package types

import "fmt"

type StructRule[T any] func(t *Struct[T]) error

func RequiredStructRule[T any]() StructRule[T] {
	return func(t *Struct[T]) error {
		if !t.isProvided {
			return fmt.Errorf("%s: not provided", t.tag)
		}

		return nil
	}
}
