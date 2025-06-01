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
			slice:       NewSlice[int](true, []int{1}, "", nil),
			expectError: assert.NoError,
		},
		"NotProvided": {
			slice:       NewSlice[int](false, []int{1}, "", nil),
			expectError: assert.Error,
		},
		"Nil value": {
			slice:       NewSlice[int](true, nil, "", nil),
			expectError: assert.Error,
		},
		"Zero value": {
			slice:       NewSlice[int](true, []int{}, "", nil),
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
			slice:       NewSlice[int](true, []int{1}, "", nil),
			ruleSize:    1,
			expectError: assert.NoError,
		},
		"Values not equal rule size": {
			slice:       NewSlice[int](false, []int{1}, "", nil),
			ruleSize:    2,
			expectError: assert.Error,
		},
		"Nil value and rule size 1": {
			slice:       NewSlice[int](true, nil, "", nil),
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
			slice:       NewSlice[int](true, []int{1}, "", nil),
			ruleSize:    1,
			expectError: assert.NoError,
		},
		"Values not equal rule size": {
			slice:       NewSlice[int](false, []int{1}, "", nil),
			ruleSize:    2,
			expectError: assert.Error,
		},
		"Nil value and rule size 1": {
			slice:       NewSlice[int](true, nil, "", nil),
			ruleSize:    1,
			expectError: assert.Error,
		},
		"Nil value and rule size 0": {
			slice:       NewSlice[int](true, nil, "", nil),
			ruleSize:    0,
			expectError: assert.NoError,
		},
		"2 values and rule size 1": {
			slice:       NewSlice[int](true, []int{1, 2}, "", nil),
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
