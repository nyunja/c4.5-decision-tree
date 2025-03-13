package modalstorage

import (
	"encoding/json"
	"os"
	"testing"

	internal "github.com/nyunja/c45-decision-tree/internal" // Change this to the correct import path

	"github.com/stretchr/testify/assert"
)

func TestSaveModel(t *testing.T) {
	// Create a sample decision tree
	tree := &internal.DecisionTree{
		Feature:   "Height",
		Threshold: 170,
		SplitType: "continuous",
		Left: &internal.DecisionTree{
			Feature:   "Hair",
			SplitType: "categorical",
			Categories: map[string]*internal.DecisionTree{
				"Blonde": {Label: "Yes"}, // Leaf node
				"Dark":   {Label: "No"},  // Leaf node
			},
		},
		Right: &internal.DecisionTree{
			Feature:   "Eyes",
			SplitType: "categorical",
			Categories: map[string]*internal.DecisionTree{
				"Blue":  {Label: "Yes"}, // Leaf node
				"Brown": {Label: "No"},  // Leaf node
			},
		},
		Metadata: &internal.Metadata{
			Features: []string{"Height", "Hair", "Eyes"},
			Types:    map[string]string{"Height": "continuous", "Hair": "categorical", "Eyes": "categorical"},
		},
	}

	// Call SaveModel function
	err := SaveModel(tree)
	assert.NoError(t, err, "SaveModel should not return an error")

	// Check if the file exists
	_, err = os.Stat("./store.dt")
	assert.NoError(t, err, "Saved model file should exist")

	// Load the saved file and verify JSON structure
	file, err := os.Open("./store.dt")
	assert.NoError(t, err, "File should open successfully")
	defer file.Close()

	decoder := json.NewDecoder(file)
	var loadedTree internal.DecisionTree
	err = decoder.Decode(&loadedTree)
	assert.NoError(t, err, "File content should be valid JSON")

	// Check if the loaded tree matches the original
	assert.Equal(t, tree.Feature, loadedTree.Feature, "Root feature should match")
	assert.Equal(t, tree.Threshold, loadedTree.Threshold, "Threshold should match")
	assert.Equal(t, tree.Metadata.Features, loadedTree.Metadata.Features, "Metadata features should match")

	// Verify leaf nodes
	assert.Equal(t, "Yes", loadedTree.Left.Categories["Blonde"].Label, "Blonde hair should lead to Yes")
	assert.Equal(t, "No", loadedTree.Left.Categories["Dark"].Label, "Dark hair should lead to No")
	assert.Equal(t, "Yes", loadedTree.Right.Categories["Blue"].Label, "Blue eyes should lead to Yes")
	assert.Equal(t, "No", loadedTree.Right.Categories["Brown"].Label, "Brown eyes should lead to No")

	os.Remove("./store.dt")
}
