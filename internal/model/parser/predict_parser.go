package parser

import (
	tcsv "github.com/nyunja/c4.5-decision-tree/internal/csv"
	t "github.com/nyunja/c4.5-decision-tree/internal/model/types"
)

// PredictionCSVParser efficiently parses a CSV file for prediction (target column may not exist)
func PredictionCSVParser(file string, hasHeader bool, chunkSize int, targetColumn string) ([]t.Instance, []string, map[string]string, error) {
	// Open file and create CSV reader
	f, csvReader, err := tcsv.OpenCSVFile(file)
	if err != nil {
		return nil, nil, nil, err
	}
	defer f.Close()

	// Read headers
	headers, err := tcsv.ReadCSVHeaders(csvReader, hasHeader)
	if err != nil {
		return nil, nil, nil, err
	}

	// First pass: collect statistics about the data
	stats, err := tcsv.CollectDatasetStatistics(f, headers, hasHeader)
	if err != nil {
		return nil, nil, nil, err
	}

	// Determine column types and ID columns
	featureTypes := tcsv.DetermineColumnTypes(stats)

	// Second pass: read and convert data
	instances, err := tcsv.LoadPredictionInstances(file, headers, featureTypes, targetColumn, stats.RowCount, chunkSize, hasHeader)
	if err != nil {
		return nil, nil, nil, err
	}

	return instances, headers, featureTypes, nil
}
