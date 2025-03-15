package predict

import (
	"testing"

	typ "github.com/nyunja/c4.5-decision-tree/internal/model/types"
)

// TestPredictClassWithLeafRoot tests the case when the root node is already a leaf
func TestPredictClassWithLeafRoot(t *testing.T) {
	model := &typ.Model{
		Root: &typ.Node{
			IsLeaf: true,
			Class:  "default",
		},
	}

	instance := typ.Instance{
		"feature1": "value1",
	}

	result := PredictClass(model, instance)

	if result != "default" {
		t.Errorf("PredictClass() with leaf root = %v, want 'default'", result)
	}
}
