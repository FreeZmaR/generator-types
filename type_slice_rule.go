package types

import "fmt"

type SliceRule[T any] func(t *Slice[T]) error

func RequiredSliceRule[T any]() SliceRule[T] {
	return func(t *Slice[T]) error {
		if !t.isProvided || len(t.value) == 0 {
			return fmt.Errorf("%s: not provided or size of data is 0", t.tag)
		}

		return nil
	}
}

func SliceSizeRule[T any](size int) SliceRule[T] {
	return func(t *Slice[T]) error {
		if len(t.value) != size {
			return fmt.Errorf("%s: size of data is not  %d", t.tag, size)
		}

		return nil
	}
}

func SliceSizeGTE[T any](size int) SliceRule[T] {
	return func(t *Slice[T]) error {
		if len(t.value) < size {
			return fmt.Errorf("%s: size of data les than %d", t.tag, size)
		}

		return nil
	}
}
