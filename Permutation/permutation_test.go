package main

import (
	"testing"
)

TestIsPermutation(t *testing.T) {
	var tests = []struct {
		input string
		permuation string
		expected bool
	}{
		{"abc", "bca", true},
	}

	for _, test := range tests {
		actual := IsPermutation(test.input, test.permuation)
	}
}
