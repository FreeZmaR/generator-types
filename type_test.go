package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type typeValidationTestCase[T any] struct {
	t     Type[T]
	rules []Rule[T]
	error assert.ErrorAssertionFunc
}

func TestTypeValidationInt(t *testing.T) {
	tt := map[string]typeValidationTestCase[int]{
		"Without tules": {
			t:     NewType(true, 1, "test", B2I),
			rules: []Rule[int]{},
			error: assert.NoError,
		},
		"Success all rules": {
			t: NewType(true, 1, "test", B2I),
			rules: []Rule[int]{
				RequiredRule[int](),
				EqualRule(1),
				NotEqualRule(2),
				GTERule(1),
				GTRule(0),
				LTERule(1),
				LTERule(2),
			},
			error: assert.NoError,
		},
		"Failed first rule": {
			t: NewType(false, 1, "test", B2I),
			rules: []Rule[int]{
				RequiredRule[int](),
				EqualRule(1),
				NotEqualRule(2),
				GTERule(1),
				GTRule(0),
				LTERule(1),
				LTERule(2),
			},
			error: assert.Error,
		},
		"Failed middle rule": {
			t: NewType(true, 1, "test", B2I),
			rules: []Rule[int]{
				RequiredRule[int](),
				EqualRule(1),
				NotEqualRule(2),
				GTERule(1),
				GTRule(2),
				LTERule(1),
				LTERule(2),
			},
			error: assert.Error,
		},
		"Failed last rule": {
			t: NewType(true, 1, "test", B2I),
			rules: []Rule[int]{
				RequiredRule[int](),
				EqualRule(1),
				NotEqualRule(2),
				GTERule(1),
				GTRule(0),
				LTERule(1),
				LTERule(0),
			},
			error: assert.Error,
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			err := tc.t.Validate(tc.rules...)
			tc.error(t, err)
		})
	}
}

func TestTypeValidationFloat64(t *testing.T) {
	tt := map[string]typeValidationTestCase[float64]{
		"Without rules": {
			t:     NewType(true, 1.0, "test", B2Float64),
			rules: []Rule[float64]{},
			error: assert.NoError,
		},
		"Success all rules": {
			t: NewType(true, 1.0, "test", B2Float64),
			rules: []Rule[float64]{
				RequiredRule[float64](),
				EqualRule(1.0),
				NotEqualRule(2.1),
				GTERule(0.6),
				GTRule(0.1),
				LTERule(1.1),
				LTERule(1.0),
			},
			error: assert.NoError,
		},
		"Failed middle rule": {
			t: NewType(true, 1.1, "test", B2Float64),
			rules: []Rule[float64]{
				RequiredRule[float64](),
				EqualRule(1.1),
				NotEqualRule(2.0),
				GTERule(1.2),
				GTRule(2.0),
				LTERule(1.0),
				LTERule(2.0),
			},
			error: assert.Error,
		},
		"Failed last rule": {
			t: NewType(true, 1.0, "test", B2Float64),
			rules: []Rule[float64]{
				RequiredRule[float64](),
				EqualRule(1.0),
				NotEqualRule(2.0),
				GTERule(1.0),
				GTRule(0.0),
				LTERule(1.0),
				LTERule(0.6),
			},
			error: assert.Error,
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			err := tc.t.Validate(tc.rules...)
			tc.error(t, err)
		})
	}
}

func TestTypeValidationString(t *testing.T) {
	tt := map[string]typeValidationTestCase[string]{
		"Without rules": {
			t:     NewType(true, "test", "test", B2S),
			rules: []Rule[string]{},
			error: assert.NoError,
		},
		"Success all rules": {
			t: NewType(true, "test", "test", B2S),
			rules: []Rule[string]{
				RequiredRule[string](),
				EqualRule("test"),
				NotEqualRule("some", "think"),
			},
			error: assert.NoError,
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			err := tc.t.Validate(tc.rules...)
			tc.error(t, err)
		})
	}
}

func TestTypeValidationBool(t *testing.T) {
	tt := map[string]typeValidationTestCase[bool]{
		"Without rules": {
			t:     NewType(true, true, "test", B2Bool),
			rules: []Rule[bool]{},
			error: assert.NoError,
		},
		"Success all rules": {
			t: NewType(true, true, "test", B2Bool),
			rules: []Rule[bool]{
				RequiredRule[bool](),
				EqualRule(true),
				NotEqualRule(false),
			},
			error: assert.NoError,
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			err := tc.t.Validate(tc.rules...)
			tc.error(t, err)
		})
	}
}
