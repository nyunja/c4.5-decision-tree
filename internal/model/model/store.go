package model

import (
	"encoding/json"
	"fmt"
	"os"

	t "github.com/nyunja/c4.5-decision-tree/internal/model/types"
)

// SaveModel saves a model to a file
func SaveModel(model *t.Model, filename string) error {
	modelJSON, err := json.MarshalIndent(model, "", "  ")
	if err != nil {
		return fmt.Errorf("error marshaling model to JSON: %v", err)
	}
	modelDir := "./decision_model"
	if _, err := os.Stat(modelDir); err != nil {
		os.MkdirAll(modelDir, 0o755)
	}

	filePath := fmt.Sprintf("%s/%s", modelDir, filename)
	err = os.WriteFile(filePath, modelJSON, 0o644)
	if err != nil {
		return fmt.Errorf("error writing model to file: %v", err)
	}

	return nil
}

// LoadModel loads a model from a file
func LoadModel(filename string) (*t.Model, error) {
	modelDir := "./decision_model"
	filePath := fmt.Sprintf("%s/%s", modelDir, filename)
	modelJSON, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("error reading model from file: %v", err)
	}

	var model t.Model
	err = json.Unmarshal(modelJSON, &model)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling model from JSON: %v", err)
	}

	return &model, nil
}
