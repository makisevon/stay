package rds

import (
	"errors"
	"reflect"
	"testing"
	"time"

	"stay/cache"
)

func TestLft(t *testing.T) {
	cases := []int{0, 1}
	for i, ttl := range cases {
		r, err := New(WithTtl(time.Duration(ttl) * time.Second))
		if err != nil {
			t.Errorf("[case %d] new err = %v", i, err)
			continue
		}

		if err := r.Fin(); err != nil {
			t.Errorf("[case %d] fin err = %v", i, err)
		}
	}
}

func TestRds_Get(t *testing.T) {
	cases := [][3]any{{0, "k", []byte("v")}, {1, "k", []byte("v")}}
	for i, cas := range cases {
		ttl := cas[0].(int)

		r, err := New(WithTtl(time.Duration(ttl) * time.Second))
		if err != nil {
			t.Errorf("[case %d] new err = %v", i, err)
			continue
		}

		k, tar := cas[1].(string), cas[2].([]byte)
		if err := r.Set(k, tar); err != nil {
			t.Errorf("[case %d] set err = %v", i, err)

			_ = r.Fin()
			continue
		}

		res, err := r.Get(k)
		if err != nil {
			t.Errorf("[case %d] get err = %v", i, err)

			_ = r.Fin()
			continue
		}

		if !reflect.DeepEqual(res, tar) {
			t.Errorf("[case %d] res = %v, tar = %v", i, res, tar)

			_ = r.Fin()
			continue
		}

		if ttl == 0 {
			_ = r.Fin()
			continue
		}

		time.Sleep(time.Duration(ttl)*time.Second + time.Millisecond)
		if _, err = r.Get(k); err == nil {
			t.Errorf("[case %d] key should be expired", i)

			_ = r.Fin()
			continue
		}

		if !errors.Is(err, cache.ErrCacheMs) {
			t.Errorf("[case %d] get err = %v", i, err)

			_ = r.Fin()
			continue
		}

		if err := r.Fin(); err != nil {
			t.Errorf("[case %d] fin err = %v", i, err)
		}
	}
}

func TestRds_Set(t *testing.T) {
	cases := [][]any{
		{0, "k", []byte("v0")}, {0, "k", []byte("v0"), []byte("v1")},
		{1, "k", []byte("0v")}, {1, "k", []byte("v0"), []byte("v1")},
	}

L:
	for i, cas := range cases {
		r, err := New(WithTtl(time.Duration(cas[0].(int)) * time.Second))
		if err != nil {
			t.Errorf("[case %d] new err = %v", i, err)
			continue
		}

		l, k := len(cas)-1, cas[1].(string)
		var tar []byte

		for j := 2; j <= l; j++ {
			val := cas[j].([]byte)
			if err := r.Set(k, val); err != nil {
				t.Errorf("[case %d] set err = %v", i, err)

				_ = r.Fin()
				continue L
			}

			if j == l {
				tar = val
			}
		}

		res, err := r.Get(k)
		if err != nil {
			t.Errorf("[case %d] get err = %v", i, err)

			_ = r.Fin()
			continue
		}

		if !reflect.DeepEqual(res, tar) {
			t.Errorf("[case %d] res = %v, tar = %v", i, res, tar)

			_ = r.Fin()
			continue
		}

		if err := r.Fin(); err != nil {
			t.Errorf("[case %d] fin err = %v", i, err)
		}
	}
}

func TestRds_Del(t *testing.T) {
	cases := [][3]any{{0, "k", []byte("v")}, {1, "k", []byte("v")}}
	for i, cas := range cases {
		r, err := New(WithTtl(time.Duration(cas[0].(int)) * time.Second))
		if err != nil {
			t.Errorf("[case %d] new err = %v", i, err)
			continue
		}

		k := cas[1].(string)
		if err := r.Set(k, cas[2].([]byte)); err != nil {
			t.Errorf("[case %d] set err = %v", i, err)

			_ = r.Fin()
			continue
		}

		if err := r.Del(k); err != nil {
			t.Errorf("[case %d] del err = %v", i, err)

			_ = r.Fin()
			continue
		}

		if _, err = r.Get(k); err == nil {
			t.Errorf("[case %d] key should be deleted", i)

			_ = r.Fin()
			continue
		}

		if !errors.Is(err, cache.ErrCacheMs) {
			t.Errorf("[case %d] get err = %v", i, err)

			_ = r.Fin()
			continue
		}

		if err := r.Fin(); err != nil {
			t.Errorf("[case %d] fin err = %v", i, err)
		}
	}
}
