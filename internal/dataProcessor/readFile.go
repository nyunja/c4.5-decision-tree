package dataprocessor

import (
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/nyunja/c45-decision-tree/internal"
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

	// convert data to appropriate types
	parsedData, err := ParseData(data, metadata)
	if err != nil {
		return nil, fmt.Errorf("error parsing data: %v", err)
	}
	return &Dataset{
		Header:   columnHeader,
		Data:     parsedData,
		Metadata: metadata,
	}, nil
}

func ReadJSONFile(dataFile string) (internal.JSONTreeNode, error) {
	// Open dataFile for reading
	file, err := os.Open(dataFile)
	if err != nil {
		return internal.JSONTreeNode{}, fmt.Errorf("Error opening JSON file: %v", err)
	}
	defer file.Close()

	// Read dataFile
	data, err := io.ReadAll(file)
	if err != nil {
		return internal.JSONTreeNode{}, fmt.Errorf("Error reading JSON file: %v", err)
	}

	// Unmarshal JSON data into DecisionTree struct
	var trainedModel internal.JSONTreeNode
	err = json.Unmarshal(data, &trainedModel)
	if err != nil {
		return internal.JSONTreeNode{}, fmt.Errorf("Error unmarshalling JSON: %v", err)
	}

	// Return the trained DecisionTree model
	return trainedModel, nil
}
