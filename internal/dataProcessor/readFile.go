package dataprocessor

import (
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"runtime"
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
		return nil, fmt.Errorf("error opening file: %v ", err)
	}
	defer file.Close()

	// read the CSV file contents
	reader := csv.NewReader(file)

	columnHeader, err := reader.Read()
	if err != nil {
		return nil, fmt.Errorf("error reading CSV header: %v", err)
	}

	var data [][]string
	for {
		row, err := reader.Read()
		if err != nil {
			break
		}
		data = append(data, row)
	}

	if len(data) == 0 {
		return nil, errors.New("CSV has no data")
	}

	metadata, err := InferColumnTypes(data, columnHeader)
	if err != nil {
		return nil, err
	}

	parsedData := make([][]any, len(data))
	var wg sync.WaitGroup
	NumWorkers := runtime.GOMAXPROCS(runtime.NumCPU())
	WorkerPool := make(chan struct {
		index int
		row   []string
	}, NumWorkers*2)

	// Worker pool
	for i := 0; i < NumWorkers; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			for job := range WorkerPool {
				parsedRowPtr := RowPool.Get().(*[]any)
				parsedRow := (*parsedRowPtr)[:len(job.row)]
				parsedRow = parsedRow[:len(job.row)]

				err := ParseData(job.row, metadata, parsedRow)
				if err != nil {
					fmt.Printf("Error parsing row %d: %v\n", job.index, err)
					return
				}
				RowPool.Put(parsedRowPtr)
				RowPool.Put(parsedRow)
			}
		}()
	}

	for i, row := range data {
		WorkerPool <- struct {
			index int
			row   []string
		}{index: i, row: row}
	}
	close(WorkerPool)
	wg.Wait()

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
