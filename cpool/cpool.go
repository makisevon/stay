package cpool

import (
	"strconv"
	"sync/atomic"

	"golang.org/x/sync/singleflight"
	"google.golang.org/grpc"

	"stay/utils/net"
	"stay/utils/opt"
)

type CPool struct {
	pool []*conn

	idx, cap uint64

	sfg singleflight.Group
}

func New(a string, os ...opt.Opt[Opts]) (*CPool, error) {
	if err := net.ChkAddr(a); err != nil {
		return nil, err
	}

	o := newOpts()
	if err := opt.Apply(o, os); err != nil {
		return nil, err
	}

	p := &CPool{
		pool: make([]*conn, o.cap),

		cap: uint64(o.cap),
	}

	for i := range p.pool {
		p.pool[i] = newConn(a, o.connOpts)
	}

	return p, nil
}

func (p *CPool) Get() (*grpc.ClientConn, error) {
	idx := int(atomic.AddUint64(&p.idx, 1) % p.cap)
	c, err, _ := p.sfg.Do(strconv.Itoa(idx), func() (any, error) {
		return p.pool[idx].get()
	})

	return c.(*grpc.ClientConn), err
}
