package entropy

import (
	"testing"

	test "github.com/nyunja/c4.5-decision-tree/internal/model/types"
)

func TestCalculateEntropy(t *testing.T) {
	tests := []struct {
		name          string
		instance      []test.Instance
		targetFeature string
		expected      float64
	}{
		{
			name:          "Empty instances",
			instance:      []test.Instance{},
			targetFeature: "class",
			expected:      0,
		},
		{
			name: "Single class",
			instance: []test.Instance{
				{"class": "A"},
				{"class": "A"},
				{"class": "A"},
			},
			targetFeature: "class",
			expected:      0,
		},
		{
			name: "Two classes",
			instance: []test.Instance{
				{"class": "A"},
				{"class": "B"},
				{"class": "A"},
				{"class": "B"},
			},
			targetFeature: "class",
			expected:      1,
		},
		{
			name: "Multiple classes",
			instance: []test.Instance{
				{"class": "A"},
				{"class": "B"},
				{"class": "C"},
				{"class": "A"},
				{"class": "B"},
				{"class": "C"},
			},
			targetFeature: "class",
			expected:      1.584962500721156,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CalculateEntropy(tt.instance, tt.targetFeature)
			if result != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}
