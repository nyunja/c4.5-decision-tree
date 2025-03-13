package internal

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"
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
	Data     [][]interface{}
	Metadata []ColumnData
}

func ReadCSVFile(filename string) (*Dataset, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("Error opening file: %v ", err)
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

	// convert data to appropriate types
	parsedData, err := ParseData(data)
	return &Dataset{
		Header:   columnHeader,
		Data:     parsedData,
		Metadata: metadata,
	}, nil
}
