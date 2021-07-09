package smallestposint

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSmallestPosInt(t *testing.T) {
	var tests = []struct {
		input    []int
		expected int
	}{
		{[]int{-3, -1}, 1},
		{[]int{1, 4, 3, 2, 6, 7}, 5},
	}
	for _, test := range tests {
		actual := SmallestPosInt(test.input)
		assert.Equal(t, actual, test.expected, fmt.Sprintf("%v should return %v", test.input, test.expected))
	}
}
