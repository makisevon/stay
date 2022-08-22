package logs

type Lv int

const (
	LvInfo Lv = iota
	LvDbg
	LvWarn
	LvErr
)
