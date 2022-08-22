package cpool

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
	"google.golang.org/grpc/credentials/insecure"

	"stay/utils/net"
)

type conn struct {
	conn *grpc.ClientConn

	addr string

	*connOpts
}

var (
	tmoOpts  = []grpc.DialOption{grpc.WithBlock(), grpc.WithReturnConnectionError()}
	credOpts = []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
)

func newConn(a string, o *connOpts) *conn {
	c := &conn{
		addr: a,

		connOpts: o,
	}

	if c.tmo > 0 {
		c.dos = append(tmoOpts, c.dos...)
	}

	c.dos = append(credOpts, c.dos...)
	return c
}

func (c *conn) get() (*grpc.ClientConn, error) {
	if c.conn != nil {
		s := c.conn.GetState()
		if s != connectivity.TransientFailure && s != connectivity.Shutdown {
			return c.conn, nil
		}
	}

	ctx, cxl := net.Ctx(c.tmo)
	defer cxl()

	var err error
	c.conn, err = grpc.DialContext(ctx, c.addr, c.dos...)
	return c.conn, err
}
