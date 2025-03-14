package modalstorage

import (
	"encoding/json"
	"os"

	"github.com/nyunja/c45-decision-tree/internal"
)

func LoadModel() (*internal.DecisionTree, error) {
	file, err := os.Open("./models/modelfile.dt")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var tree internal.DecisionTree
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&tree)
	if err != nil {
		return nil, err
	}
	return &tree, nil
}
