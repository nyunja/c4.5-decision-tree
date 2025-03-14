package predict

import dataprocessor "github.com/nyunja/c45-decision-tree/internal/dataProcessor"

func Predic(modelFile, inputFile, outputFile string) error {
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
