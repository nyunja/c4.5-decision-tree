package dataprocessor

import (
	"errors"
	"fmt"
	"strconv"
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

	for colIdx, colName := range columnHeader {
		isNumeric, isDate, isTimestamp := true, true, true

		// Check types
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

		// Assign type
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
	}

	// Parse Data
	for rowIdx, row := range data {
		parsedRow := make([]any, numCols)
		for colIdx, value := range row {
			if value == "" {
				parsedRow[colIdx] = nil
				continue
			}

			var err error
			switch columnTypes[colIdx].Type {
			case NumType:
				parsedRow[colIdx], err = strconv.ParseFloat(value, 64)
			case DateType:
				parsedRow[colIdx], err = time.Parse("2006-01-02", value)
			case TimestampType:
				parsedRow[colIdx], err = time.Parse(time.RFC3339, value)
			default:
				parsedRow[colIdx] = value
			}

			if err != nil {
				return nil, nil, fmt.Errorf("error parsing column %d: %v", colIdx, err)
			}
		}
		parsedData[rowIdx] = parsedRow
	}

	return parsedData, columnTypes, nil
}
