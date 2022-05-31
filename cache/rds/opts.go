package rds

import (
	"errors"
	"time"

	"github.com/go-redis/redis/v8"

	"stay/utils/opt"
)

type Opts struct {
	ros *redis.Options

	ttl time.Duration
}

func newOpts() *Opts {
	return &Opts{
		ros: new(redis.Options),
	}
}

func WithRos(os *redis.Options) opt.Opt[Opts] {
	return func(o *Opts) error {
		if os == nil {
			o.ros = new(redis.Options)
		} else {
			o.ros = os
		}

		return nil
	}
}

var ErrInvTtl = errors.New("invalid TTL for Redis cache")

func WithTtl(t time.Duration) opt.Opt[Opts] {
	return func(o *Opts) error {
		if t < 0 {
			return ErrInvTtl
		}

		o.ttl = t
		return nil
	}
}
