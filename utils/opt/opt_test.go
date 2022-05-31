package opt

import (
	"reflect"
	"testing"
)

type testOpts struct {
	b bool
	i int
}

func newTestOpts(b bool, i int) *testOpts {
	return &testOpts{
		b: b,
		i: i,
	}
}

func TestMt(t *testing.T) {
	cases := []*testOpts{
		newTestOpts(false, 0), newTestOpts(false, 1), newTestOpts(true, 0), newTestOpts(true, 1),
	}

	for i, tar := range cases {
		res := &*tar
		if err := Mt[testOpts]()(res); err != nil {
			t.Errorf("[case %d] opt err = %v", i, err)
			continue
		}

		if !reflect.DeepEqual(res, tar) {
			t.Errorf("[case %d] res = %#v, tar = %#v", i, res, tar)
		}
	}
}

func TestApply(t *testing.T) {
	cases := [][2]any{
		{[]Opt[testOpts]{testWithB(true)}, newTestOpts(true, 0)},
		{[]Opt[testOpts]{testWithI(1)}, newTestOpts(false, 1)},

		{[]Opt[testOpts]{testWithB(false), testWithI(1)}, newTestOpts(false, 1)},
		{[]Opt[testOpts]{testWithB(true), testWithI(0)}, newTestOpts(true, 0)},
		{[]Opt[testOpts]{testWithB(true), testWithI(1)}, newTestOpts(true, 1)},
	}

	for i, cas := range cases {
		res := newTestOpts(false, 0)
		if err := Apply(res, cas[0].([]Opt[testOpts])); err != nil {
			t.Errorf("[case %d] apply err = %v", i, err)
			continue
		}

		if tar := cas[1]; !reflect.DeepEqual(res, tar) {
			t.Errorf("[case %d] res = %#v, tar = %#v", i, res, tar)
		}
	}
}

func testWithB(b bool) Opt[testOpts] {
	return func(o *testOpts) error {
		o.b = b
		return nil
	}
}

func testWithI(i int) Opt[testOpts] {
	return func(o *testOpts) error {
		o.i = i
		return nil
	}
}
