package net

import (
	"testing"
	"time"
)

func TestCtx(t *testing.T) {
	cases := []time.Duration{-1, 0, time.Second}
	for i, tmo := range cases {
		func() {
			defer func() {
				if err := recover(); err != nil {
					t.Errorf("[case %d] cxl err = %v", i, err)
				}
			}()

			_, cxl := Ctx(tmo)
			cxl()
		}()
	}
}

func TestChkAddr(t *testing.T) {
	cases := []string{"", ":", ":0", "0.0.0.0", "[::]", "0.0.0.0:0", "[::]:0"}
	for i, a := range cases {
		if err := ChkAddr(a); err != nil {
			t.Errorf("[case %d] chk_addr err = %v", i, err)
		}
	}
}
