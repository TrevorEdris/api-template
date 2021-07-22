package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	description    string
	input          int
	expectedOutput int
	expectedError  error
}

func TestFibonacci(t *testing.T) {
	/*
		Fib sequence is n = (n-1) + (n-2) with base case fib(0) = 0 and fib(1) = 1.
		Negative numbers are considered invalid input.

		Example:
		0 1 2 3 4 5 6  7  8  9 10
		0 1 1 2 3 5 8 13 21 34 55
	*/
	cases := []testCase{
		{
			description:    "negative starting number",
			input:          -1,
			expectedOutput: -1,
			expectedError:  InvalidStartingNumberError,
		},
		{
			description:    "base case 0",
			input:          0,
			expectedOutput: 0,
			expectedError:  nil,
		},
		{
			description:    "base case 1",
			input:          1,
			expectedOutput: 1,
			expectedError:  nil,
		},
		{
			description:    "non base case",
			input:          10,
			expectedOutput: 55,
			expectedError:  nil,
		},
	}

	for _, tc := range cases {
		output, err := fibonacci(tc.input)
		assert.Equal(t, tc.expectedOutput, output)
		assert.Equal(t, tc.expectedError, err)
	}
}
