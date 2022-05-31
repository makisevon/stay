package rds

import (
	"context"
	"errors"
	"time"

	"github.com/go-redis/redis/v8"

	"stay/cache"
	"stay/utils/opt"
)

type Rds struct {
	db *redis.Client

	ttl time.Duration
}

func New(os ...opt.Opt[Opts]) (*Rds, error) {
	o := newOpts()
	if err := opt.Apply(o, os); err != nil {
		return nil, err
	}

	db := redis.NewClient(o.ros)
	if _, err := db.Ping(context.Background()).Result(); err != nil {
		return nil, err
	}

	r := &Rds{
		db: db,

		ttl: o.ttl,
	}

	return r, nil
}

func (r *Rds) Fin() error {
	return r.db.Close()
}

var _ cache.Cache = (*Rds)(nil)

func (r *Rds) Get(k string) ([]byte, error) {
	if v, err := r.db.Get(context.Background(), k).Bytes(); !errors.Is(err, redis.Nil) {
		return v, err
	}

	return nil, cache.ErrCacheMs
}

func (r *Rds) Set(k string, v []byte) error {
	if _, err := r.db.Set(context.Background(), k, v, r.ttl).Result(); !errors.Is(err, redis.Nil) {
		return err
	}

	return cache.ErrCacheMs
}

func (r *Rds) Del(k string) error {
	if _, err := r.db.Del(context.Background(), k).Result(); !errors.Is(err, redis.Nil) {
		return err
	}

	return cache.ErrCacheMs
}
