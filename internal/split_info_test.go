package internal

import (
	"math"
	"testing"
)

/*
* TestSplitInfo - unit test for the SplitInfo function
*
* input:
*  test cases with slices of integers to check if the function calculates correct split information
*
* output:
*  a pass or fail result based on expected vs actual entropy calculation
 */

func TestSplitInfo(t *testing.T) {
	table := []struct {
		input    []int
		expected float64
	}{
		{
			// Test case 1: Uniform distribution (split entropy should be high)
			input:    []int{1, 1, 2, 2, 3, 3},
			expected: 1.584962500721156,
		},
		{
			// Test case 2: Single value (split entropy should be 0)
			input:    []int{1, 1, 1, 1},
			expected: 0.0,
		},
		{
			// Test case 4: Large numbers of repeated values (entropy should be close to 0)

			input:    []int{100, 100, 100, 100, 100},
			expected: 0.0,
		},
		{
			// Test case 5: Small number of values (entropy should be high)

			input:    []int{1, 2},
			expected: 1.0,
		},
	}

	for _, test := range table {
		result := SplitInfo(test.input)
		if math.Abs(result-test.expected) > 1e-6 {
			t.Errorf("Test failed: expected %v, got %v", test.expected, result)
		}
	}
}
