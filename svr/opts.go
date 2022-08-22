package svr

import (
	"errors"
	"time"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"

	"stay/cache"
	"stay/cache/mem/lru"
	"stay/chash"
	"stay/cpool"
	"stay/utils/net"
	"stay/utils/opt"
)

type Opts struct {
	hos []opt.Opt[chash.Opts]

	*baseOpts
}

type baseOpts struct {
	hot cache.Cache
	rnd int

	src Src

	sos []grpc.ServerOption

	mode string
	tps  []string

	mws gin.HandlersChain

	http string

	*peerOpts
}

type peerOpts struct {
	pos []opt.Opt[cpool.Opts]

	*peerBaseOpts
}

type peerBaseOpts struct {
	tmo time.Duration

	cos []grpc.CallOption
}

func newOpts() (*Opts, error) {
	o, err := newBaseOpts()
	if err != nil {
		return nil, err
	}

	os := &Opts{
		baseOpts: o,
	}

	return os, nil
}

const rnd = 10

func newBaseOpts() (*baseOpts, error) {
	l, err := lru.New()
	if err != nil {
		return nil, err
	}

	o := &baseOpts{
		hot: l,
		rnd: rnd,

		src: src,

		peerOpts: newPeerOpts(),
	}

	return o, nil
}

func newPeerOpts() *peerOpts {
	return &peerOpts{
		peerBaseOpts: newPeerBaseOpts(),
	}
}

const tmo = 2 * time.Second

func newPeerBaseOpts() *peerBaseOpts {
	return &peerBaseOpts{
		tmo: tmo,
	}
}

func WithHos(os ...opt.Opt[chash.Opts]) opt.Opt[Opts] {
	return func(o *Opts) error {
		if len(os) == 0 {
			o.hos = nil
		} else {
			o.hos = os
		}

		return nil
	}
}

func WithHot(c cache.Cache) opt.Opt[Opts] {
	return func(o *Opts) error {
		o.hot = c
		return nil
	}
}

var ErrInvRnd = errors.New("invalid random range for server hot cache")

func WithRnd(r int) opt.Opt[Opts] {
	return func(o *Opts) error {
		switch {
		case r < 0:
			return ErrInvRnd

		case r == 0:
			o.rnd = rnd
		default:
			o.rnd = r
		}

		return nil
	}
}

func WithSrc(s Src) opt.Opt[Opts] {
	return func(o *Opts) error {
		o.src = s
		return nil
	}
}

func WithSrcFn(f SrcFn) opt.Opt[Opts] {
	return func(o *Opts) error {
		if f == nil {
			o.src = src
		} else {
			o.src = f
		}

		return nil
	}
}

func WithSos(os ...grpc.ServerOption) opt.Opt[Opts] {
	return func(o *Opts) error {
		if len(os) == 0 {
			o.sos = nil
		} else {
			o.sos = os
		}

		return nil
	}
}

var ErrInvMode = errors.New("invalid Gin mode for http server")

func WithMode(m string) opt.Opt[Opts] {
	return func(o *Opts) error {
		if m != "" && m != gin.EnvGinMode && m != gin.DebugMode && m != gin.ReleaseMode && m != gin.TestMode {
			return ErrInvMode
		}

		o.mode = m
		return nil
	}
}

func WithTps(ps ...string) opt.Opt[Opts] {
	return func(o *Opts) error {
		if len(ps) == 0 {
			o.tps = nil
		} else {
			o.tps = ps
		}

		return nil
	}
}

func WithMws(ms ...gin.HandlerFunc) opt.Opt[Opts] {
	return func(o *Opts) error {
		if len(ms) == 0 {
			o.mws = nil
		} else {
			o.mws = ms
		}

		return nil
	}
}

func WithHttp(h string) opt.Opt[Opts] {
	return func(o *Opts) error {
		if err := net.ChkAddr(h); err != nil {
			return err
		}

		o.http = h
		return nil
	}
}

func WithPos(os ...opt.Opt[cpool.Opts]) opt.Opt[Opts] {
	return func(o *Opts) error {
		if len(os) == 0 {
			o.pos = nil
		} else {
			o.pos = os
		}

		return nil
	}
}

var ErrInvTmo = errors.New("invalid peer RPC timeout")

func WithTmo(t time.Duration) opt.Opt[Opts] {
	return func(o *Opts) error {
		switch {
		case t < -1:
			return ErrInvTmo

		case t == -1:
			o.tmo = tmo
		default:
			o.tmo = t
		}

		return nil
	}
}

func WithCos(os ...grpc.CallOption) opt.Opt[Opts] {
	return func(o *Opts) error {
		if len(os) == 0 {
			o.cos = nil
		} else {
			o.cos = os
		}

		return nil
	}
}
