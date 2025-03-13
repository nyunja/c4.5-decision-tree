package cmd

import (
	"encoding/csv"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"dt/internal/c45"
)

// TrainModel represents the train command
func TrainModel(input, target, output string) {
	data, err := loadCSV(inputFile)
	if err != nil {
		log.Fatalf("Failed to read input file: %v", err)
	}
	if len(data) == 0 {
		log.Fatalf("Error: Empty dataset")
		return
	}
	// Find the target column index
	labelIndex, err := getLabelIndex(data[0], target)
	if err != nil {
		log.Fatalf("%v", err)
	}
	tree := c45.Train(data, labelIndex)
	err = c45.SaveModel(tree, output)
	if err != nil {
		log.Fatalf("Failed to save model: %v", err)
	}

	fmt.Println("✅ Model training complete! Saved to:", output)
}

func loadCSV(filename string) ([][]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	return reader.ReadAll()
}

func getLabelIndex(header []string, target string) (int, error) {
	targetIndex := -1
	for i, col := range header {
		fmt.Println(col)
		if strings.TrimSpace(col) == target {
			targetIndex = i
			break
		}
	}
	if targetIndex == -1 {
		return 0, errors.New("target column not found in dataset")
	}
	return targetIndex, nil
}
