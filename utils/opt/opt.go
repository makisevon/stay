package opt

type Opt[T any] func(*T) error

func Mt[T any]() Opt[T] {
	return func(*T) error {
		return nil
	}
}

func Apply[T any](o *T, os []Opt[T]) error {
	if o == nil {
		return nil
	}

	for _, opt := range os {
		if opt == nil {
			continue
		}

		if err := opt(o); err != nil {
			return err
		}
	}

	return nil
}
