package types

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type typeValidationTestCase[T any] struct {
	t     Type[T]
	error assert.ErrorAssertionFunc
}

func TestTypeValidationInt(t *testing.T) {
	tt := map[string]typeValidationTestCase[int]{
		"Without rules": {
			t: NewType(
				WithIsProvided[int](true),
				WithValue[int](1),
				WithTag[int]("test"),
				WithCastFn[int](B2I),
			),
			error: assert.NoError,
		},
		"Success all rules": {
			t: NewType(
				WithIsProvided[int](true),
				WithValue[int](1),
				WithTag[int]("test"),
				WithCastFn[int](B2I),
				WithRules[int](
					RequiredRule[int](),
					EqualRule(1),
					NotEqualRule(2),
					GTERule(1),
					GTRule(0),
					LTERule(1),
					LTERule(2),
				),
			),
			error: assert.NoError,
		},
		"Failed first rule": {
			t: NewType(
				WithValue[int](1),
				WithTag[int]("test"),
				WithCastFn[int](B2I),
				WithRules[int](
					RequiredRule[int](),
					EqualRule(1),
					NotEqualRule(2),
					GTERule(1),
					GTRule(0),
					LTERule(1),
					LTERule(2),
				),
			),
			error: assert.Error,
		},
		"Failed middle rule": {
			t: NewType(
				WithIsProvided[int](true),
				WithValue[int](1),
				WithTag[int]("test"),
				WithCastFn[int](B2I),
				WithRules[int](
					RequiredRule[int](),
					EqualRule(1),
					NotEqualRule(2),
					GTERule(1),
					GTRule(2),
					LTERule(1),
					LTERule(2),
				),
			),
			error: assert.Error,
		},
		"Failed last rule": {
			t: NewType(
				WithIsProvided[int](true),
				WithValue[int](1),
				WithTag[int]("test"),
				WithCastFn[int](B2I),
				WithRules[int](
					RequiredRule[int](),
					EqualRule(1),
					NotEqualRule(2),
					GTERule(1),
					GTRule(0),
					LTERule(1),
					LTERule(0),
				),
			),
			error: assert.Error,
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			err := tc.t.Validate()
			tc.error(t, err)
		})
	}
}

func TestTypeValidationFloat64(t *testing.T) {
	tt := map[string]typeValidationTestCase[float64]{
		"Without rules": {
			t: NewType(
				WithIsProvided[float64](true),
				WithValue(1.0),
				WithTag[float64]("test"),
				WithCastFn[float64](B2Float64),
			),
			error: assert.NoError,
		},
		"Success all rules": {
			t: NewType(
				WithIsProvided[float64](true),
				WithValue(1.0),
				WithTag[float64]("test"),
				WithCastFn[float64](B2Float64),
				WithRules[float64](
					RequiredRule[float64](),
					EqualRule(1.0),
					NotEqualRule(2.1),
					GTERule(0.6),
					GTRule(0.1),
					LTERule(1.1),
					LTERule(1.0),
				),
			),
			error: assert.NoError,
		},
		"Failed middle rule": {
			t: NewType(
				WithIsProvided[float64](true),
				WithValue(1.1),
				WithTag[float64]("test"),
				WithCastFn[float64](B2Float64),
				WithRules[float64](
					RequiredRule[float64](),
					EqualRule(1.1),
					NotEqualRule(2.0),
					GTERule(1.2),
					GTRule(2.0),
					LTERule(1.0),
					LTERule(2.0),
				),
			),
			error: assert.Error,
		},
		"Failed last rule": {
			t: NewType(
				WithIsProvided[float64](true),
				WithValue(1.0),
				WithTag[float64]("test"),
				WithCastFn[float64](B2Float64),
				WithRules[float64](
					RequiredRule[float64](),
					EqualRule(1.0),
					NotEqualRule(2.0),
					GTERule(1.0),
					GTRule(0.0),
					LTERule(1.0),
					LTERule(0.6),
				),
			),
			error: assert.Error,
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			err := tc.t.Validate()
			tc.error(t, err)
		})
	}
}

func TestTypeValidationString(t *testing.T) {
	tt := map[string]typeValidationTestCase[string]{
		"Without rules": {
			t: NewType(
				WithIsProvided[string](true),
				WithValue("test"),
				WithTag[string]("test"),
				WithCastFn[string](B2S),
			),
			error: assert.NoError,
		},
		"Success all rules": {
			t: NewType(
				WithIsProvided[string](true),
				WithValue("test"),
				WithTag[string]("test"),
				WithCastFn[string](B2S),
				WithRules[string](
					RequiredRule[string](),
					EqualRule("test"),
					NotEqualRule("some", "think"),
				),
			),
			error: assert.NoError,
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			err := tc.t.Validate()
			tc.error(t, err)
		})
	}
}

func TestTypeValidationBool(t *testing.T) {
	tt := map[string]typeValidationTestCase[bool]{
		"Without rules": {
			t: NewType(
				WithIsProvided[bool](true),
				WithValue(true),
				WithTag[bool]("test"),
				WithCastFn[bool](B2Bool),
			),
			error: assert.NoError,
		},
		"Success all rules": {
			t: NewType(
				WithIsProvided[bool](true),
				WithValue(true),
				WithTag[bool]("test"),
				WithCastFn[bool](B2Bool),
				WithRules[bool](
					RequiredRule[bool](),
					EqualRule(true),
					NotEqualRule(false),
				),
			),
			error: assert.NoError,
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			err := tc.t.Validate()
			tc.error(t, err)
		})
	}
}
