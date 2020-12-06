package triangle

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_TriangleType_With_Negative_Values_Returns_Error(t *testing.T) {
	testCases := []struct {
		name          string
		triangle      *Triangle
		expectedError string
	}{
		{
			name:          "add sides have negative values",
			triangle:      NewTriangle(-1, -1, -1),
			expectedError: "not a triangle",
		}, {
			name:          "side1 has a negative value",
			triangle:      NewTriangle(-1, 1, 1),
			expectedError: "not a triangle",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			_, err := testCase.triangle.TriangleType()
			if err == nil {
				assert.FailNow(t, "err is nil")
			}
			assert.Equal(t, testCase.expectedError, err.Error())
		})
	}
}

func Test_TriangleType_With_Zero_Values_Return_Error(t *testing.T) {
	testCases := []struct {
		name          string
		triangle      *Triangle
		expectedError string
	}{
		{
			name:          "all sides are 0",
			triangle:      NewTriangle(0, 0, 0),
			expectedError: "not a triangle",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			_, err := testCase.triangle.TriangleType()
			if err == nil {
				assert.FailNow(t, "err is nil")
			}
			assert.Equal(t, testCase.expectedError, err.Error())
		})
	}
}

func Test_TriangleType_With_Two_Sides_Greater_Then_The_Other_Side_Returns_Error(t *testing.T) {
	testCases := []struct {
		name          string
		triangle      *Triangle
		expectedError string
	}{
		{
			name:          "side1 and side2 are greater then side3",
			triangle:      NewTriangle(8, 6, 3),
			expectedError: "not a triangle",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			_, err := testCase.triangle.TriangleType()
			if err == nil {
				assert.FailNow(t, "err is nil")
			}
			assert.Equal(t, testCase.expectedError, err.Error())
		})
	}
}
