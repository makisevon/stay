package lru

type Mem int64

const (
	MemB Mem = 1
	MemK     = MemB << 10
	MemM     = MemK << 10
	MemG     = MemM << 10
	MemT     = MemG << 10
)
