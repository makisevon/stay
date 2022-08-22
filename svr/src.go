package svr

import "stay/cache"

type Src interface {
	Get(string) ([]byte, error)
}

type SrcFn func(string) ([]byte, error)

func (f SrcFn) Get(k string) ([]byte, error) {
	return f(k)
}

var src = SrcFn(func(string) ([]byte, error) {
	return nil, cache.ErrCacheMs
})
