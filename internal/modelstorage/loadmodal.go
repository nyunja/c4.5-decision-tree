package modalstorage

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/nyunja/c45-decision-tree/internal"
)

func RetrieveModelTree(filename string) (*internal.JSONTreeNode, error) {
	filePath := fmt.Sprintf("./models/%s", filename)
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var tree internal.JSONTreeNode
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&tree)
	if err != nil {
		return nil, err
	}
	return &tree, nil
}
