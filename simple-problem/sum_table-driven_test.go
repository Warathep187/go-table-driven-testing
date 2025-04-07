package main

import (
	"testing"
)

type SumTestCase struct {
	name     string
	input    []int
	expected int
}

func TestSum(t *testing.T) {
	tests := []SumTestCase{
		{
			name:     "empty slice",
			input:    []int{},
			expected: 0,
		},
		{
			name:     "single number",
			input:    []int{1},
			expected: 1,
		},
		{
			name:     "multiple numbers",
			input:    []int{1, 2, 3, 4, 5},
			expected: 15,
		},
		{
			name:     "negative numbers",
			input:    []int{-1, -2, -3, -4, -5},
			expected: -15,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := Sum(test.input...)
			if got != test.expected {
				t.Errorf("Sum(%v) = %d; want %d", test.input, got, test.expected)
			}
		})
	}
}
