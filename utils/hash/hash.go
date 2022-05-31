package hash

type Hash func(string) int

type StdHash[T uint32 | uint64] func([]byte) T

func Wrap[T uint32 | uint64](h StdHash[T]) Hash {
	if h == nil {
		return nil
	}

	return func(s string) int {
		return int(h([]byte(s)))
	}
}
