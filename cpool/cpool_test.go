package cpool

import (
	_net "net"
	"testing"
	"time"

	"google.golang.org/grpc"

	"stay/utils/net"
	"stay/utils/opt"
)

func TestNew(t *testing.T) {
	cases := [][]opt.Opt[Opts]{
		nil, {},

		{WithCap(1)}, {WithTmo(time.Second)}, {WithDos(grpc.EmptyDialOption{})},

		{WithCap(1), WithTmo(time.Second)}, {WithCap(1), WithDos(grpc.EmptyDialOption{})},
		{WithTmo(time.Second), WithDos(grpc.EmptyDialOption{})},

		{WithCap(1), WithTmo(time.Second), WithDos(grpc.EmptyDialOption{})},
	}

	for i, os := range cases {
		if _, err := New("", os...); err != nil {
			t.Errorf("[case %d] new err = %v", i, err)
		}
	}
}

func TestCPool_Get(t *testing.T) {
	addr := testBenchSvr(t)
	if addr == "" {
		return
	}

	cases := []opt.Opt[Opts]{nil, WithTmo(0), WithTmo(time.Second)}
	for i, o := range cases {
		p, err := New(addr, o)
		if err != nil {
			t.Errorf("[case %d] new err = %v", i, err)
			continue
		}

		if _, err := p.Get(); err != nil {
			t.Errorf("[case %d] get err = %v", i, err)
		}
	}
}

func BenchmarkCPool_Get(b *testing.B) {
	addr := testBenchSvr(b)
	if addr == "" {
		return
	}

	p, err := New(addr)
	if err != nil {
		b.Errorf("new err = %v", err)
		return
	}

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if _, err := p.Get(); err != nil {
				b.Errorf("get err = %v", err)
			}
		}
	})
}

func BenchmarkCPool_Get_Bl(b *testing.B) {
	addr := testBenchSvr(b)
	if addr == "" {
		return
	}

	opts := append(credOpts, tmoOpts...)

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			ctx, cxl := net.Ctx(tmo)
			if _, err := grpc.DialContext(ctx, addr, opts...); err != nil {
				b.Errorf("dial err = %v", err)
			}

			cxl()
		}
	})
}

func testBenchSvr(tb testing.TB) string {
	l, err := _net.Listen("tcp", "")
	if err != nil {
		tb.Errorf("listen err = %v", err)
		return ""
	}

	go func() {
		s := grpc.NewServer()
		defer s.Stop()

		if err := s.Serve(l); err != nil {
			tb.Errorf("serve err = %v", err)
		}
	}()

	return l.Addr().String()
}
