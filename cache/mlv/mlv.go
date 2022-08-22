package mlv

import (
	"errors"
	"sync"

	"stay/cache"
)

type Mlv struct {
	caches []cache.Cache

	mtx sync.RWMutex
}

var ErrFewCaches = errors.New("too few caches to new multi-level cache")

func New(cs ...cache.Cache) (*Mlv, error) {
	if len(cs) < 1 {
		return nil, ErrFewCaches
	}

	m := &Mlv{
		caches: cs,
	}

	return m, nil
}

var _ cache.Cache = (*Mlv)(nil)

func (m *Mlv) Get(k string) ([]byte, error) {
	m.mtx.RLock()
	defer m.mtx.RUnlock()

	for _, c := range m.caches {
		if v, err := c.Get(k); err == nil {
			return v, nil
		}
	}

	return nil, cache.ErrCacheMs
}

func (m *Mlv) Set(k string, v []byte) error {
	m.mtx.Lock()
	defer m.mtx.Unlock()

	var err error
	for _, c := range m.caches {
		if e := c.Set(k, v); err == nil && e != nil {
			err = e
		}
	}

	return err
}

func (m *Mlv) Del(k string) error {
	m.mtx.Lock()
	defer m.mtx.Unlock()

	var err error
	for _, c := range m.caches {
		if e := c.Del(k); err == nil && e != nil {
			err = e
		}
	}

	return err
}
