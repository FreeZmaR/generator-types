package main

import (
	"fmt"
)

type typeWithoutValue interface {
	IsProvided() bool
	Tag() string
}

type Rule[T Types] interface {
	Check(val Type[T]) error
}

type RequiredRule struct{}

func (r RequiredRule) Check(t typeWithoutValue) error {
	if !t.IsProvided() {
		return fmt.Errorf("%s: not provided", t.Tag())
	}

	return nil
}

type EqualRule[T Types] struct {
	vals []T
}

func NewEqualRule[T Types](vals ...T) *EqualRule[T] {
	return &EqualRule[T]{vals: vals}
}

func (r EqualRule[T]) Check(t Type[T]) error {
	if !t.IsProvided() {
		return nil
	}

	for i := range r.vals {
		if r.vals[i] == t.Value() {
			return nil
		}
	}

	return fmt.Errorf("%s: not equal %v", t.Tag(), r.vals)
}

type NotEqualRule[T Types] struct {
	vals []T
}

func NewNotEqualRule[T Types](vals ...T) *NotEqualRule[T] {
	return &NotEqualRule[T]{vals: vals}
}

func (r NotEqualRule[T]) Check(t Type[T]) error {
	if !t.IsProvided() {
		return nil
	}

	for i := range r.vals {
		if r.vals[i] == t.Value() {
			return fmt.Errorf("%s: equal %v", t.Tag(), r.vals)
		}
	}

	return nil
}

type GTERule[T int | float64] struct {
	val T
}

func NewGTERule[T int | float64](val T) GTERule[T] {
	return GTERule[T]{val: val}
}

func (r GTERule[T]) Check(t Type[T]) error {
	if !t.IsProvided() || t.Value() >= r.val {
		return nil
	}

	return fmt.Errorf("%s: not greater than or equal %v", t.Tag(), r.val)
}

type GTRule[T int | float64] struct {
	val T
}

func NewGTRule[T int | float64](val T) GTRule[T] {
	return GTRule[T]{val: val}
}

func (r GTRule[T]) Check(t Type[T]) error {
	if !t.IsProvided() || t.Value() > r.val {
		return nil
	}

	return fmt.Errorf("%s: not greater than %v", t.Tag(), r.val)
}

type LTERule[T int | float64] struct {
	val T
}

func NewLTERule[T int | float64](val T) LTERule[T] {
	return LTERule[T]{val: val}
}

func (r LTERule[T]) Check(t Type[T]) error {
	if !t.IsProvided() || t.Value() <= r.val {
		return nil
	}

	return fmt.Errorf("%s: not less than or equal %v", t.Tag(), r.val)
}

type LTRule[T int | float64] struct {
	val T
}

func NewLTRule[T int | float64](val T) LTRule[T] {
	return LTRule[T]{val: val}
}

func (r LTRule[T]) Check(t Type[T]) error {
	if !t.IsProvided() || t.Value() < r.val {
		return nil
	}

	return fmt.Errorf("%s: not less than %v", t.Tag(), r.val)
}
