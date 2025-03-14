package predict

import (
	"github.com/nyunja/c45-decision-tree/internal"
	dataprocessor "github.com/nyunja/c45-decision-tree/internal/dataProcessor"
	"github.com/nyunja/c45-decision-tree/util"
)

func Predict(modelFile, inputFile, outputFile string) error {
	// load data
	model, err := dataprocessor.ReadJSONFile(modelFile)
	if err != nil {
		return err
	}

	input, err := dataprocessor.ReadCSVFile(inputFile)
	if err != nil {
		return err
	}

	headers := IndexHeader(input.Header)
	predictions := make([]internal.Prediction, 0, len(input.Data))
	for _, row := range input.Data {
		pred, err := traverseTreeAndPredict(&model, row, headers)
		if err != nil {
			return err
		}
		predictions = append(predictions, pred)
	}

	// write predictions to output file
	err = dataprocessor.WritePredictionCSVFile(outputFile, input, predictions)
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

func traverseTreeAndPredict(node *internal.JSONTreeNode, row []any, headers map[string]int) (internal.Prediction, error) {
	if node.Children == nil || len(node.Children) == 0 {
		return calculateBestClassAndConfidence(node.ClassDistribution), nil
	}
	idx, found := headers[node.SplitType]
	if !found {
		util.LogError("target_column_not_found")
	}
	value := row[idx]
	if node.SplitType == "categorical" {
		child, found := node.Children[value.(string)]
		if found {
			traverseTreeAndPredict(child, row, headers)
		}
	}
	return calculateBestClassAndConfidence(node.ClassDistribution), nil
}

func calculateBestClassAndConfidence(classDistribution map[string]int) internal.Prediction {
	total := 0
	var bestClass string
	maxProb := 0.0

	// Calculate total instances in the class distribution
	for _, count := range classDistribution {
		total += count
	}

	if total == 0 {
		return internal.Prediction{Class: "unknown", Confidence: 0.0}
	}

	// Get probabilities
	for class, count := range classDistribution {
		prob := float64(count) / float64(total)
		if prob > maxProb {
			maxProb = prob
			bestClass = class
		}
	}
	return internal.Prediction{Class: bestClass, Confidence: maxProb}
}
