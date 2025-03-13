package internal

import (
	"fmt"
	"strconv"
	"time"
)

func ParseData(data [][]string, metadata []ColumnData) ([][]interface{}, error) {
	parsedData := make([][]interface{}, len(data))

	for _, row := range data {
		parsedRow := make([]interface{}, len(row))
		for j, value := range row {
			if value == "" { // Handle empty/missing values
				parsedRow[j] = nil
				continue
			}

			// parse metadata values
			switch metadata[j].Type {
			case NumType:
				num, err := strconv.ParseFloat(value, 64)
				if err != nil {
					return nil, fmt.Errorf("invalid numerical value: %v", err)
				}
				parsedRow[j] = num
			case DateType:
				date, err := time.Parse("2006-01-02", value)
				if err != nil {
					return nil, fmt.Errorf("invalid date value: %v", err)
				}
				parsedRow[j] = date
			case TimestampType:
				timeStamp, err := time.Parse(time.RFC3339, value)
				if err != nil {
					return nil, fmt.Errorf("invalid timestamp: %v", err)
				}
				parsedRow[j] = timeStamp
			default:
				parsedRow[j] = value
			}
		}
	}
	return parsedData, nil
}
