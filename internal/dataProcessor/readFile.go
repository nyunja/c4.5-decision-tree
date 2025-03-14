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

// chunk data to reduce processing time
var RowPool = sync.Pool{
	New: func() any {
		return make([]any, 100)
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

func Init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func ReadCSVFile(filename string) (*Dataset, error) {
	// Maximize CPU usage
	Init()

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
	for range NumWorkers {
		wg.Add(1)

		go func() {
			defer wg.Done()
			for job := range WorkerPool {
				// convert data to appropriate types
				parsedRow := RowPool.Get().([]any)
				parsedRow = parsedRow[:len(job.row)]

				err := ParseData(job.row, metadata, parsedRow)
				if err != nil {
					fmt.Printf("Error parsing row %d: %v\n", job.index, err)
					return
				}
				parsedData[job.index] = parsedRow
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
