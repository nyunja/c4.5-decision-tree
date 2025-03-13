package internal

import (
	"errors"
	"strconv"
	"time"
)

// InferColumnTypes detects and sorts data according to the data type
func InferColumnTypes(data [][]string, columnHeader []string) ([]ColumnData, error) {
	if len(data) == 0 {
		return nil, errors.New("empty dataset, cannot infer datatype")
	}

	ColumnType := make([]ColumnData, len(columnHeader))

	for columnIndex, columnName := range columnHeader {
		isNumeric, isDate, isTimeStamp := true, true, true

		for _, row := range data {
			value := row[columnIndex]
			if value == "" { // skip empty values
				continue
			}

			// numeric metadata
			if _, err := strconv.ParseFloat(value, 64); err != nil {
				isNumeric = false
			}

			// date metadata
			if _, err := time.Parse("2006-01-02", value); err != nil {
				isDate = false
			}

			// timestamp metadata
			if _, err := time.Parse(time.RFC3339, value); err != nil {
				isTimeStamp = false
			}
		}

		// parse the metadata & sort accordingly
		switch {
		case isNumeric:
			ColumnType[columnIndex] = ColumnData{Name: columnName, Type: NumType}
		case isDate:
			ColumnType[columnIndex] = ColumnData{Name: columnName, Type: DateType}
		case isTimeStamp:
			ColumnType[columnIndex] = ColumnData{Name: columnName, Type: TimestampType}
		default:
			ColumnType[columnIndex] = ColumnData{Name: columnName, Type: CategoryType}
		}
	}
	return ColumnType, nil
}
