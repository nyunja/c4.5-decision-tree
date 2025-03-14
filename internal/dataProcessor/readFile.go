package dataprocessor

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"runtime"
	"sync"
)

type ColumnType string

const (
	NumType       ColumnType = "Numeric"
	CategoryType  ColumnType = "Categorical"
	DateType      ColumnType = "Date"
	TimestampType ColumnType = "Timestamp"
)

// store column names & their types
type ColumnData struct {
	Name string
	Type ColumnType
}

// Dataset represent parsed CSV file
type Dataset struct {
	Header   []string
	Data     [][]any
	Metadata []ColumnData
}

func ReadCSVFile(filename string) (*Dataset, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %v ", err)
	}
	defer file.Close()

	// read the CSV file contents
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("error reading CSV: %v", err)
	}

	if len(records) < 2 {
		return nil, errors.New("CSV must have at least one row of data")
	}

	columnHeader := records[0]
	data := records[1:]

	metadata, err := InferColumnTypes(data, columnHeader)
	if err != nil {
		return nil, err
	}

	parsedData := make([][]any, len(data))
	var wg sync.WaitGroup
	numWorkers := runtime.NumCPU()
	WorkerPool := make(chan struct{}, numWorkers)

	for i, row := range data {
		wg.Add(1)
		WorkerPool <- struct{}{}

		go func(i int, row []string) {
			defer func() { <-WorkerPool }()
			defer wg.Done()

			// convert data to appropriate types
			parsedRowData, err := ParseData(row, metadata)
			if err != nil {
				fmt.Printf("Error parsing row %d: %v\n", i, err)
				return
			}
			parsedData[i] = parsedRowData
		}(i, row)
	}
	wg.Wait()

	return &Dataset{
		Header:   columnHeader,
		Data:     parsedData,
		Metadata: metadata,
	}, nil
}
