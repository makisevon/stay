package main

import (
	"errors"
	"flag"
	"strings"

	"github.com/gin-gonic/gin"

	"stay/cache"
	"stay/cache/cbf"
	"stay/cache/mem/exp"
	"stay/cache/mem/lru"
	"stay/cache/mlv"
	"stay/cache/rds"
	"stay/logs"
	"stay/svr"
)

var (
	addr string
	http string

	peers string
)

func init() {
	if err := logs.Init(); err != nil {
		panic(err)
	}

	defer flag.Parse()

	flag.StringVar(&addr, "addr", "", "RPC server address")
	flag.StringVar(&http, "http", "", "HTTP server address")

	flag.StringVar(&peers, "peers", "", "RPC peer addresses")
}

func main() {
	if addr == "" || peers == "" {
		logs.Err("missing required args")
		return
	}

	c, err := newCache()
	if err != nil {
		logs.Err("new cache err = %v", err)
		return
	}

	s, err := svr.New(c, addr, svr.WithMode(gin.ReleaseMode), svr.WithHttp(http))
	if err != nil {
		logs.Err("new svr err = %v", err)
		return
	}

	if err := s.Reg(strings.Split(peers, ",")...); err != nil {
		logs.Err("reg peers err = %v", err)
		return
	}

	if err := s.Run(); err != nil {
		logs.Err("svr running err = %v", err)
	}
}

func newCache() (cache.Cache, error) {
	ts := flag.Args()

	l := len(ts)
	if l == 0 {
		return lru.New()
	}

	if l == 1 {
		return newCacheByTyp(ts[0])
	}

	cs := make([]cache.Cache, l)
	for i, t := range ts {
		var err error
		if cs[i], err = newCacheByTyp(t); err != nil {
			return nil, err
		}
	}

	m, err := mlv.New(cs...)
	if err != nil {
		return nil, err
	}

	c, err := cbf.WithCbf(m)
	if err != nil {
		return nil, err
	}

	return cache.WithShell(c), nil
}

var errInvTyp = errors.New("invalid cache type")

func newCacheByTyp(t string) (cache.Cache, error) {
	switch t {
	case "exp":
		return exp.New()
	case "lru":
		return lru.New()
	case "rds":
		return rds.New()
	}

	return nil, errInvTyp
}
