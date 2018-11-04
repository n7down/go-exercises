package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	var tests = []struct {
		input    string
		expected bool
	}{
		{"", true},
		{"a", true},
		{"aa", true},
		{"ab", false},
		{"aba", true},
		{"abc", false},
	}

	for _, test := range tests {
		actual := IsPalindrome(test.input)
		assert.Equal(t, actual, test.expected)
	}
}
