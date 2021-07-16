package leftrotation

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLeftRotationWithZeroElements(t *testing.T) {
	var tests = []struct {
		input           []int32
		numberRotations int32
		expected        []int32
	}{
		{
			[]int32{},
			1,
			[]int32{},
		},
	}
	for _, test := range tests {
		actual := RotateLeft(test.input, test.numberRotations)
		assert.Equal(t, test.expected, actual, fmt.Sprintf("%v should return %v", actual, test.expected))
	}
}

func TestLeftRotationWithOneElement(t *testing.T) {
	var tests = []struct {
		input           []int32
		numberRotations int32
		expected        []int32
	}{
		{
			[]int32{1},
			1,
			[]int32{1},
		},
	}
	for _, test := range tests {
		actual := RotateLeft(test.input, test.numberRotations)
		assert.Equal(t, test.expected, actual, fmt.Sprintf("%v should return %v", actual, test.expected))
	}
}

func TestLeftRotationWithTwoElements(t *testing.T) {
	var tests = []struct {
		input           []int32
		numberRotations int32
		expected        []int32
	}{
		{
			[]int32{1, 2},
			1,
			[]int32{2, 1},
		},
	}
	for _, test := range tests {
		actual := RotateLeft(test.input, test.numberRotations)
		assert.Equal(t, test.expected, actual, fmt.Sprintf("%v should return %v", actual, test.expected))
	}
}

func TestLeftRotationWithFiveElements(t *testing.T) {
	var tests = []struct {
		input           []int32
		numberRotations int32
		expected        []int32
	}{
		{
			[]int32{1, 2, 3, 4, 5},
			4,
			[]int32{5, 1, 2, 3, 4},
		},
	}
	for _, test := range tests {
		actual := RotateLeft(test.input, test.numberRotations)
		assert.Equal(t, test.expected, actual, fmt.Sprintf("%v should return %v", actual, test.expected))
	}
}

func TestLeftRotationWithEightElements(t *testing.T) {
	var tests = []struct {
		input           []int32
		numberRotations int32
		expected        []int32
	}{
		{
			[]int32{1, 2, 3, 4, 5, 6, 7, 8},
			4,
			[]int32{5, 6, 7, 8, 1, 2, 3, 4},
		},
	}
	for _, test := range tests {
		actual := RotateLeft(test.input, test.numberRotations)
		assert.Equal(t, test.expected, actual, fmt.Sprintf("%v should return %v", actual, test.expected))
	}
}
