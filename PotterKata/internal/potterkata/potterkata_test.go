package potterkata

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	twoBookDiscount   = 0.05
	threeBookDiscount = 0.1
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
		actual := CalculatePotterBookPrice(test.firstBook, test.secondBook, test.thirdBook, test.fourthBook, test.fifthBook)
		assert.Equal(t, actual, test.expected, fmt.Sprintf("%f should return %f", actual, test.expected))
	}
}

func TestTwoDiscount(t *testing.T) {
	var tests = []struct {
		firstBook  int
		secondBook int
		thirdBook  int
		fourthBook int
		fifthBook  int
		expected   float64
	}{
		{1, 1, 0, 0, 0, (8 * 2) - (8 * 2 * twoBookDiscount)},
	}

	for _, test := range tests {
		actual := CalculatePotterBookPrice(test.firstBook, test.secondBook, test.thirdBook, test.fourthBook, test.fifthBook)
		assert.Equal(t, actual, test.expected, fmt.Sprintf("%f should return %f", actual, test.expected))
	}
}

func TestThreeDiscount(t *testing.T) {
	var tests = []struct {
		firstBook  int
		secondBook int
		thirdBook  int
		fourthBook int
		fifthBook  int
		expected   float64
	}{
		{1, 1, 1, 0, 0, (8 * 3) - (8 * 3 * threeBookDiscount)},
	}

	for _, test := range tests {
		actual := CalculatePotterBookPrice(test.firstBook, test.secondBook, test.thirdBook, test.fourthBook, test.fifthBook)
		assert.Equal(t, actual, test.expected, fmt.Sprintf("%f should return %f", actual, test.expected))
	}
}

func TestMultipleDiscount(t *testing.T) {
	var tests = []struct {
		firstBook  int
		secondBook int
		thirdBook  int
		fourthBook int
		fifthBook  int
		expected   float64
	}{
		{2, 2, 2, 1, 1, (8 * 8) - (8 * 3 * threeBookDiscount) - (8 * 3 * threeBookDiscount)},
	}

	for _, test := range tests {
		actual := CalculatePotterBookPrice(test.firstBook, test.secondBook, test.thirdBook, test.fourthBook, test.fifthBook)
		assert.Equal(t, actual, test.expected, fmt.Sprintf("%f should return %f", actual, test.expected))
	}
}
