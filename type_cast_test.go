package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestB2S(t *testing.T) {
	type testCase struct {
		input    []byte
		expected string
	}

	tt := map[string]testCase{
		"empty": {
			input:    []byte{},
			expected: "",
		},
		"nil": {},
		"simple string": {
			input:    []byte("string"),
			expected: "string",
		},
		"With quotes": {
			input:    []byte(`"some string"`),
			expected: "some string",
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			get, _ := B2S(tc.input)
			assert.Equal(t, tc.expected, get)
		})
	}
}

func TestB2I(t *testing.T) {
	type testCase struct {
		input    []byte
		expected int
		error    assert.ErrorAssertionFunc
	}

	tt := map[string]testCase{
		"empty": {
			input:    []byte{},
			expected: 0,
			error:    assert.NoError,
		},
		"nil": {
			expected: 0,
			error:    assert.NoError,
		},
		"base number": {
			input:    []byte("123"),
			expected: 123,
			error:    assert.NoError,
		},
		"number in quotes": {
			input:    []byte(`"123"`),
			expected: 123,
			error:    assert.NoError,
		},
		"invalid number": {
			input:    []byte("wrong"),
			expected: 0,
			error:    assert.Error,
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			get, err := B2I(tc.input)
			tc.error(t, err)
			assert.Equal(t, tc.expected, get)
		})
	}
}

func TestB2Bool(t *testing.T) {
	type testCase struct {
		input    []byte
		expected bool
		error    assert.ErrorAssertionFunc
	}

	tt := map[string]testCase{
		"empty": {
			input:    []byte{},
			expected: false,
			error:    assert.NoError,
		},
		"nil": {
			expected: false,
			error:    assert.NoError,
		},
		"true": {
			input:    []byte("true"),
			expected: true,
			error:    assert.NoError,
		},
		"false": {
			input:    []byte("false"),
			expected: false,
			error:    assert.NoError,
		},
		"true as number": {
			input:    []byte("1"),
			expected: true,
			error:    assert.NoError,
		},
		"false as number": {
			input:    []byte("0"),
			expected: false,
			error:    assert.NoError,
		},
		"true in quotes": {
			input:    []byte(`"true"`),
			expected: true,
			error:    assert.NoError,
		},
		"false in quoted": {
			input:    []byte(`"false"`),
			expected: false,
			error:    assert.NoError,
		},
		"true as number in quotes": {
			input:    []byte(`"1"`),
			expected: true,
			error:    assert.NoError,
		},
		"false as number in quotes": {
			input:    []byte(`"0"`),
			expected: false,
			error:    assert.NoError,
		},
		"invalid value": {
			input:    []byte("wrong"),
			expected: false,
			error:    assert.Error,
		},
		"invalid number": {
			input:    []byte("2"),
			expected: false,
			error:    assert.Error,
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			get, err := B2Bool(tc.input)
			tc.error(t, err)
			assert.Equal(t, tc.expected, get)
		})
	}
}

func TestB2Float64(t *testing.T) {
	type testCase struct {
		input    []byte
		expected float64
		error    assert.ErrorAssertionFunc
	}

	tt := map[string]testCase{
		"empty": {
			input:    []byte{},
			expected: 0,
			error:    assert.NoError,
		},
		"nil": {
			expected: 0,
			error:    assert.NoError,
		},
		"base number": {
			input:    []byte("123"),
			expected: 123,
			error:    assert.NoError,
		},
		"base number in quotes": {
			input:    []byte(`"123"`),
			expected: 123,
			error:    assert.NoError,
		},
		"float number": {
			input:    []byte("1.23"),
			expected: 1.23,
			error:    assert.NoError,
		},
		"float number in quoted": {
			input:    []byte(`"1.23"`),
			expected: 1.23,
			error:    assert.NoError,
		},
		"invalid number": {
			input:    []byte("wrong"),
			expected: 0,
			error:    assert.Error,
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			get, err := B2Float64(tc.input)
			tc.error(t, err)
			assert.Equal(t, tc.expected, get)
		})
	}
}
