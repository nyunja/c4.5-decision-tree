package model

import (
	"math"
	"testing"
)

// TestEntropy tests the Entropy function.
func TestEntropy(t *testing.T) {
	tests := []struct {
		Classes  []int
		Expected float64
	}{
		{[]int{}, 0},           // empty array
		{[]int{1, 1, 1, 1}, 0}, // if all elements are the same, then there is 0 entropy
		{[]int{1, 2}, 1},       // 2 classes with equal probability
		{[]int{1, 2, 3, 4}, 2}, // 4 with equal probability
	}

	for _, test := range tests {
		got := Entropy(test.Classes)
		if math.Abs(got-test.Expected) > 1e-9 {
			t.Errorf("expected Entropy of %v to be %f; but got %f", test.Classes, test.Expected, got)
		}
	}
}
