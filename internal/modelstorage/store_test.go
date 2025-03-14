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
	tree := &internal.JSONTreeNode{
		SplitFeature: "Hair Color",
		SplitType:    "Categorical",
		Children: map[string]*internal.JSONTreeNode{
			"Blonde": {ClassDistribution: map[string]int{"Yes": 10, "No": 2}},
			"Dark":   {ClassDistribution: map[string]int{"Yes": 3, "No": 15}},
		},
		ClassDistribution: map[string]int{"Yes": 13, "No": 17},
	}

	// Call SaveModel function
	err := SaveModelTree(tree, "./test.dt")
	assert.NoError(t, err, "SaveModel should not return an error")

	// Check if the file exists
	_, err = os.Stat("./test.dt")
	assert.NoError(t, err, "Saved model file should exist")

	// Load the saved file and verify JSON structure
	file, err := os.Open("./test.dt")
	assert.NoError(t, err, "File should open successfully")
	defer file.Close()

	decoder := json.NewDecoder(file)
	var loadedTree internal.JSONTreeNode
	err = decoder.Decode(&loadedTree)
	assert.NoError(t, err, "File content should be valid JSON")

	// Check if the loaded tree matches the original
	assert.Equal(t, tree.SplitFeature, loadedTree.SplitFeature, "Root feature should match")
	assert.Equal(t, tree.SplitType, loadedTree.SplitType, "Split type should match")
	assert.Equal(t, tree.ClassDistribution, loadedTree.ClassDistribution, "Class distribution should match")

	os.Remove("./test.dt")
}
