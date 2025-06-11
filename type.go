package types

import (
	"fmt"
	"github.com/goccy/go-json"
)

type Type[T any] struct {
	isProvided bool
	value      T
	tag        string
	castFN     TypeCastFn[T]
	rules      []TypeRule[T]
}

type TypeCastFn[T any] func(b []byte) (T, error)

func NewType[T any](options ...TypeOption[T]) Type[T] {
	var t Type[T]

	for i := range options {
		options[i](&t)
	}

	return t
}

func (t *Type[T]) IsProvided() bool {
	return t.isProvided
}

func (t *Type[T]) Value() T {
	return t.value
}

func (t *Type[T]) Tag() string {
	return t.tag
}

func (t *Type[T]) SetTag(tag string) {
	t.tag = tag
}

func (t *Type[T]) String() string {
	return fmt.Sprintf("%s:%v", t.tag, t.value)
}

func (t *Type[T]) Validate(rules ...TypeRule[T]) error {
	if len(t.rules) != 0 {
		rules = append(t.rules, rules...)
	}

	for i := range rules {
		if err := rules[i](t); err != nil {
			return err
		}
	}

	return nil
}

func (t *Type[T]) UnmarshalJSON(b []byte) error {
	var err error

	t.isProvided = true
	t.value, err = t.cast(b)
	if err != nil {
		return err
	}

	return nil
}

func (t *Type[T]) cast(b []byte) (T, error) {
	var value T

	if t.castFN != nil {
		return t.castFN(b)
	}

	switch any(t.value).(type) {
	case int:
		val, err := B2I(b)
		if err != nil {
			return value, err
		}

		value = any(val).(T)
	case float64:
		val, err := B2Float64(b)
		if err != nil {
			return value, err
		}

		value = any(val).(T)
	case string:
		val, err := B2S(b)
		if err != nil {
			return value, err
		}

		value = any(val).(T)
	case bool:
		val, err := B2Bool(b)
		if err != nil {
			return value, err
		}

		value = any(val).(T)
	default:
		if err := json.Unmarshal(b, &value); err != nil {
			return value, err
		}
	}

	return value, nil
}
