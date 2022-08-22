package hash

import (
	"hash/crc32"
	"testing"
)

func TestWrap(t *testing.T) {
	hash := Wrap(crc32.ChecksumIEEE)
	cases := []string{"", "k", "v"}

	for i, s := range cases {
		if res, tar := hash(s), int(crc32.ChecksumIEEE([]byte(s))); res != tar {
			t.Errorf("[case %d] res = %d, tar = %d", i, res, tar)
		}
	}
}
