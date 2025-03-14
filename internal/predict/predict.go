package predict

import dataprocessor "github.com/nyunja/c45-decision-tree/internal/dataProcessor"

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

	return nil
}

func IndexHeader(header []string) map[string]int {
	indexMap := make(map[string]int)
	for i, value := range header {
		indexMap[value] = i
	}
	return indexMap
}
