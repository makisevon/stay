package cbf

import (
	"errors"
	"hash/fnv"

	"stay/utils/hash"
	"stay/utils/opt"
)

type Opts struct {
	num  int
	prob float64

	hash hash.Hash
}

const (
	num  = 4000
	prob = 1e-7
)

func newOpts() *Opts {
	return &Opts{
		num:  num,
		prob: prob,

		hash: genHash(),
	}
}

var ErrInvNum = errors.New("invalid number of keys for CBF")

func WithNum(n int) opt.Opt[Opts] {
	return func(o *Opts) error {
		switch {
		case n < 0:
			return ErrInvNum

		case n == 0:
			o.num = num
		default:
			o.num = n
		}

		return nil
	}
}

var ErrInvProb = errors.New("invalid probability of false positives for CBF")

func WithProb(p float64) opt.Opt[Opts] {
	return func(o *Opts) error {
		switch {
		case p < -prob:
			return ErrInvProb

		case p < prob:
			o.prob = prob
		default:
			o.prob = p
		}

		return nil
	}
}

func WithHash(h hash.Hash) opt.Opt[Opts] {
	return func(o *Opts) error {
		if h == nil {
			o.hash = genHash()
		} else {
			o.hash = h
		}

		return nil
	}
}

func WithStdHash[T uint32 | uint64](h hash.StdHash[T]) opt.Opt[Opts] {
	return func(o *Opts) error {
		if h == nil {
			o.hash = genHash()
		} else {
			o.hash = hash.Wrap(h)
		}

		return nil
	}
}

func genHash() hash.Hash {
	f := fnv.New32()
	return hash.Wrap(func(bs []byte) uint32 {
		f.Reset()
		_, _ = f.Write(bs)

		return f.Sum32()
	})
}
