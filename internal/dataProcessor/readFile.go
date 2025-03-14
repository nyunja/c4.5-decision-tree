package dataprocessor

import (
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"sync"

	"github.com/nyunja/c45-decision-tree/internal"
)

type ColumnType string

// chunk data to reduce processing time
var RowPool = sync.Pool{
	New: func() any {
		row := make([]any, 100)
		return &row
	},
}

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
		return nil, fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	// Read all rows at once
	reader := csv.NewReader(file)
	// Channel to receive rows
	rowChan := make(chan []string)
	errChan := make(chan error, 1)
	doneChan := make(chan struct{})

	// Goroutine to read rows
	go func() {
		defer close(rowChan)
		for {
			row, err := reader.Read()
			if err == io.EOF {
				break
			}
			if err != nil {
				errChan <- fmt.Errorf("error reading CSV file: %v", err)
				return
			}
			rowChan <- row
		}
		doneChan <- struct{}{}
	}()

	// Collect all rows
	var allRows [][]string
	for {
		select {
		case row := <-rowChan:
			allRows = append(allRows, row)
		case err := <-errChan:
			return nil, err
		case <-doneChan:
			close(errChan)
			close(doneChan)
			goto PARSE
		}
	}

PARSE:
	if len(allRows) < 1 {
		return nil, errors.New("CSV is empty")
	}

	columnHeader := allRows[0]
	data := allRows[1:] // Exclude the header row

	// Infer and parse in a single step
	parsedData, metadata, err := InferColumnTypes(data, columnHeader)
	if err != nil {
		return nil, err
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
		return internal.JSONTreeNode{}, fmt.Errorf("error opening JSON file: %v", err)
	}
	defer file.Close()

	// Read dataFile
	data, err := io.ReadAll(file)
	if err != nil {
		return internal.JSONTreeNode{}, fmt.Errorf("error reading JSON file: %v", err)
	}

	// Unmarshal JSON data into DecisionTree struct
	var trainedModel internal.JSONTreeNode
	err = json.Unmarshal(data, &trainedModel)
	if err != nil {
		return internal.JSONTreeNode{}, fmt.Errorf("rror unmarshalling JSON: %v", err)
	}

	// Return the trained DecisionTree model
	return trainedModel, nil
}
