package c45

import (
	"encoding/json"
	"os"
)

// Decision tree node
type Node struct {
	Attribute   string
	Threshold   string
	Children    map[string]*Node
	Class       string
	IsLeaf      bool
}

func (n *Node) Predict(row []string) string {
	return "prediction"
}

func LoadModel(path string) (*Node, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var tree Node
	err = json.Unmarshal(data, &tree)
	return &tree, err
}

func SaveModel(tree *Node, path string) error {
	data, err := json.MarshalIndent(tree, "", "  ")
    if err != nil {
        return err
    }
    return os.WriteFile(path, data, 0644)
}
