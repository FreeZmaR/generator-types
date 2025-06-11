package types

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRequiredSliceRule(t *testing.T) {
	type testCase struct {
		slice       Slice[int]
		expectError assert.ErrorAssertionFunc
	}

	tt := map[string]testCase{
		"Provided": {
			slice:       NewSlice[int](WithSliceIsProvided[int](true), WithSliceValue([]int{1})),
			expectError: assert.NoError,
		},
		"NotProvided": {
			slice:       NewSlice[int](WithSliceValue([]int{1})),
			expectError: assert.Error,
		},
		"Nil value": {
			slice:       NewSlice[int](WithSliceIsProvided[int](true)),
			expectError: assert.Error,
		},
		"Zero value": {
			slice:       NewSlice[int](WithSliceIsProvided[int](true), WithSliceValue([]int{})),
			expectError: assert.Error,
		},
	}

	rule := RequiredSliceRule[int]()

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			err := tc.slice.Validate(rule)
			tc.expectError(t, err)
		})
	}
}

func TestSliceSizeRule(t *testing.T) {
	type testCase struct {
		slice       Slice[int]
		ruleSize    int
		expectError assert.ErrorAssertionFunc
	}

	tt := map[string]testCase{
		"Values equal rule size": {
			slice:       NewSlice[int](WithSliceIsProvided[int](true), WithSliceValue([]int{1})),
			ruleSize:    1,
			expectError: assert.NoError,
		},
		"Values not equal rule size": {
			slice:       NewSlice[int](WithSliceIsProvided[int](true), WithSliceValue([]int{1})),
			ruleSize:    2,
			expectError: assert.Error,
		},
		"Nil value and rule size 1": {
			slice:       NewSlice[int](WithSliceIsProvided[int](true)),
			ruleSize:    1,
			expectError: assert.Error,
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			err := tc.slice.Validate(SliceSizeRule[int](tc.ruleSize))
			tc.expectError(t, err)
		})
	}
}

func TestSliceSizeGTE(t *testing.T) {
	type testCase struct {
		slice       Slice[int]
		ruleSize    int
		expectError assert.ErrorAssertionFunc
	}

	tt := map[string]testCase{
		"Values equal rule size": {
			slice:       NewSlice[int](WithSliceIsProvided[int](true), WithSliceValue([]int{1})),
			ruleSize:    1,
			expectError: assert.NoError,
		},
		"Values not equal rule size": {
			slice:       NewSlice[int](WithSliceIsProvided[int](true), WithSliceValue([]int{1})),
			ruleSize:    2,
			expectError: assert.Error,
		},
		"Nil value and rule size 1": {
			slice:       NewSlice[int](WithSliceIsProvided[int](true)),
			ruleSize:    1,
			expectError: assert.Error,
		},
		"Nil value and rule size 0": {
			slice:       NewSlice[int](WithSliceIsProvided[int](true)),
			ruleSize:    0,
			expectError: assert.NoError,
		},
		"2 values and rule size 1": {
			slice:       NewSlice[int](WithSliceIsProvided[int](true), WithSliceValue([]int{1})),
			ruleSize:    1,
			expectError: assert.NoError,
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			err := tc.slice.Validate(SliceSizeGTE[int](tc.ruleSize))
			tc.expectError(t, err)
		})
	}
}
