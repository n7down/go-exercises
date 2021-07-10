package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNoDiscount(t *testing.T) {
	var tests = []struct {
		firstBook  int
		secondBook int
		thirdBook  int
		fourthBook int
		fifthBook  int
		expected   float64
	}{
		{1, 0, 0, 0, 0, 8.0},
	}

	for _, test := range tests {
		actual := CalculatePotterDiscount(test.firstBook, test.secondBook, test.thirdBook, test.fourthBook, test.fifthBook)
		assert.Equal(t, actual, test.expected, fmt.Sprintf("%f should return %f", actual, test.expected))
	}
}

func TestTwoDiscount(t *testing.T) {
	var twoBookDiscount = 0.05
	var tests = []struct {
		firstBook  int
		secondBook int
		thirdBook  int
		fourthBook int
		fifthBook  int
		expected   float64
	}{
		{1, 1, 0, 0, 0, (8 * 2 * twoBookDiscount) - (8 * 2)},
	}

	for _, test := range tests {
		actual := CalculatePotterDiscount(test.firstBook, test.secondBook, test.thirdBook, test.fourthBook, test.fifthBook)
		assert.Equal(t, actual, test.expected, fmt.Sprintf("%f should return %f", actual, test.expected))
	}
}
