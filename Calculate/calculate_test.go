package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalculateTwo(t *testing.T) {
	output := CalculateTwo(2)
	assert.Equal(t, output, 4, fmt.Sprintf("test failed: expected: %d, actual: %d", output, 4))
}

func TestCalculateTwoTable(t *testing.T) {
	var tests = []struct {
		input    int
		expected int
	}{
		{2, 4},
		{-1, 1},
		{0, 2},
		{99999, 100001},
	}

	for _, test := range tests {
		output := CalculateTwo(test.input)
		assert.Equal(t, output, test.expected, fmt.Sprintf("test failed: input: %d, expected: %d, recieved: %d", test.input, test.expected, output))
	}
}
