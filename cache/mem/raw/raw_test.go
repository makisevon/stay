package raw

import (
	"errors"
	"reflect"
	"sync"
	"testing"

	"stay/cache"
)

func TestSto(t *testing.T) {
	cases := [][]any{{"k", []byte("v0")}, {"k", []byte("v0"), []byte("v1")}}

L:
	for i, cas := range cases {
		r, l, k := New(), len(cas)-1, cas[0].(string)
		var tar []byte

		for j := 1; j <= l; j++ {
			v := cas[j].([]byte)
			if err := r.Set(k, v); err != nil {
				t.Errorf("[case %d] set err = %v", i, err)
				continue L
			}

			if j == l {
				tar = v
			}
		}

		res, err := r.Get(k)
		if err != nil {
			t.Errorf("[case %d] get err = %v", i, err)
			continue
		}

		if !reflect.DeepEqual(res, tar) {
			t.Errorf("[case %d] res = %v, tar = %v", i, res, tar)
			continue
		}

		if err := r.Del(k); err != nil {
			t.Errorf("[case %d] del err = %v", i, err)
			continue
		}

		if _, err = r.Get(k); err == nil {
			t.Errorf("[case %d] key should be deleted", i)
			continue
		}

		if !errors.Is(err, cache.ErrCacheMs) {
			t.Errorf("[case %d] get err = %v", i, err)
		}
	}
}

func BenchmarkSto(b *testing.B) {
	benchSto(b, New())
}

type benchCache struct {
	cache map[string][]byte

	mtx sync.RWMutex
}

func newBenchCache() *benchCache {
	return &benchCache{
		cache: make(map[string][]byte),
	}
}

var _ cache.Cache = (*benchCache)(nil)

func (c *benchCache) Get(k string) ([]byte, error) {
	c.mtx.RLock()
	defer c.mtx.RUnlock()

	if v, ok := c.cache[k]; ok {
		return v, nil
	}

	return nil, cache.ErrCacheMs
}

func (c *benchCache) Set(k string, v []byte) error {
	c.mtx.Lock()
	defer c.mtx.Unlock()

	c.cache[k] = v
	return nil
}

func (c *benchCache) Del(k string) error {
	c.mtx.Lock()
	defer c.mtx.Unlock()

	delete(c.cache, k)
	return nil
}

func BenchmarkSto_Bl(b *testing.B) {
	benchSto(b, newBenchCache())
}

const key = "k"

var val = []byte("v")

func benchSto(b *testing.B, c cache.Cache) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if err := c.Set(key, val); err != nil {
				b.Errorf("set err = %v", err)
				return
			}

			if _, err := c.Get(key); err != nil && !errors.Is(err, cache.ErrCacheMs) {
				b.Errorf("get err = %v", err)
				return
			}

			if err := c.Del(key); err != nil {
				b.Errorf("del err = %v", err)
				return
			}
		}
	})
}
