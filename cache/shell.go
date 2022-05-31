package cache

import "errors"

type Shell struct {
	cache Cache
}

func WithShell(c Cache) *Shell {
	return &Shell{
		cache: c,
	}
}

var _ Cache = (*Shell)(nil)

var ErrMtKey = errors.New("empty cache key")

func (s *Shell) Get(k string) ([]byte, error) {
	if k == "" {
		return nil, ErrMtKey
	}

	v, err := s.cache.Get(k)
	return view(v), err
}

var ErrNilVal = errors.New("nil cache value")

func (s *Shell) Set(k string, v []byte) error {
	if k == "" {
		return ErrMtKey
	}

	if v = view(v); v == nil {
		return ErrNilVal
	}

	return s.cache.Set(k, v)
}

func view(bs []byte) []byte {
	l := len(bs)
	if l == 0 {
		return nil
	}

	r := make([]byte, l)
	if n := copy(r, bs); n > 0 {
		return r[:n:n]
	}

	return nil
}

func (s *Shell) Del(k string) error {
	if k == "" {
		return ErrMtKey
	}

	return s.cache.Del(k)
}
