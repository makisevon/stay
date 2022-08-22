package svr

import (
	"context"
	"errors"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"

	"stay/cache/mem/raw"
	"stay/chash"
	"stay/logs"
	"stay/utils/net"
)

func TestSvr(t *testing.T) {
	if err := logs.Init(); err != nil {
		t.Errorf("logs init err = %v", err)
		return
	}

	c, cxl := net.Ctx(time.Second)
	defer cxl()

	g, c := errgroup.WithContext(c)
	defer func() {
		if err := g.Wait(); err != nil && !errors.Is(err, context.DeadlineExceeded) {
			t.Errorf("svr running err = %v", err)
		}
	}()

	as := []string{":8081", ":8082", ":8083"}

	h := ":8080"
	hs := []string{h, "", ""}

	for i, a := range as {
		s, err := New(raw.New(), a, WithMode(gin.TestMode), WithHttp(hs[i]), WithHos(chash.WithHash(testHash)))
		if err != nil {
			t.Errorf("[svr %s] new err = %v", a, err)
			return
		}

		if err := s.Reg(as...); err != nil {
			t.Errorf("[svr %s] reg err = %v", a, err)
			return
		}

		g.Go(func() error {
			e := make(chan error)
			go func() {
				e <- s.Run()
			}()

			select {
			case <-c.Done():
				return c.Err()

			case err := <-e:
				return err
			}
		})
	}

	p := "http://localhost" + h
	cases := []string{"80809", "80812", "80813", "80822", "80823", "80832", "80833"}

	for i, kv := range cases {
		v := url.Values{
			_val: []string{kv},
		}

		if _, err := http.PostForm(p+"/set/"+kv, v); err != nil {
			t.Errorf("[case %d] set err = %v", i, err)
			continue
		}

		if _, err := http.Get(p + "/get/" + kv); err != nil {
			t.Errorf("[case %d] get err = %v", i, err)
			continue
		}

		if _, err := http.PostForm(p+"/del/"+kv, nil); err != nil {
			t.Errorf("[case %d] del err = %v", i, err)
		}
	}
}

func testHash(s string) int {
	v, _ := strconv.Atoi(strings.Replace(strings.Replace(s, ":", "", 1), "-", "", 1))
	return v
}
