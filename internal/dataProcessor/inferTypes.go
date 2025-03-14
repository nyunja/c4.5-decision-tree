package dataprocessor

import (
	"errors"
	"fmt"
	"strconv"
	"sync"
	"time"
)

// InferAndParseData infers column types and parses data simultaneously.
func InferColumnTypes(data [][]string, columnHeader []string) ([][]any, []ColumnData, error) {
	if len(data) == 0 {
		return nil, nil, errors.New("empty dataset, cannot infer datatype")
	}

	numRows := len(data)
	numCols := len(columnHeader)

	parsedData := make([][]any, numRows)
	columnTypes := make([]ColumnData, numCols)

	var wg sync.WaitGroup
	errChan := make(chan error, numCols)

	for colIdx, colName := range columnHeader {
		wg.Add(1)
		go func(colIdx int, colName string) {
			defer wg.Done()

			isNumeric, isDate, isTimestamp := true, true, true

			for _, row := range data {
				if colIdx >= len(row) {
					continue
				}

				value := row[colIdx]
				if value == "" {
					continue
				}

				if _, err := strconv.ParseFloat(value, 64); err != nil {
					isNumeric = false
				}
				if _, err := time.Parse("2006-01-02", value); err != nil {
					isDate = false
				}
				if _, err := time.Parse(time.RFC3339, value); err != nil {
					isTimestamp = false
				}
			}

			switch {
			case isNumeric:
				columnTypes[colIdx] = ColumnData{Name: colName, Type: NumType}
			case isDate:
				columnTypes[colIdx] = ColumnData{Name: colName, Type: DateType}
			case isTimestamp:
				columnTypes[colIdx] = ColumnData{Name: colName, Type: TimestampType}
			default:
				columnTypes[colIdx] = ColumnData{Name: colName, Type: CategoryType}
			}

			for rowIdx, row := range data {
				if parsedData[rowIdx] == nil {
					parsedData[rowIdx] = make([]any, numCols)
				}

				value := row[colIdx]
				if value == "" {
					parsedData[rowIdx][colIdx] = nil
					continue
				}

				var err error
				switch columnTypes[colIdx].Type {
				case NumType:
					parsedData[rowIdx][colIdx], err = strconv.ParseFloat(value, 64)
				case DateType:
					parsedData[rowIdx][colIdx], err = time.Parse("2006-01-02", value)
				case TimestampType:
					parsedData[rowIdx][colIdx], err = time.Parse(time.RFC3339, value)
				default:
					parsedData[rowIdx][colIdx] = value
				}

				if err != nil {
					errChan <- fmt.Errorf("error parsing column %d: %v", colIdx, err)
					return
				}
			}
		}(colIdx, colName)
	}

	wg.Wait()
	close(errChan)

	if err, ok := <-errChan; ok {
		return nil, nil, err
	}

	return parsedData, columnTypes, nil
}
