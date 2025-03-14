package dataprocessor

import (
	"fmt"
	"strconv"
	"time"
)

func ParseData(row []string, metadata []ColumnData) ([]any, error) {
	parsedRow := make([]any, len(row))
	for j, value := range row {
		if value == "" {
			parsedRow[j] = nil
			continue
		}

		var err error
		switch metadata[j].Type {
		case NumType:
			parsedRow[j], err = strconv.ParseFloat(value, 64)
		case DateType:
			parsedRow[j], err = time.Parse("2006-01-02", value)
		case TimestampType:
			parsedRow[j], err = time.Parse(time.RFC3339, value)
		default:
			parsedRow[j] = value
		}

		if err != nil {
			return nil, fmt.Errorf("invalid value at column %d: %v", j, err)
		}
	}
	return parsedRow, nil
}
