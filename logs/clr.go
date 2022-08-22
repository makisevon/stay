package logs

type clr []byte

var off = []byte("\033[0m")

func (c clr) dye(s string) string {
	return string(append(append(c, []byte(s)...), off...))
}

var (
	clrR clr = []byte("\033[1;31m")
	clrG clr = []byte("\033[1;32m")
	clrY clr = []byte("\033[1;33m")
	clrB clr = []byte("\033[1;34m")
	clrC clr = []byte("\033[1;36m")
)
