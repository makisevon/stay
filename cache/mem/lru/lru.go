package lru

import (
	"container/list"
	"sync"

	"stay/cache"
	"stay/utils/opt"
)

type Lru struct {
	list  *list.List
	cache map[string]*list.Element

	len, lim Mem

	mtx sync.Mutex
}

type ent struct {
	key string
	val []byte
}

func New(os ...opt.Opt[Opts]) (*Lru, error) {
	o := newOpts()
	if err := opt.Apply(o, os); err != nil {
		return nil, err
	}

	l := &Lru{
		list:  list.New(),
		cache: make(map[string]*list.Element),

		lim: o.lim,
	}

	return l, nil
}

var _ cache.Cache = (*Lru)(nil)

func (l *Lru) Get(k string) ([]byte, error) {
	l.mtx.Lock()
	defer l.mtx.Unlock()

	if e, ok := l.cache[k]; ok {
		l.list.MoveToBack(e)
		return e.Value.(*ent).val, nil
	}

	return nil, cache.ErrCacheMs
}

func (l *Lru) Set(k string, v []byte) error {
	l.mtx.Lock()
	defer l.mtx.Unlock()

	if e, ok := l.cache[k]; ok {
		l.list.MoveToBack(e)
		e.Value.(*ent).val = v
	} else {
		l.len += Mem(len(k)) + Mem(len(v))
		l.cache[k] = l.list.PushBack(&ent{
			key: k,
			val: v,
		})
	}

	for l.len > l.lim {
		e := l.list.Remove(l.list.Front()).(*ent)
		l.len -= Mem(len(e.key)) + Mem(len(e.val))
		delete(l.cache, e.key)
	}

	return nil
}

func (l *Lru) Del(k string) error {
	l.mtx.Lock()
	defer l.mtx.Unlock()

	if e, ok := l.cache[k]; ok {
		l.len -= Mem(len(k)) + Mem(len(l.list.Remove(e).(*ent).val))
		delete(l.cache, k)
	}

	return nil
}
