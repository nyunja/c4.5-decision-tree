package parser

import (
	tcsv "github.com/nyunja/c4.5-decision-tree/internal/csv"
	t "github.com/nyunja/c4.5-decision-tree/internal/model/types"
)

// StreamingCSVParser efficiently parses a CSV file in chunks
func StreamingCSVParser(file string, hasHeader bool, chunkSize int, targetColumn string) ([]t.Instance, []string, map[string]string, error) {
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
	// idColumns := utils.DetectIDColumns(stats, headers)
	// fmt.Printf("Detected ID columns: %v\n", idColumns)

	// Second pass: read and convert data
	instances, err := tcsv.LoadInstances(file, headers, featureTypes, targetColumn, stats.RowCount, chunkSize, hasHeader)
	if err != nil {
		return nil, nil, nil, err
	}

	return instances, headers, featureTypes, nil
}
