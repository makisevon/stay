package cache

import (
	"errors"
	"testing"
)

func TestShell_Get(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Errorf("get err = %v", err)
		}
	}()

	if _, err := WithShell(Cache(nil)).Get(""); !errors.Is(err, ErrMtKey) {
		t.Errorf("get err = %v", err)
	}
}

func TestShell_Set(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Errorf("set err = %v", err)
		}
	}()

	s := WithShell(Cache(nil))
	cases := [][3]any{{"", []byte(nil), ErrMtKey}, {"k", []byte(nil), ErrNilVal}, {"k", []byte(""), ErrNilVal}}

	for i, cas := range cases {
		if err := s.Set(cas[0].(string), cas[1].([]byte)); !errors.Is(err, cas[2].(error)) {
			t.Errorf("[case %d] get err = %v", i, err)
		}
	}
}

func TestShell_Del(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Errorf("del err = %v", err)
		}
	}()

	if err := WithShell(Cache(nil)).Del(""); !errors.Is(err, ErrMtKey) {
		t.Errorf("del err = %v", err)
	}
}
