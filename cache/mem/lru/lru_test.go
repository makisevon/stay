package lru

import (
	"errors"
	"reflect"
	"testing"

	"stay/cache"
)

func TestNew(t *testing.T) {
	cases := []Mem{0, 1}
	for i, l := range cases {
		if _, err := New(WithLim(l * MemM)); err != nil {
			t.Errorf("[case %d] new err = %v", i, err)
		}
	}
}

func TestLru_Get(t *testing.T) {
	cases := [][3]any{{0, "k", []byte("v")}, {1, "k", []byte("v")}}
	for i, cas := range cases {
		l, err := New(WithLim(Mem(cas[0].(int)) * MemM))
		if err != nil {
			t.Errorf("[case %d] new err = %v", i, err)
			continue
		}

		k, tar := cas[1].(string), cas[2].([]byte)
		if err := l.Set(k, tar); err != nil {
			t.Errorf("[case %d] set err = %v", i, err)
			continue
		}

		res, err := l.Get(k)
		if err != nil {
			t.Errorf("[case %d] get err = %v", i, err)
			continue
		}

		if !reflect.DeepEqual(res, tar) {
			t.Errorf("[case %d] res = %v, tar = %v", i, res, tar)
		}
	}
}

func TestLru_Set(t *testing.T) {
	cases := [][]any{
		{0, [2]string{"k0", "v0"}, [2]any{"k0", toMBytes("v0")}},
		{0, [2]string{"k0", "v0"}, [2]string{"k1", "v1"}, [2]any{"k0", toMBytes("v0")}},
		{0, [2]string{"k0", "v0"}, [2]string{"k1", "v1"}, [2]any{"k1", toMBytes("v1")}},

		{2, [2]string{"k0", "v0"}, [2]any{"k0", toMBytes("v0")}},
		{2, [2]string{"k0", "v0"}, [2]string{"k1", "v1"}, [2]any{"k0", []byte(nil)}},
		{2, [2]string{"k0", "v0"}, [2]string{"k1", "v1"}, [2]any{"k1", toMBytes("v1")}},

		{6, [2]string{"k0", "v0"}, [2]any{"k0", toMBytes("v0")}},
		{6, [2]string{"k0", "v0"}, [2]string{"k1", "v1"}, [2]any{"k0", toMBytes("v0")}},
		{6, [2]string{"k0", "v0"}, [2]string{"k1", "v1"}, [2]any{"k1", toMBytes("v1")}},
		{6, [2]string{"k0", "v0"}, [2]string{"k1", "v1"}, [2]string{"k2", "v2"}, [2]any{"k0", toMBytes("v0")}},
		{6, [2]string{"k0", "v0"}, [2]string{"k1", "v1"}, [2]string{"k2", "v2"}, [2]any{"k1", toMBytes("v1")}},
		{6, [2]string{"k0", "v0"}, [2]string{"k1", "v1"}, [2]string{"k2", "v2"}, [2]any{"k2", toMBytes("v2")}},

		{
			6,
			[2]string{"k0", "v0"}, [2]string{"k1", "v1"}, [2]string{"k2", "v2"},
			[2]string{"k3", "v3"},
			[2]any{"k0", []byte(nil)},
		}, {
			6,
			[2]string{"k0", "v0"}, [2]string{"k1", "v1"}, [2]string{"k2", "v2"},
			[2]string{"k3", "v3"},
			[2]any{"k1", toMBytes("v1")},
		}, {
			6,
			[2]string{"k0", "v0"}, [2]string{"k1", "v1"}, [2]string{"k2", "v2"},
			[2]string{"k3", "v3"},
			[2]any{"k2", toMBytes("v2")},
		}, {
			6,
			[2]string{"k0", "v0"}, [2]string{"k1", "v1"}, [2]string{"k2", "v2"},
			[2]string{"k3", "v3"},
			[2]any{"k3", toMBytes("v3")},
		}, {
			6,
			[2]string{"k0", "v0"}, [2]string{"k1", "v1"}, [2]string{"k2", "v2"},
			[2]string{"k3", "v3"}, [2]string{"k4", "v4"},
			[2]any{"k1", []byte(nil)},
		}, {
			6,
			[2]string{"k0", "v0"}, [2]string{"k1", "v1"}, [2]string{"k2", "v2"},
			[2]string{"k3", "v3"}, [2]string{"k4", "v4"},
			[2]any{"k2", toMBytes("v2")},
		}, {
			6,
			[2]string{"k0", "v0"}, [2]string{"k1", "v1"}, [2]string{"k2", "v2"},
			[2]string{"k3", "v3"}, [2]string{"k4", "v4"},
			[2]any{"k3", toMBytes("v3")},
		}, {
			6,
			[2]string{"k0", "v0"}, [2]string{"k1", "v1"}, [2]string{"k2", "v2"},
			[2]string{"k3", "v3"}, [2]string{"k4", "v4"},
			[2]any{"k4", toMBytes("v4")},
		},

		{
			6,
			[2]string{"k0", "v0"}, [2]string{"k1", "v1"}, [2]string{"k2", "v2"},
			[2]string{"k0", "v0"},
			[2]any{"k0", toMBytes("v0")},
		}, {
			6,
			[2]string{"k0", "v0"}, [2]string{"k1", "v1"}, [2]string{"k2", "v2"},
			[2]string{"k0", "v0"},
			[2]any{"k1", toMBytes("v1")},
		}, {
			6,
			[2]string{"k0", "v0"}, [2]string{"k1", "v1"}, [2]string{"k2", "v2"},
			[2]string{"k0", "v0"},
			[2]any{"k2", toMBytes("v2")},
		}, {
			6,
			[2]string{"k0", "v0"}, [2]string{"k1", "v1"}, [2]string{"k2", "v2"},
			[2]string{"k0", "v0"}, [2]string{"k3", "v3"},
			[2]any{"k0", toMBytes("v0")},
		}, {
			6,
			[2]string{"k0", "v0"}, [2]string{"k1", "v1"}, [2]string{"k2", "v2"},
			[2]string{"k0", "v0"}, [2]string{"k3", "v3"},
			[2]any{"k1", []byte(nil)},
		}, {
			6,
			[2]string{"k0", "v0"}, [2]string{"k1", "v1"}, [2]string{"k2", "v2"},
			[2]string{"k0", "v0"}, [2]string{"k3", "v3"},
			[2]any{"k2", toMBytes("v2")},
		}, {
			6,
			[2]string{"k0", "v0"}, [2]string{"k1", "v1"}, [2]string{"k2", "v2"},
			[2]string{"k0", "v0"}, [2]string{"k3", "v3"},
			[2]any{"k3", toMBytes("v3")},
		},
	}

L:
	for i, cas := range cases {
		l, err := New(WithLim(Mem(cas[0].(int)) * MemM))
		if err != nil {
			t.Errorf("[case %d] new err = %v", i, err)
			continue
		}

		lc := len(cas) - 1
		for j := 1; j < lc; j++ {
			e := cas[j].([2]string)
			if err := l.Set(string(toMBytes(e[0])), toMBytes(e[1])); err != nil {
				t.Errorf("[case %d] set err = %v", i, err)
				continue L
			}
		}

		e := cas[lc].([2]any)

		res, err := l.Get(string(toMBytes(e[0].(string))))
		if err != nil && !errors.Is(err, cache.ErrCacheMs) {
			t.Errorf("[case %d] get err = %v", i, err)
			continue
		}

		if tar := e[1].([]byte); !reflect.DeepEqual(res, tar) {
			t.Errorf("[case %d] res = %v, tar = %v", i, res, tar)
		}
	}
}

const mil = 1 << 20

func toMBytes(s string) []byte {
	bs := make([]byte, mil)
	copy(bs, s)

	return bs
}

func TestLru_Del(t *testing.T) {
	cases := [][3]any{{0, "k", []byte("v")}, {1, "k", []byte("v")}}
	for i, cas := range cases {
		l, err := New(WithLim(Mem(cas[0].(int)) * MemM))
		if err != nil {
			t.Errorf("[case %d] new err = %v", i, err)
			continue
		}

		k := cas[1].(string)
		if err := l.Set(k, cas[2].([]byte)); err != nil {
			t.Errorf("[case %d] set err = %v", i, err)
			continue
		}

		if err := l.Del(k); err != nil {
			t.Errorf("[case %d] del err = %v", i, err)
			continue
		}

		if _, err = l.Get(k); err == nil {
			t.Errorf("[case %d] key should be deleted", i)
			continue
		}

		if !errors.Is(err, cache.ErrCacheMs) {
			t.Errorf("[case %d] get err = %v", i, err)
		}
	}
}
