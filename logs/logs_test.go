package logs

import "testing"

func TestLogs(t *testing.T) {
	cases := []Lv{LvDbg, LvInfo, LvWarn, LvErr}
	for i, l := range cases {
		if err := Init(WithLv(l)); err != nil {
			t.Errorf("[case %d] init err = %v", i, err)
			continue
		}

		Dbg("[case %d]", i)
		Info("[case %d]", i)
		Warn("[case %d]", i)
		Err("[case %d]", i)
	}
}
