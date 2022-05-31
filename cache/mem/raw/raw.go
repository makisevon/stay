package raw

import (
	"sync"

	"stay/cache"
)

type Raw struct {
	cache sync.Map
}

func New() *Raw {
	return new(Raw)
}

var _ cache.Cache = (*Raw)(nil)

func (r *Raw) Get(k string) ([]byte, error) {
	if v, ok := r.cache.Load(k); ok {
		return v.([]byte), nil
	}

	return nil, cache.ErrCacheMs
}

func (r *Raw) Set(k string, v []byte) error {
	r.cache.Store(k, v)
	return nil
}

func (r *Raw) Del(k string) error {
	r.cache.Delete(k)
	return nil
}
