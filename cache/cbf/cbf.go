package cbf

import (
	"errors"
	"math"
	"sync"

	"stay/cache"
	"stay/utils/hash"
	"stay/utils/opt"
)

type Cbf struct {
	cache cache.Cache

	bits []int

	size, ops int

	hash hash.Hash

	mtx sync.RWMutex
}

func WithCbf(c cache.Cache, os ...opt.Opt[Opts]) (*Cbf, error) {
	o := newOpts()
	if err := opt.Apply(o, os); err != nil {
		return nil, err
	}

	n := float64(o.num)
	s := int(math.Ceil(-n * math.Log(o.prob) / math.Ln2 / math.Ln2))

	cbf := &Cbf{
		cache: c,

		bits: make([]int, s),

		size: s,
		ops:  int(math.Round(float64(s) / n * math.Ln2)),

		hash: o.hash,
	}

	return cbf, nil
}

var _ cache.Cache = (*Cbf)(nil)

func (c *Cbf) Get(k string) ([]byte, error) {
	if !c.get(k) {
		return nil, cache.ErrCacheMs
	}

	v, err := c.cache.Get(k)
	if err != nil && errors.Is(err, cache.ErrCacheMs) {
		c.del(k)
	}

	return v, err
}

func (c *Cbf) Set(k string, v []byte) error {
	if err := c.cache.Set(k, v); err != nil {
		return err
	}

	if !c.get(k) {
		c.set(k)
	}

	return nil
}

func (c *Cbf) Del(k string) error {
	if err := c.cache.Del(k); err != nil {
		return err
	}

	c.del(k)
	return nil
}

func (c *Cbf) get(k string) bool {
	c.mtx.RLock()
	defer c.mtx.RUnlock()

	for i := 0; i < c.ops; i++ {
		if c.bits[c.idx(k, i)] == 0 {
			return false
		}
	}

	return true
}

func (c *Cbf) set(k string) {
	c.add(k, 1)
}

func (c *Cbf) del(k string) {
	c.add(k, -1)
}

func (c *Cbf) add(k string, d int) {
	c.mtx.Lock()
	defer c.mtx.Unlock()

	for i := 0; i < c.ops; i++ {
		c.bits[c.idx(k, i)] += d
	}
}

func (c *Cbf) idx(k string, i int) int {
	return (c.hash(k)%c.size + c.size + i) % c.size
}
