package util

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

func ReadCSVFile(filename string) ([]string, []float64, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	// read the CSV file contents
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, nil, err
	}

	var labels []string
	var probalities []float64

	for _, record := range records {
		if len(record) < 2 {
			continue
		}

		label := record[0]
		p, err := strconv.ParseFloat(record[1], 64)
		if err != nil {
			fmt.Printf("Error converting to float64: %v", err)
		}

		// append labels and probabilities
		labels = append(labels, label)
		probalities = append(probalities, p)
	}

	return labels, probalities, nil
}

