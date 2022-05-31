package logs

import (
	"errors"
	_log "log"

	"stay/utils/opt"
)

type Opts struct {
	lv Lv

	flag int
}

const flag = _log.Lshortfile | _log.LstdFlags

func newOpts() *Opts {
	return &Opts{
		lv: LvInfo,

		flag: flag,
	}
}

var ErrInvLv = errors.New("invalid log level")

func WithLv(l Lv) opt.Opt[Opts] {
	return func(o *Opts) error {
		switch l {
		case -1:
			o.lv = LvInfo
		case LvInfo, LvDbg, LvWarn, LvErr:
			o.lv = l

		default:
			return ErrInvLv
		}

		return nil
	}
}

var ErrInvFlag = errors.New("invalid log flag")

func WithFlag(f int) opt.Opt[Opts] {
	return func(o *Opts) error {
		switch {
		case o.flag < -1:
			return ErrInvFlag

		case o.flag == -1:
			o.flag = flag
		default:
			o.flag = f
		}

		return nil
	}
}
