package main

import "testing"

func CalculateTwoTest(t *testing.T) {
	if CalculateTwo(2) != 4 {
		t.Error("expected 2 + 2 to equal 4")
	}
}

func CalculateTwoMultipleTest(t *testing.T) {
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
		if output := CalculateTwo(test.input); output != test.expected {
			t.Error("test failed: {} input, {} expected, recieved: {}", test.input, test.expected, output)
		}
	}
}
