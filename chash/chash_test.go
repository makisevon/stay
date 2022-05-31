package chash

import (
	"hash/crc32"
	"strconv"
	"strings"
	"testing"

	"stay/utils/opt"
)

func TestNew(t *testing.T) {
	cases := [][]opt.Opt[Opts]{
		nil, {},

		{WithRep(1)}, {WithStdHash(crc32.ChecksumIEEE)}, {WithRep(1), WithStdHash(crc32.ChecksumIEEE)},
	}

	for i, os := range cases {
		if _, err := New(os...); err != nil {
			t.Errorf("[case %d] new err = %v", i, err)
		}
	}
}

func TestCHash(t *testing.T) {
	cases := [][2]any{
		{
			1,
			[][3]any{
				{
					false, []string{"1"},
					[][2]string{
						{"09", "1"}, {"10", "1"}, {"11", "1"},
						{"19", "1"}, {"20", "1"}, {"21", "1"},
						{"29", "1"}, {"30", "1"}, {"31", "1"},
					},
				}, {
					false, []string{"2", "3"},
					[][2]string{
						{"09", "1"}, {"10", "1"}, {"11", "2"},
						{"19", "2"}, {"20", "2"}, {"21", "3"},
						{"29", "3"}, {"30", "3"}, {"31", "1"},
					},
				}, {
					true, []string{"2"},
					[][2]string{
						{"09", "1"}, {"10", "1"}, {"11", "3"},
						{"19", "3"}, {"20", "3"}, {"21", "3"},
						{"29", "3"}, {"30", "3"}, {"31", "1"},
					},
				}, {
					true, []string{"1", "3"},
					[][2]string{{"09"}, {"10"}, {"11"}, {"19"}, {"20"}, {"21"}, {"29"}, {"30"}, {"31"}},
				},
			},
		}, {
			2,
			[][3]any{
				{
					false, []string{"1"},
					[][2]string{
						{"09", "1"}, {"10", "1"}, {"11", "1"}, {"12", "1"},
						{"19", "1"}, {"20", "1"}, {"21", "1"}, {"22", "1"},
						{"29", "1"}, {"30", "1"}, {"31", "1"}, {"32", "1"},
					},
				}, {
					false, []string{"2", "3"},
					[][2]string{
						{"09", "1"}, {"10", "1"}, {"11", "1"}, {"12", "2"},
						{"19", "2"}, {"20", "2"}, {"21", "2"}, {"22", "3"},
						{"29", "3"}, {"30", "3"}, {"31", "3"}, {"32", "1"},
					},
				}, {
					true, []string{"2"},
					[][2]string{
						{"09", "1"}, {"10", "1"}, {"11", "1"}, {"12", "3"},
						{"19", "3"}, {"20", "3"}, {"21", "3"}, {"22", "3"},
						{"29", "3"}, {"30", "3"}, {"31", "3"}, {"32", "1"},
					},
				}, {
					true, []string{"1", "3"},
					[][2]string{
						{"09"}, {"10"}, {"11"}, {"12"},
						{"19"}, {"20"}, {"21"}, {"22"},
						{"29"}, {"30"}, {"31"}, {"32"},
					},
				},
			},
		},
	}

L:
	for i, cas := range cases {
		h, err := New(WithRep(cas[0].(int)), WithHash(testHash))
		if err != nil {
			t.Errorf("[case %d] new err = %v", i, err)
			continue
		}

		for j, op := range cas[1].([][3]any) {
			fn := h.Add
			if op[0].(bool) {
				fn = h.Rm
			}

			fn(op[1].([]string)...)
			for _, pair := range op[2].([][2]string) {
				key := pair[0]
				if res, tar := h.Get(key), pair[1]; res != tar {
					t.Errorf("[case %d] [oper %d] key = %s, res = %s, tar = %s", i, j, key, res, tar)
					continue L
				}
			}
		}
	}
}

func testHash(s string) int {
	v, _ := strconv.Atoi(strings.Replace(s, "-", "", 1))
	return v
}
