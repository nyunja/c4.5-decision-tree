package modalstorage

import (
	"encoding/json"
	"os"

	"github.com/nyunja/c45-decision-tree/internal"
)

func SaveModelTree(tree *internal.JSONTreeNode, filePath string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(tree)
}
