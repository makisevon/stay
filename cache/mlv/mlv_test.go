package mlv

import (
	"errors"
	"reflect"
	"testing"

	"stay/cache"
	"stay/cache/mem/raw"
)

func TestNew(t *testing.T) {
	cases := [][]cache.Cache{{raw.New()}, {raw.New(), raw.New()}}
	for i, cs := range cases {
		if _, err := New(cs...); err != nil {
			t.Errorf("[case %d] new err = %v", i, err)
		}
	}
}

func TestSto(t *testing.T) {
	cases := [][]any{
		{[]cache.Cache{raw.New()}, "k", []byte("v0")},
		{[]cache.Cache{raw.New()}, "k", []byte("v0"), []byte("v1")},

		{[]cache.Cache{raw.New(), raw.New()}, "k", []byte("v0")},
		{[]cache.Cache{raw.New(), raw.New()}, "k", []byte("v0"), []byte("v1")},
	}

L:
	for i, cas := range cases {
		c, err := New(cas[0].([]cache.Cache)...)
		if err != nil {
			t.Errorf("[case %d] new err = %v", i, err)
			continue
		}

		l, k := len(cas)-1, cas[1].(string)
		var tar []byte

		for j := 2; j <= l; j++ {
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
