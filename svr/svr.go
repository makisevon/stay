package svr

import (
	"context"
	"math/rand"
	_net "net"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
	"golang.org/x/sync/singleflight"
	"google.golang.org/grpc"

	"stay/cache"
	"stay/chash"
	"stay/logs"
	"stay/pb"
	"stay/utils/net"
	"stay/utils/opt"
)

type Svr struct {
	intl *svr
}

type svr struct {
	cache cache.Cache

	peers map[string]*peer
	hash  *chash.CHash

	addr string

	mtx sync.Mutex
	gsf singleflight.Group
	dsf singleflight.Group

	*baseOpts

	pb.UnimplementedCacheServer
}

func New(c cache.Cache, a string, os ...opt.Opt[Opts]) (*Svr, error) {
	i, err := newSvr(c, a, os...)
	if err != nil {
		return nil, err
	}

	s := &Svr{
		intl: i,
	}

	return s, nil
}

func newSvr(c cache.Cache, a string, os ...opt.Opt[Opts]) (*svr, error) {
	if err := net.ChkAddr(a); err != nil {
		return nil, err
	}

	o, err := newOpts()
	if err != nil {
		return nil, err
	}

	if err := opt.Apply(o, os); err != nil {
		return nil, err
	}

	h, err := chash.New(o.hos...)
	if err != nil {
		return nil, err
	}

	s := &svr{
		cache: c,

		peers: make(map[string]*peer),
		hash:  h,

		addr: a,

		baseOpts: o.baseOpts,
	}

	return s, nil
}

func (s *Svr) Reg(as ...string) error {
	return s.intl.reg(as)
}

func (s *Svr) Run() error {
	return s.intl.run()
}

func (s *svr) reg(as []string) error {
	s.mtx.Lock()
	defer s.mtx.Unlock()

	ps := make([]*peer, len(as))
	for i, a := range as {
		var err error
		if ps[i], err = newPeer(a, s.peerOpts); err != nil {
			return err
		}
	}

	for i, a := range as {
		s.peers[a] = ps[i]
	}

	s.hash.Add(as...)
	return nil
}

func (s *svr) run() error {
	s.mtx.Lock()
	defer s.mtx.Unlock()

	var g errgroup.Group
	g.Go(s.runRpc)
	g.Go(s.runHttp)

	return g.Wait()
}

const nw = "tcp"

func (s *svr) runRpc() error {
	l, err := _net.Listen(nw, s.addr)
	if err != nil {
		return err
	}

	r := grpc.NewServer(s.sos...)
	defer r.GracefulStop()

	pb.RegisterCacheServer(r, s)

	logs.Info("[svr %s] RPC listening at %s", s.addr, l.Addr().String())
	return r.Serve(l)
}

const _key = "key"

func (s *svr) runHttp() error {
	l, err := _net.Listen(nw, s.http)
	if err != nil {
		return err
	}

	gin.SetMode(s.mode)

	r := gin.New()
	if err := r.SetTrustedProxies(s.tps); err != nil {
		return err
	}

	r.Use(gin.Recovery())
	r.Use(s.mws...)

	r.GET("/get/:"+_key, s.get)
	r.POST("/set/:"+_key, s.set)
	r.POST("/del/:"+_key, s.del)

	logs.Info("[svr %s] HTTP listening at %s", s.addr, l.Addr().String())
	return r.RunListener(l)
}

var _ pb.CacheServer = (*svr)(nil)

func (s *svr) Get(_ context.Context, r *pb.GetReq) (*pb.GetResp, error) {
	k := r.GetKey()
	logs.Info("[svr %s] get key = %s", s.addr, k)

	v, err, _ := s.gsf.Do(k, func() (any, error) {
		for _, c := range []cache.Cache{s.cache, s.hot} {
			if v, err := c.Get(k); err == nil {
				return v, nil
			}
		}

		p := s.pick(k)
		if p == nil {
			v, err := s.src.Get(k)
			if err == nil {
				_ = s.cache.Set(k, v)
			}

			return v, err
		}

		v, err := p.get(k)
		if err == nil && rand.Intn(s.rnd) == 0 {
			_ = s.hot.Set(k, v)
		}

		return v, err
	})

	if err != nil {
		logs.Err("[svr %s] get err = %v, key = %s", s.addr, err, k)
	}

	resp := &pb.GetResp{
		Val: v.([]byte),
		Err: unwrapErr(err),
	}

	return resp, nil
}

func (s *svr) Set(_ context.Context, r *pb.SetReq) (*pb.SetResp, error) {
	k, v := r.GetKey(), r.GetVal()
	logs.Info("[svr %s] set key = %s", s.addr, k)

	err := func() error {
		p := s.pick(k)
		if p == nil {
			return s.cache.Set(k, v)
		}

		if err := p.set(k, v); err != nil {
			return err
		}

		_ = s.hot.Del(k)
		return nil
	}()

	if err != nil {
		logs.Err("[svr %s] set err = %v, key = %s", s.addr, err, k)
	}

	resp := &pb.SetResp{
		Err: unwrapErr(err),
	}

	return resp, nil
}

func (s *svr) Del(_ context.Context, r *pb.DelReq) (*pb.DelResp, error) {
	k := r.GetKey()
	logs.Info("[svr %s] del key = %s", s.addr, k)

	_, err, _ := s.dsf.Do(k, func() (any, error) {
		p := s.pick(k)
		if p == nil {
			return nil, s.cache.Del(k)
		}

		if err := p.del(k); err != nil {
			return nil, err
		}

		_ = s.hot.Del(k)
		return nil, nil
	})

	if err != nil {
		logs.Err("[svr %s] del err = %v, key = %s", s.addr, err, k)
	}

	resp := &pb.DelResp{
		Err: unwrapErr(err),
	}

	return resp, nil
}

func (s *svr) pick(k string) *peer {
	if p, ok := s.peers[s.hash.Get(k)]; ok && p.addr != s.addr {
		return p
	}

	return nil
}

func unwrapErr(e error) string {
	if e == nil {
		return ""
	}

	return e.Error()
}

const (
	_val = "val"
	_err = "err"
)

func (s *svr) get(c *gin.Context) {
	req := &pb.GetReq{
		Key: c.Param(_key),
	}

	r, _ := s.Get(nil, req)
	c.JSON(http.StatusOK, &gin.H{
		_val: r.GetVal(),
		_err: r.GetErr(),
	})
}

func (s *svr) set(c *gin.Context) {
	req := &pb.SetReq{
		Key: c.Param(_key),
		Val: []byte(c.PostForm(_val)),
	}

	r, _ := s.Set(nil, req)
	c.JSON(http.StatusOK, &gin.H{
		_err: r.GetErr(),
	})
}

func (s *svr) del(c *gin.Context) {
	req := &pb.DelReq{
		Key: c.Param(_key),
	}

	r, _ := s.Del(nil, req)
	c.JSON(http.StatusOK, &gin.H{
		_err: r.GetErr(),
	})
}
