package chash

import (
	"errors"
	"hash/crc32"

	"stay/utils/hash"
	"stay/utils/opt"
)

type Opts struct {
	rep int

	hash hash.Hash
}

const rep = 3

var _hash = hash.Wrap(crc32.ChecksumIEEE)

func newOpts() *Opts {
	return &Opts{
		rep: rep,

		hash: _hash,
	}
}

var ErrInvRep = errors.New("invalid key replicas for consistent hash")

func WithRep(r int) opt.Opt[Opts] {
	return func(o *Opts) error {
		switch {
		case r < 0:
			return ErrInvRep

		case r == 0:
			o.rep = rep
		default:
			o.rep = r
		}

		return nil
	}
}

func WithHash(h hash.Hash) opt.Opt[Opts] {
	return func(o *Opts) error {
		if h == nil {
			o.hash = _hash
		} else {
			o.hash = h
		}

		return nil
	}
}

func WithStdHash[T uint32 | uint64](h hash.StdHash[T]) opt.Opt[Opts] {
	return func(o *Opts) error {
		if h == nil {
			o.hash = _hash
		} else {
			o.hash = hash.Wrap(h)
		}

		return nil
	}
}
