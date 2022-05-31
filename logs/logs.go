package logs

import (
	"io"
	_os "os"
	"sync"

	"stay/utils/opt"
)

var (
	dpf = clrB.dye("Dbg") + "  "
	ipf = clrG.dye("Info") + " "
	wpf = clrY.dye("Warn") + " "
	epf = clrR.dye("Err") + "  "

	dlg *log
	ilg *log
	wlg *log
	elg *log

	mtx sync.Mutex
)

func Init(os ...opt.Opt[Opts]) error {
	o := newOpts()
	if err := opt.Apply(o, os); err != nil {
		return err
	}

	dw, iw, ww, ew := io.Discard, io.Writer(_os.Stdout), io.Writer(_os.Stdout), io.Writer(_os.Stderr)
	switch o.lv {
	case LvDbg:
		dw = _os.Stdout
	case LvWarn:
		dw, iw = io.Discard, io.Discard
	case LvErr:
		dw, iw, ww = io.Discard, io.Discard, io.Discard
	}

	mtx.Lock()
	defer mtx.Unlock()

	dlg, ilg = newLog(dw, dpf, o.flag), newLog(iw, ipf, o.flag)
	wlg, elg = newLog(ww, wpf, o.flag), newLog(ew, epf, o.flag)

	return nil
}

func Dbg(f string, vs ...any) {
	_ = dlg.op(1, f, vs...)
}

func Info(f string, vs ...any) {
	_ = ilg.op(1, f, vs...)
}

func Warn(f string, vs ...any) {
	_ = wlg.op(1, f, vs...)
}

func Err(f string, vs ...any) {
	_ = elg.op(1, f, vs...)
}
