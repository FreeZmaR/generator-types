package types

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// TEST Required

func TestRequiredRule_Int(t *testing.T) {
	tp := NewType(WithIsProvided[int](true))
	err := RequiredRule[int]()(&tp)
	assert.NoError(t, err)

	tp = NewType(WithIsProvided[int](false), WithTag[int]("test"))
	err = RequiredRule[int]()(&tp)
	assert.ErrorContains(t, err, "test: not provided")
}

func TestRequiredRule_Float64(t *testing.T) {
	tp := NewType(WithIsProvided[float64](true))
	err := RequiredRule[float64]()(&tp)
	assert.NoError(t, err)

	tp = NewType(WithIsProvided[float64](false), WithTag[float64]("test"))
	err = RequiredRule[float64]()(&tp)
	assert.ErrorContains(t, err, "test: not provided")
}

func TestRequiredRule_String(t *testing.T) {
	tp := NewType(WithIsProvided[string](true))
	err := RequiredRule[string]()(&tp)
	assert.NoError(t, err)

	tp = NewType(WithIsProvided[string](false), WithTag[string]("test"))
	err = RequiredRule[string]()(&tp)
	assert.ErrorContains(t, err, "test: not provided")
}

func TestRequiredRule_Bool(t *testing.T) {
	tp := NewType(WithIsProvided[bool](true))
	err := RequiredRule[bool]()(&tp)
	assert.NoError(t, err)

	tp = NewType(WithIsProvided[bool](false), WithTag[bool]("test"))
	err = RequiredRule[bool]()(&tp)
	assert.ErrorContains(t, err, "test: not provided")
}

// TEST Equal

func TestEqualRule_Int(t *testing.T) {
	tp := NewType(WithIsProvided[int](true), WithValue(1), WithTag[int]("test"))
	err := EqualRule[int](1)(&tp)
	assert.NoError(t, err)

	err = EqualRule[int](1, 2, 3, 4, 5)(&tp)
	assert.NoError(t, err)

	err = EqualRule[int](0)(&tp)
	assert.ErrorContains(t, err, "test: not equal [0]")

	err = EqualRule[int](0, 2, 3, 4, 5)(&tp)
	assert.ErrorContains(t, err, "test: not equal [0 2 3 4 5]")
}

func TestEqualRule_Float64(t *testing.T) {
	tp := NewType(WithIsProvided[float64](true), WithValue(1.0), WithTag[float64]("test"))
	err := EqualRule[float64](1.0)(&tp)
	assert.NoError(t, err)

	err = EqualRule[float64](0)(&tp)
	assert.ErrorContains(t, err, "test: not equal [0]")

	err = EqualRule[float64](0, 2.0, 3.0)(&tp)
	assert.ErrorContains(t, err, "test: not equal [0 2 3]")

	tp = NewType(WithValue(1.0))
	err = EqualRule[float64](2)(&tp)
	assert.NoError(t, err)

	tp = NewType(WithIsProvided[float64](true), WithValue(1.1))
	err = EqualRule[float64](1.1, 2.0, 3.0)(&tp)
	assert.NoError(t, err)
}

func TestEqualRule_String(t *testing.T) {
	tp := NewType(WithIsProvided[string](true), WithValue("some"), WithTag[string]("test"))
	err := EqualRule[string]("some")(&tp)
	assert.NoError(t, err)

	err = EqualRule[string]("some", "other")(&tp)
	assert.NoError(t, err)

	err = EqualRule[string]("different")(&tp)
	assert.ErrorContains(t, err, "test: not equal [different]")

	err = EqualRule[string]("diff", "erent")(&tp)
	assert.ErrorContains(t, err, "test: not equal [diff erent]")

	tp = NewType(WithValue("some"))
	err = EqualRule[string]("other")(&tp)
	assert.NoError(t, err)
}

func TestEqualRule_Bool(t *testing.T) {
	tp := NewType(WithIsProvided[bool](true), WithValue(true), WithTag[bool]("test"))
	err := EqualRule[bool](true)(&tp)
	assert.NoError(t, err)

	err = EqualRule[bool](false)(&tp)
	assert.ErrorContains(t, err, "test: not equal [false]")

	tp = NewType(WithIsProvided[bool](true), WithValue(false), WithTag[bool]("test"))
	err = EqualRule[bool](true)(&tp)
	assert.ErrorContains(t, err, "test: not equal [true]")

	tp = NewType(WithValue(false))
	err = EqualRule[bool](false, true)(&tp)
	assert.NoError(t, err)

	err = EqualRule[bool](false)(&tp)
	assert.NoError(t, err)
}

// TEST NotEqual

func TestNotEqualRule_Int(t *testing.T) {
	tp := NewType(WithIsProvided[int](true), WithValue(1), WithTag[int]("test"))
	err := NotEqualRule[int](2)(&tp)
	assert.NoError(t, err)

	err = NotEqualRule[int](0, 2, 3)(&tp)
	assert.NoError(t, err)

	err = NotEqualRule[int](1)(&tp)
	assert.ErrorContains(t, err, "test: equal [1]")

	err = NotEqualRule[int](0, 1, 2)(&tp)
	assert.ErrorContains(t, err, "test: equal [0 1 2]")
}

func TestNotEqualRule_Float64(t *testing.T) {
	tp := NewType(WithIsProvided[float64](true), WithValue(1.0), WithTag[float64]("test"))
	err := NotEqualRule[float64](2.0)(&tp)
	assert.NoError(t, err)

	err = NotEqualRule[float64](0.5, 2.0)(&tp)
	assert.NoError(t, err)

	err = NotEqualRule[float64](1.0)(&tp)
	assert.ErrorContains(t, err, "test: equal [1]")

	err = NotEqualRule[float64](0.0, 1.0, 2.0)(&tp)
	assert.ErrorContains(t, err, "test: equal [0 1 2]")
}

func TestNotEqualRule_String(t *testing.T) {
	tp := NewType(WithIsProvided[string](true), WithValue("some"), WithTag[string]("test"))
	err := NotEqualRule[string]("other")(&tp)
	assert.NoError(t, err)

	err = NotEqualRule[string]("diff", "erent")(&tp)
	assert.NoError(t, err)

	err = NotEqualRule[string]("some")(&tp)
	assert.ErrorContains(t, err, "test: equal [some]")

	err = NotEqualRule[string]("any", "some", "thing")(&tp)
	assert.ErrorContains(t, err, "test: equal [any some thing]")
}

func TestNotEqualRule_Bool(t *testing.T) {
	tp := NewType(WithIsProvided[bool](true), WithValue(true), WithTag[bool]("test"))
	err := NotEqualRule[bool](false)(&tp)
	assert.NoError(t, err)

	err = NotEqualRule[bool](true)(&tp)
	assert.ErrorContains(t, err, "test: equal [true]")

	tp = NewType(WithValue(false))
	err = NotEqualRule[bool](true)(&tp)
	assert.NoError(t, err)

	tp = NewType(WithIsProvided[bool](true), WithValue(false), WithTag[bool]("test"))
	err = NotEqualRule[bool](false)(&tp)
	assert.ErrorContains(t, err, "test: equal [false]")
}

// TEST GTE

func TestGTERule_Int(t *testing.T) {
	tp := NewType(WithIsProvided[int](true), WithValue(5), WithTag[int]("test"))
	err := GTERule[int](5)(&tp)
	assert.NoError(t, err)

	err = GTERule[int](4)(&tp)
	assert.NoError(t, err)

	tp = NewType(WithValue(3))
	err = GTERule[int](5)(&tp)
	assert.NoError(t, err)

	tp = NewType(WithIsProvided[int](true), WithValue(3), WithTag[int]("test"))
	err = GTERule[int](5)(&tp)
	assert.ErrorContains(t, err, "test: not greater than or equal 5")
}

func TestGTERule_Float64(t *testing.T) {
	tp := NewType(WithIsProvided[float64](true), WithValue(5.0))
	err := GTERule[float64](5.0)(&tp)
	assert.NoError(t, err)

	err = GTERule[float64](4.9)(&tp)
	assert.NoError(t, err)

	tp = NewType(WithIsProvided[float64](true), WithValue(3.0), WithTag[float64]("test"))
	err = GTERule[float64](5.0)(&tp)
	assert.ErrorContains(t, err, "test: not greater than or equal 5")

	tp = NewType(WithValue(3.0))
	err = GTERule[float64](5.0)(&tp)
	assert.NoError(t, err)
}

// TEST GT

func TestGTRule_Int(t *testing.T) {
	tp := NewType(WithIsProvided[int](true), WithValue(6))
	err := GTRule[int](5)(&tp)
	assert.NoError(t, err)

	tp = NewType(WithIsProvided[int](true), WithValue(5), WithTag[int]("test"))
	err = GTRule[int](5)(&tp)
	assert.ErrorContains(t, err, "test: not greater than 5")

	tp = NewType(WithIsProvided[int](true), WithValue(4), WithTag[int]("test"))
	err = GTRule[int](5)(&tp)
	assert.ErrorContains(t, err, "test: not greater than 5")

	tp = NewType(WithValue(3))
	err = GTRule[int](5)(&tp)
	assert.NoError(t, err)
}

func TestGTRule_Float64(t *testing.T) {
	tp := NewType(WithIsProvided[float64](true), WithValue(5.1))
	err := GTRule[float64](5.0)(&tp)
	assert.NoError(t, err)

	tp = NewType(WithIsProvided[float64](true), WithValue(5.0), WithTag[float64]("test"))
	err = GTRule[float64](5.0)(&tp)
	assert.ErrorContains(t, err, "test: not greater than 5")

	tp = NewType(WithIsProvided[float64](true), WithValue(4.9), WithTag[float64]("test"))
	err = GTRule[float64](5.0)(&tp)
	assert.ErrorContains(t, err, "test: not greater than 5")

	tp = NewType(WithValue(4.9))
	err = GTRule[float64](5.0)(&tp)
	assert.NoError(t, err)
}

// TEST LTE

func TestLTERule_Int(t *testing.T) {
	tp := NewType(WithIsProvided[int](true), WithValue(5))
	err := LTERule[int](5)(&tp)
	assert.NoError(t, err)

	err = LTERule[int](6)(&tp)
	assert.NoError(t, err)

	tp = NewType(WithIsProvided[int](true), WithValue(7), WithTag[int]("test"))
	err = LTERule[int](5)(&tp)
	assert.ErrorContains(t, err, "test: not less than or equal 5")

	tp = NewType(WithValue(6))
	err = LTERule[int](5)(&tp)
	assert.NoError(t, err)
}

func TestLTERule_Float64(t *testing.T) {
	tp := NewType(WithIsProvided[float64](true), WithValue(5.0))
	err := LTERule[float64](5.0)(&tp)
	assert.NoError(t, err)

	err = LTERule[float64](5.1)(&tp)
	assert.NoError(t, err)

	tp = NewType(WithIsProvided[float64](true), WithValue(5.2), WithTag[float64]("test"))
	err = LTERule[float64](5.0)(&tp)
	assert.ErrorContains(t, err, "test: not less than or equal 5")

	tp = NewType(WithValue(5.2))
	err = LTERule[float64](5.0)(&tp)
	assert.NoError(t, err)
}

// TEST LT

func TestLTRule_Int(t *testing.T) {
	tp := NewType(WithIsProvided[int](true), WithValue(4))
	err := LTRule[int](5)(&tp)
	assert.NoError(t, err)

	tp = NewType(WithIsProvided[int](true), WithValue(5), WithTag[int]("test"))
	err = LTRule[int](5)(&tp)
	assert.ErrorContains(t, err, "test: not less than 5")

	tp = NewType(WithIsProvided[int](true), WithValue(6), WithTag[int]("test"))
	err = LTRule[int](5)(&tp)
	assert.ErrorContains(t, err, "test: not less than 5")

	tp = NewType(WithValue(5))
	err = LTRule[int](5)(&tp)
	assert.NoError(t, err)
}

func TestLTRule_Float64(t *testing.T) {
	tp := NewType(WithIsProvided[float64](true), WithValue(4.9))
	err := LTRule[float64](5.0)(&tp)
	assert.NoError(t, err)

	tp = NewType(WithIsProvided[float64](true), WithValue(5.0), WithTag[float64]("test"))
	err = LTRule[float64](5.0)(&tp)
	assert.ErrorContains(t, err, "test: not less than 5")

	tp = NewType(WithIsProvided[float64](true), WithValue(5.1), WithTag[float64]("test"))
	err = LTRule[float64](5.0)(&tp)
	assert.ErrorContains(t, err, "test: not less than 5")

	tp = NewType(WithValue(5.1))
	err = LTRule[float64](5.0)(&tp)
	assert.NoError(t, err)
}
