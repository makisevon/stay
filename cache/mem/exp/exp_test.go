package exp

import (
	"errors"
	"reflect"
	"runtime"
	"testing"
	"time"

	"stay/cache"
)

func TestLft(t *testing.T) {
	cases := [][3]int{{0, 0, 0}, {1, 0, 0}, {1, 1, 2}}
	for i, cas := range cases {
		tar2 := runtime.NumGoroutine()
		tar1 := tar2 + cas[2]

		e, err := New(WithTtl(time.Duration(cas[0])*time.Second), WithSpan(time.Duration(cas[1])*time.Second))
		if err != nil {
			t.Errorf("[case %d] new err = %v", i, err)
			continue
		}

		time.Sleep(time.Millisecond)
		if res := runtime.NumGoroutine(); res != tar1 {
			t.Errorf("[case %d] res = %d, tar_1 = %d", i, res, tar1)

			_ = e.Fin()
			time.Sleep(time.Millisecond)

			continue
		}

		if err := e.Fin(); err != nil {
			t.Errorf("[case %d] fin err = %v", i, err)
			continue
		}

		time.Sleep(time.Millisecond)
		if res := runtime.NumGoroutine(); res != tar2 {
			t.Errorf("[case %d] res = %d, tar_2 = %d", i, res, tar2)
		}
	}
}

func TestExp_Get(t *testing.T) {
	cases := [][4]any{{0, 0, "k", []byte("v")}, {1, 0, "k", []byte("v")}, {1, 1, "k", []byte("v")}}
	for i, cas := range cases {
		ttl := cas[0].(int)

		e, err := New(WithTtl(time.Duration(ttl)*time.Second), WithSpan(time.Duration(cas[1].(int))*time.Second))
		if err != nil {
			t.Errorf("[case %d] new err = %v", i, err)
			continue
		}

		k, tar := cas[2].(string), cas[3].([]byte)
		if err := e.Set(k, tar); err != nil {
			t.Errorf("[case %d] set err = %v", i, err)

			_ = e.Fin()
			continue
		}

		res, err := e.Get(k)
		if err != nil {
			t.Errorf("[case %d] get err = %v", i, err)

			_ = e.Fin()
			continue
		}

		if !reflect.DeepEqual(res, tar) {
			t.Errorf("[case %d] res = %v, tar = %v", i, res, tar)

			_ = e.Fin()
			continue
		}

		if ttl == 0 {
			_ = e.Fin()
			continue
		}

		time.Sleep(time.Duration(ttl)*time.Second + time.Millisecond)
		if _, err = e.Get(k); err == nil {
			t.Errorf("[case %d] key should be expired", i)

			_ = e.Fin()
			continue
		}

		if !errors.Is(err, cache.ErrCacheMs) {
			t.Errorf("[case %d] get err = %v", i, err)

			_ = e.Fin()
			continue
		}

		if err := e.Fin(); err != nil {
			t.Errorf("[case %d] fin err = %v", i, err)
		}
	}
}

func TestExp_Set(t *testing.T) {
	cases := [][]any{
		{0, 0, "k", []byte("v0")}, {0, 0, "k", []byte("v0"), []byte("v1")},
		{1, 0, "k", []byte("v0")}, {1, 0, "k", []byte("v0"), []byte("v1")},
		{1, 1, "k", []byte("v0")}, {1, 1, "k", []byte("v0"), []byte("v1")},
	}

L:
	for i, cas := range cases {
		e, err := New(
			WithTtl(time.Duration(cas[0].(int))*time.Second),
			WithSpan(time.Duration(cas[1].(int))*time.Second),
		)

		if err != nil {
			t.Errorf("[case %d] new err = %v", i, err)
			continue
		}

		l, k := len(cas)-1, cas[2].(string)
		var tar []byte

		for j := 3; j <= l; j++ {
			v := cas[j].([]byte)
			if err := e.Set(k, v); err != nil {
				t.Errorf("[case %d] set err = %v", i, err)

				_ = e.Fin()
				continue L
			}

			if j == l {
				tar = v
			}
		}

		res, err := e.Get(k)
		if err != nil {
			t.Errorf("[case %d] get err = %v", i, err)

			_ = e.Fin()
			continue
		}

		if !reflect.DeepEqual(res, tar) {
			t.Errorf("[case %d] res = %v, tar = %v", i, res, tar)

			_ = e.Fin()
			continue
		}

		if err := e.Fin(); err != nil {
			t.Errorf("[case %d] fin err = %v", i, err)
		}
	}
}

func TestExp_Del(t *testing.T) {
	cases := [][4]any{{0, 0, "k", []byte("v")}, {1, 0, "k", []byte("v")}, {1, 1, "k", []byte("v")}}
	for i, cas := range cases {
		e, err := New(
			WithTtl(time.Duration(cas[0].(int))*time.Second),
			WithSpan(time.Duration(cas[1].(int))*time.Second),
		)

		if err != nil {
			t.Errorf("[case %d] new err = %v", i, err)
			continue
		}

		k := cas[2].(string)
		if err := e.Set(k, cas[3].([]byte)); err != nil {
			t.Errorf("[case %d] set err = %v", i, err)

			_ = e.Fin()
			continue
		}

		if err := e.Del(k); err != nil {
			t.Errorf("[case %d] del err = %v", i, err)

			_ = e.Fin()
			continue
		}

		if _, err = e.Get(k); err == nil {
			t.Errorf("[case %d] key should be deleted", i)

			_ = e.Fin()
			continue
		}

		if !errors.Is(err, cache.ErrCacheMs) {
			t.Errorf("[case %d] get err = %v", i, err)

			_ = e.Fin()
			continue
		}

		if err := e.Fin(); err != nil {
			t.Errorf("[case %d] fin err = %v", i, err)
		}
	}
}
