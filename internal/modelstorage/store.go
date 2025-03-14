package modalstorage

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/nyunja/c45-decision-tree/internal"
)

func SaveModelTree(tree *internal.JSONTreeNode, fileinput string) error {
	directorypath := "./tree_models"
	if err := os.MkdirAll(directorypath, os.ModePerm); err != nil {
		return err
	}

	filePath := fmt.Sprintf("%s/%s", directorypath, fileinput)

	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(tree)
}
