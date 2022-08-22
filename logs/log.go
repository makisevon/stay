package logs

import (
	"fmt"
	"io"
	_log "log"
)

type log struct {
	log *_log.Logger

	pref string

	dsc bool
}

var pref = clrC.dye("[Stay]") + " "

func newLog(w io.Writer, p string, f int) *log {
	return &log{
		log: _log.New(w, pref, f),

		pref: p,

		dsc: w == io.Discard,
	}
}

func (l *log) op(d int, f string, vs ...any) error {
	if l.dsc {
		return nil
	}

	return l.log.Output(d+2, fmt.Sprintf(l.pref+f, vs...))
}
