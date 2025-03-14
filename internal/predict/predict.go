package predict

import (
	"github.com/nyunja/c45-decision-tree/internal"
	dataprocessor "github.com/nyunja/c45-decision-tree/internal/dataProcessor"
	"github.com/nyunja/c45-decision-tree/util"
)

func Predict(modelFile, inputFile, outputFile string) error {
	// load data
	_, err := dataprocessor.ReadJSONFile(modelFile)
	if err != nil {
		return err
	}

	_, err = dataprocessor.ReadCSVFile(inputFile)
	if err != nil {
		return err
	}

	return nil
}

func IndexHeader(header []string) map[string]int {
	indexMap := make(map[string]int)
	for i, value := range header {
		indexMap[value] = i
	}
	return indexMap
}

func traverseTreeAndPredict(node *internal.JSONTreeNode, row []string, headers map[string]int) (Prediction, error) {
	if node.Children == nil || len(node.Children) == 0 {
		return calculateBestClassAndConfidence(node.ClassDistribution), nil
	}
	idx, found := headers[node.SplitType]
	if !found {
		util.LogError("target_column_not_found")
	}
	value := row[idx]
	if node.SplitType == "categorical" {
		child, found := node.Children[value]
		if found {
			traverseTreeAndPredict(child, row, headers)
		}
	}
	return calculateBestClassAndConfidence(node.ClassDistribution), nil
}

type Prediction struct {
	Class      string
	Confidence float64
}

func calculateBestClassAndConfidence(classDistribution map[string]int) Prediction {
	total := 0
	var bestClass string
	maxProb := 0.0

	// Calculate total instances in the class distribution
	for _, count := range classDistribution {
		total += count
	}

	if total == 0 {
		return Prediction{Class: "unknown", Confidence: 0.0}
	}

	// Get probabilities
	for class, count := range classDistribution {
		prob := float64(count) / float64(total)
		if prob > maxProb {
			maxProb = prob
			bestClass = class
		}
	}
	return Prediction{Class: bestClass, Confidence: maxProb}
}
