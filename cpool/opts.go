package cpool

import (
	"errors"
	"runtime"
	"time"

	"google.golang.org/grpc"

	"stay/utils/opt"
)

type Opts struct {
	cap int

	*connOpts
}

type connOpts struct {
	tmo time.Duration

	dos []grpc.DialOption
}

func newOpts() *Opts {
	return &Opts{
		cap: runtime.GOMAXPROCS(0) * 3 / 2,

		connOpts: newConnOpts(),
	}
}

const tmo = 3 * time.Second

func newConnOpts() *connOpts {
	return &connOpts{
		tmo: tmo,
	}
}

var ErrInvCap = errors.New("invalid capacity of connection pool")

func WithCap(c int) opt.Opt[Opts] {
	return func(o *Opts) error {
		switch {
		case c < 0:
			return ErrInvCap

		case c == 0:
			o.cap = runtime.GOMAXPROCS(0) * 3 / 2
		default:
			o.cap = c
		}

		return nil
	}
}

var ErrInvTmo = errors.New("invalid connection dial timeout")

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

func WithDos(os ...grpc.DialOption) opt.Opt[Opts] {
	return func(o *Opts) error {
		if len(os) == 0 {
			o.dos = nil
		} else {
			o.dos = os
		}

		return nil
	}
}
