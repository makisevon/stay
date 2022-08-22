package chash

import (
	"sort"
	"strconv"

	"stay/utils/opt"
)

type CHash struct {
	keys  map[int]string
	idxes []int

	*Opts
}

func New(os ...opt.Opt[Opts]) (*CHash, error) {
	o := newOpts()
	if err := opt.Apply(o, os); err != nil {
		return nil, err
	}

	h := &CHash{
		keys: make(map[int]string),

		Opts: o,
	}

	return h, nil
}

func (h *CHash) Add(ks ...string) {
	if len(ks) == 0 {
		return
	}

L:
	for _, k := range ks {
		for i := 0; i < h.rep; i++ {
			idx := h.idx(k, i)
			if _, ok := h.keys[idx]; ok {
				continue L
			}

			h.keys[idx] = k
			h.idxes = append(h.idxes, idx)
		}
	}

	sort.Ints(h.idxes)
}

func (h *CHash) Get(k string) string {
	l := len(h.idxes)
	if l == 0 {
		return ""
	}

	return h.keys[h.idxes[sort.SearchInts(h.idxes, h.hash(k))%l]]
}

func (h *CHash) Rm(ks ...string) {
	if len(ks) == 0 || len(h.idxes) == 0 {
		return
	}

L:
	for _, k := range ks {
		for i := 0; i < h.rep; i++ {
			idx := h.idx(k, i)
			if _, ok := h.keys[idx]; !ok {
				continue L
			}

			delete(h.keys, idx)

			j := sort.SearchInts(h.idxes, idx)
			h.idxes = append(h.idxes[:j], h.idxes[j+1:]...)
		}
	}

	l := len(h.idxes)
	h.idxes = h.idxes[:l:l]
}

func (h *CHash) idx(k string, i int) int {
	return h.hash(k + "-" + strconv.Itoa(i))
}
