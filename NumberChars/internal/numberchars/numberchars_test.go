package numberchars

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumberChars(t *testing.T) {
	var tests = []struct {
		input    string
		expected string
	}{
		{"aabbcc", "2a2b2c"},
		{"aabbccaa", "2a2b2c2a"},
	}

	for _, test := range tests {
		actual := NumberChars(test.input)
		assert.Equal(t, actual, test.expected, fmt.Sprintf("%s should return %s", test.input, test.expected))
	}
}
