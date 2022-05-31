package net

import (
	"context"
	"net/url"
	"time"

	"google.golang.org/grpc/resolver"
)

func Ctx(t time.Duration) (context.Context, context.CancelFunc) {
	if t <= 0 {
		return context.Background(), func() {}
	}

	return context.WithTimeout(context.Background(), t)
}

func ChkAddr(a string) error {
	if _, err := url.Parse(a); err == nil {
		return nil
	}

	_, err := url.Parse(resolver.GetDefaultScheme() + ":///" + a)
	return err
}
