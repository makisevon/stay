package cache

import "errors"

type Cache interface {
	Get(string) ([]byte, error)
	Set(string, []byte) error
	Del(string) error
}

var ErrCacheMs = errors.New("cache miss")
