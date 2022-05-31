package svr

import (
	"errors"

	"stay/cpool"
	"stay/logs"
	"stay/pb"
	"stay/utils/net"
)

type peer struct {
	pool *cpool.CPool

	addr string

	*peerBaseOpts
}

func newPeer(a string, o *peerOpts) (*peer, error) {
	cp, err := cpool.New(a, o.pos...)
	if err != nil {
		return nil, err
	}

	p := &peer{
		pool: cp,

		addr: a,

		peerBaseOpts: o.peerBaseOpts,
	}

	return p, nil
}

func (p *peer) get(k string) ([]byte, error) {
	c, err := p.cli()
	if err != nil {
		return nil, err
	}

	logs.Info("[peer %s] get key = %s", p.addr, k)

	ctx, cxl := net.Ctx(p.tmo)
	defer cxl()

	req := &pb.GetReq{
		Key: k,
	}

	r, err := c.Get(ctx, req, p.cos...)
	if err != nil {
		logs.Err("[peer %s] get err = %v, key = %s", p.addr, err, k)
		return nil, err
	}

	return r.GetVal(), wrapErr(r.GetErr())
}

func (p *peer) set(k string, v []byte) error {
	c, err := p.cli()
	if err != nil {
		return err
	}

	logs.Info("[peer %s] set key = %s", p.addr, k)

	ctx, cxl := net.Ctx(p.tmo)
	defer cxl()

	req := &pb.SetReq{
		Key: k,
		Val: v,
	}

	r, err := c.Set(ctx, req, p.cos...)
	if err != nil {
		logs.Err("[peer %s] set err = %v, key = %s", p.addr, err, k)
		return err
	}

	return wrapErr(r.GetErr())
}

func (p *peer) del(k string) error {
	c, err := p.cli()
	if err != nil {
		return err
	}

	logs.Info("[peer %s] del key = %s", p.addr, k)

	ctx, cxl := net.Ctx(p.tmo)
	defer cxl()

	req := &pb.DelReq{
		Key: k,
	}

	r, err := c.Del(ctx, req, p.cos...)
	if err != nil {
		logs.Err("[peer %s] del err = %v, key = %s", p.addr, err, k)
		return err
	}

	return wrapErr(r.GetErr())
}

func (p *peer) cli() (pb.CacheClient, error) {
	c, err := p.pool.Get()
	if err != nil {
		return nil, err
	}

	return pb.NewCacheClient(c), nil
}

func wrapErr(e string) error {
	if e == "" {
		return nil
	}

	return errors.New(e)
}
