package cbf

import (
	"errors"
	"reflect"
	"testing"
	"time"

	"stay/cache"
	"stay/cache/mem/raw"
)

func TestSto(t *testing.T) {
	cases := [][]any{{"k", []byte("v0")}, {"k", []byte("v0"), []byte("v1")}}

L:
	for i, cas := range cases {
		c, err := WithCbf(raw.New())
		if err != nil {
			t.Errorf("[case %d] new err = %v", i, err)
			continue
		}

		l, k := len(cas)-1, cas[0].(string)
		var tar []byte

		for j := 1; j <= l; j++ {
			v := cas[j].([]byte)
			if err := c.Set(k, v); err != nil {
				t.Errorf("[case %d] set err = %v", i, err)
				continue L
			}

			if j == l {
				tar = v
			}
		}

		res, err := c.Get(k)
		if err != nil {
			t.Errorf("[case %d] get err = %v", i, err)
			continue
		}

		if !reflect.DeepEqual(res, tar) {
			t.Errorf("[case %d] res = %v, tar = %v", i, res, tar)
			continue
		}

		if err := c.Del(k); err != nil {
			t.Errorf("[case %d] del err = %v", i, err)
			continue
		}

		if _, err = c.Get(k); err == nil {
			t.Errorf("[case %d] key should be deleted", i)
			continue
		}

		if !errors.Is(err, cache.ErrCacheMs) {
			t.Errorf("[case %d] get err = %v", i, err)
		}
	}
}

type benchCache struct {
	cache cache.Cache

	dur time.Duration
}

func newBenchCache() *benchCache {
	return &benchCache{
		cache: raw.New(),

		dur: time.Millisecond,
	}
}

var _ cache.Cache = (*benchCache)(nil)

func (c *benchCache) Get(k string) ([]byte, error) {
	time.Sleep(c.dur)
	return c.cache.Get(k)
}

func (c *benchCache) Set(k string, v []byte) error {
	return c.cache.Set(k, v)
}

func (c *benchCache) Del(k string) error {
	return c.cache.Del(k)
}

func BenchmarkCbf_Get(b *testing.B) {
	c, err := WithCbf(newBenchCache())
	if err != nil {
		b.Errorf("new err = %v", err)
	}

	benchGet(b, c)
}

func BenchmarkCbf_Get_Bl(b *testing.B) {
	benchGet(b, newBenchCache())
}

func benchGet(b *testing.B, c cache.Cache) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if _, err := c.Get("k"); !errors.Is(err, cache.ErrCacheMs) {
				b.Errorf("get err = %v", err)
				return
			}
		}
	})
}
