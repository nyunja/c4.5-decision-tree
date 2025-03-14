package modalstorage

import (
	"encoding/json"
	"fmt"
	"os"

	model "github.com/nyunja/c45-decision-tree/internal/model"
)

func SaveModel(tree *model.JSONTreeNode, fileinput string) error {
	filePath := fmt.Sprintf("./models/%s", fileinput)

	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(tree)
}
