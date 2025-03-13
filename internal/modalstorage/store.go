package modalstorage

import (
	"encoding/json"
	"os"

	internal "github.com/nyunja/c45-decision-tree/internal"
)

func SaveModel(tree *internal.DecisionTree) error {
	file, err := os.Create("./models/modelfile.dt")
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(tree)
}
