package types

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRequiredStructRule(t *testing.T) {
	type Data struct{}

	type testCase struct {
		input       Struct[Data]
		expectError assert.ErrorAssertionFunc
	}

	tt := map[string]testCase{
		"Provided": {
			input:       NewStruct[Data](WithStructIsProvided[Data](true)),
			expectError: assert.NoError,
		},
		"NotProvided": {
			input:       NewStruct[Data](),
			expectError: assert.Error,
		},
	}

	rule := RequiredStructRule[Data]()

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			err := tc.input.Validate(rule)
			tc.expectError(t, err)
		})
	}
}
