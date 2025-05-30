package types

import (
	"fmt"
)

type Rule[T any] func(t *Type[T]) error

func RequiredRule[T any]() Rule[T] {
	return func(t *Type[T]) error {
		if !t.IsProvided() {
			return fmt.Errorf("%s: not provided", t.Tag())
		}

		return nil
	}
}

func EqualRule[T comparable](vals ...T) Rule[T] {
	return func(t *Type[T]) error {
		if !t.IsProvided() {
			return nil
		}

		for i := range vals {
			if vals[i] == t.Value() {
				return nil
			}
		}

		return fmt.Errorf("%s: not equal %v", t.Tag(), vals)
	}
}

func NotEqualRule[T comparable](vals ...T) Rule[T] {
	return func(t *Type[T]) error {
		if !t.IsProvided() {
			return nil
		}

		for i := range vals {
			if vals[i] == t.Value() {
				return fmt.Errorf("%s: equal %v", t.Tag(), vals)
			}
		}

		return nil
	}
}

func GTERule[T int | float64](val T) Rule[T] {
	return func(t *Type[T]) error {
		if !t.IsProvided() || t.Value() >= val {
			return nil
		}

		return fmt.Errorf("%s: not greater than or equal %v", t.Tag(), val)
	}
}

func GTRule[T int | float64](val T) Rule[T] {
	return func(t *Type[T]) error {
		if !t.IsProvided() || t.Value() > val {
			return nil
		}

		return fmt.Errorf("%s: not greater than %v", t.Tag(), val)
	}
}

func LTERule[T int | float64](val T) Rule[T] {
	return func(t *Type[T]) error {
		if !t.IsProvided() || t.Value() <= val {
			return nil
		}

		return fmt.Errorf("%s: not less than or equal %v", t.Tag(), val)
	}
}

func LTRule[T int | float64](val T) Rule[T] {
	return func(t *Type[T]) error {
		if !t.IsProvided() || t.Value() < val {
			return nil
		}

		return fmt.Errorf("%s: not less than %v", t.Tag(), val)
	}
}
