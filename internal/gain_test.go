package internal

import (
	"testing"
)

func TestGain(t *testing.T) {
	tests := []struct {
		attributeValues []int
		classLabels     []int
		expectedGain    float64
	}{
		{
			attributeValues: []int{1, 2, 1, 2, 1, 2},
			classLabels:     []int{0, 1, 0, 1, 0, 1},
			expectedGain:    1.0,
		},
		{
			attributeValues: []int{1, 1, 1, 2, 2, 2},
			classLabels:     []int{0, 0, 0, 1, 1, 1},
			expectedGain:    1.0,
		},
		{
			attributeValues: []int{1, 1, 2, 2},
			classLabels:     []int{0, 1, 0, 1},
			expectedGain:    0.0,
		},
	}

	for _, test := range tests {
		gain := Gain(test.attributeValues, test.classLabels)
		if gain != test.expectedGain {
			t.Errorf("Gain(%v, %v) = %f; want %f", test.attributeValues, test.classLabels, gain, test.expectedGain)
		}
	}
}
