package model

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"

	t "github.com/nyunja/c4.5-decision-tree/internal/model/types" // You'll need to replace this with your actual package name
)

func TestSaveAndLoadModel(t *testing.T) {
	// Setup - create a temporary directory for test files
	testDir := "./decision_model"
	defer os.RemoveAll(testDir) // Clean up after the test

	// Create a test model
	testModel := createTestModel()

	// Test SaveModel
	filename := "test_model.json"
	err := SaveModel(testModel, filename)
	if err != nil {
		t.Fatalf("SaveModel failed: %v", err)
	}

	// Verify file exists
	filePath := filepath.Join(testDir, filename)
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		t.Fatalf("SaveModel did not create file at expected path: %s", filename)
	}

	// Test LoadModel
	loadedModel, err := LoadModel(filename)
	if err != nil {
		t.Fatalf("LoadModel failed: %v", err)
	}

	// Verify model was loaded correctly
	if !reflect.DeepEqual(testModel, loadedModel) {
		t.Fatalf("Loaded model does not match saved model")
	}
}

// Helper function to create a test model
func createTestModel() *t.Model {
	child1 := &t.Node{
		Feature: "",
		IsLeaf:  true,
		Class:   "category_a",
	}

	child2 := &t.Node{
		Feature: "",
		IsLeaf:  true,
		Class:   "category_b",
	}

	root := &t.Node{
		Feature:    "age",
		IsLeaf:     false,
		Continuous: true,
		Threshold:  30.0,
		Children:   []*t.Node{child1, child2},
	}

	return &t.Model{
		Root:         root,
		FeatureTypes: map[string]string{"age": "numerical", "gender": "categorical"},
		FeatureNames: []string{"age", "gender"},
		TargetName:   "category",
	}
}
