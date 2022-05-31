package lru

import (
	"errors"

	"stay/utils/opt"
)

type Opts struct {
	lim Mem
}

const lim = MemM << 4

func newOpts() *Opts {
	return &Opts{
		lim: lim,
	}
}

var ErrInvLim = errors.New("invalid memory limit of LRU cache")

func WithLim(l Mem) opt.Opt[Opts] {
	return func(o *Opts) error {
		switch {
		case l < 0:
			return ErrInvLim

		case l == 0:
			o.lim = lim
		default:
			o.lim = l
		}

		return nil
	}
}
