package exp

import (
	"errors"
	"sync"
	"time"

	"stay/cache"
	"stay/utils/opt"
)

type Exp struct {
	cache map[string]*ent

	keys map[int64][]string

	fst chan struct{}
	fed chan struct{}

	once sync.Once
	mtx  sync.Mutex

	*Opts
}

type ent struct {
	val []byte
	exp int64
}

func New(os ...opt.Opt[Opts]) (*Exp, error) {
	o := newOpts()
	if err := opt.Apply(o, os); err != nil {
		return nil, err
	}

	e := &Exp{
		cache: make(map[string]*ent),

		Opts: o,
	}

	if e.span > 0 {
		e.keys, e.fst, e.fed = make(map[int64][]string), make(chan struct{}), make(chan struct{})
		go e.run(time.Now().Unix())
	}

	return e, nil
}

func (e *Exp) run(t int64) {
	ts := make(chan int64, e.proc)
	defer close(ts)

	go func() {
		for t := range ts {
			ks, ok := e.keys[t]
			if !ok {
				continue
			}

			e.mtx.Lock()

			for _, k := range ks {
				if v, ok := e.cache[k]; ok && v.exp == t {
					delete(e.cache, k)
				}
			}

			delete(e.keys, t)

			e.mtx.Unlock()
		}

		e.fed <- struct{}{}
	}()

	tk := time.NewTicker(e.span)
	defer tk.Stop()

	s := int(e.span / time.Second)

	for {
		select {
		case <-e.fst:
			return

		case <-tk.C:
			for i := 0; i < s; i++ {
				t++
				ts <- t
			}
		}
	}
}

var ErrLtFin = errors.New("lifetime of expiring cache is finished")

func (e *Exp) Fin() error {
	err := ErrLtFin
	e.once.Do(func() {
		err = nil
		if e.span == 0 {
			return
		}

		defer func() {
			close(e.fst)
			close(e.fed)
		}()

		e.fst <- struct{}{}
		<-e.fed
	})

	return err
}

var _ cache.Cache = (*Exp)(nil)

func (e *Exp) Get(k string) ([]byte, error) {
	e.mtx.Lock()
	defer e.mtx.Unlock()

	ent, ok := e.cache[k]
	if !ok {
		return nil, cache.ErrCacheMs
	}

	if ent.exp > time.Now().Unix() {
		return ent.val, nil
	}

	delete(e.cache, k)
	return nil, cache.ErrCacheMs
}

func (e *Exp) Set(k string, v []byte) error {
	e.mtx.Lock()
	defer e.mtx.Unlock()

	ent := &ent{
		val: v,
	}

	ent.exp = time.Now().Add(e.ttl).Unix()
	if e.span > 0 {
		e.keys[ent.exp] = append(e.keys[ent.exp], k)
	}

	e.cache[k] = ent
	return nil
}

func (e *Exp) Del(k string) error {
	e.mtx.Lock()
	defer e.mtx.Unlock()

	delete(e.cache, k)
	return nil
}
