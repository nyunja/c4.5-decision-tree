package internal

import (
	"math"
	"testing"
)

/*
* TestGainRatio - unit test for the GainRatio function using table-driven tests
*
* input:
*  test cases with slices of attribute values and corresponding class labels to check if the function calculates correct gain ratio
*
* output:
*  a pass or fail result based on expected vs actual gain ratio calculation
 */

func TestGainRatio(t *testing.T) {
	tests := []struct {
		attributeValues []int
		classLabels     []int
		expected        float64
	}{
		{
			// Edge case where gain is 0 (no information gain)
			attributeValues: []int{1, 1, 1, 1},
			classLabels:     []int{0, 0, 0, 0},
			expected:        0.0, // No gain when all labels are the same
		},
		{
			// Edge case where split info is very small
			attributeValues: []int{1, 1, 2, 2},
			classLabels:     []int{0, 1, 0, 1},
			expected:        0.0, // Split info is small, resulting in zero gain ratio
		},
		{
			// Small number of values with no gain
			attributeValues: []int{1, 1},
			classLabels:     []int{0, 0},
			expected:        0.0, // No gain as the class labels are identical
		},
	}

	// Iterate over the test cases
	for _, test := range tests {
		result := GainRatio(test.attributeValues, test.classLabels)
		if math.Abs(result-test.expected) > 1e-6 {
			t.Errorf("expected %v, got %v", test.expected, result)
		}
	}
}
