package exp

import (
	"errors"
	"runtime"
	"time"

	"stay/utils/opt"
)

type Opts struct {
	ttl, span time.Duration

	proc int
}

const ttl = 10 * time.Second

func newOpts() *Opts {
	return &Opts{
		ttl: ttl,

		proc: runtime.GOMAXPROCS(0) * 3 / 2,
	}
}

var ErrInvTtl = errors.New("invalid TTL for expiring cache")

func WithTtl(t time.Duration) opt.Opt[Opts] {
	return func(o *Opts) error {
		switch {
		case t < 0:
			return ErrInvTtl

		case t == 0:
			o.ttl = ttl
		default:
			o.ttl = t
		}

		return nil
	}
}

var ErrInvSpan = errors.New("invalid span between active cache expiring")

func WithSpan(s time.Duration) opt.Opt[Opts] {
	return func(o *Opts) error {
		if s < 0 {
			return ErrInvSpan
		}

		o.span = s
		return nil
	}
}

var ErrInvProc = errors.New("invalid max number of processes for active cache expiring")

func WithProc(p int) opt.Opt[Opts] {
	return func(o *Opts) error {
		switch {
		case p < 0:
			return ErrInvProc

		case p == 0:
			o.proc = runtime.GOMAXPROCS(0) * 3 / 2
		default:
			o.proc = p
		}

		return nil
	}
}
