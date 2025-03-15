package utils

import (
	"errors"
	"strconv"
	"strings"
	"time"

	t "github.com/nyunja/c4.5-decision-tree/internal/model/types"
)

func ConvertStringToTimestamp(value string) (*time.Time, error) {
	timeVal, err := time.Parse(time.RFC3339, value)
	if err != nil {
		timeVal, err = time.Parse("2006-01-02 15:04:05", value)
		if err != nil {
			return nil, errors.New("could not convert to timestamp")
		}
	}
	return &timeVal, nil
}

func ConvertStringToNumerical(value string) (float64, error) {
	floatVal, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return 0.0, errors.New("could not convert to numerical")
	}
	return floatVal, nil
}

func ConvertStringToDate(value string) (*time.Time, error) {
	value = strings.TrimSpace(value) // Remove any extra spaces

	dateFormats := []string{
		"2006-01-02",      // YYYY-MM-DD
		"01/02/2006",      // MM/DD/YYYY (US)
		"02/01/2006",      // DD/MM/YYYY (Europe)
		"2006/01/02",      // YYYY/MM/DD
		"2006.01.02",      // YYYY.MM.DD
		"02-01-2006",      // DD-MM-YYYY
		"02 January 2006", // DD Month YYYY (e.g., 02 January 2024)
	}

	for _, format := range dateFormats {
		if dateVal, err := time.Parse(format, value); err == nil {
			return &dateVal, nil
		}
	}

	return nil, errors.New("could not convert string to date")
}

// isDateValue checks if a string is a valid date
func IsDateValue(value string) bool {
	_, err1 := time.Parse("2006-01-02", value)
	if err1 == nil {
		return true
	}

	_, err2 := time.Parse("01/02/2006", value)
	if err2 == nil {
		return true
	}

	_, err3 := time.Parse("2006/02/01", value)

	return err3 == nil
}

// isTimestampValue checks if a string is a valid timestamp
func IsTimestampValue(value string) bool {
	_, err1 := time.Parse(time.RFC3339, value)
	if err1 == nil {
		return true
	}

	_, err2 := time.Parse("2006-01-02 15:04:05", value)
	return err2 == nil
}

// convertPredictionRecordToInstance converts a CSV record to an Instance object for prediction
func ConvertPredictionRecordToInstance(record []string, headers []string, featureTypes map[string]string) t.Instance {
	instance := make(t.Instance, len(headers))

	for i, value := range record {
		header := headers[i]

		// Convert value based on feature type
		switch featureTypes[header] {
		case "numerical":
			if value == "" {
				instance[header] = nil
			} else {
				floatVal, err := strconv.ParseFloat(value, 64)
				if err != nil {
					instance[header] = value
				} else {
					instance[header] = floatVal
				}
			}
		case "date":
			if value == "" {
				instance[header] = nil
			} else {
				instance[header] = ConvertToDateValue(value)
			}
		case "timestamp":
			if value == "" {
				instance[header] = nil
			} else {
				instance[header] = ConvertToTimestampValue(value)
			}
		default:
			instance[header] = value
		}
	}

	return instance
}

// convertToDateValue converts a string to a date value
func ConvertToDateValue(value string) interface{} {
	dateVal, err := time.Parse("2006-01-02", value)
	if err == nil {
		return dateVal
	}

	dateVal, err = time.Parse("01/02/2006", value)
	if err == nil {
		return dateVal
	}

	dateVal, err = time.Parse("02/01/2006", value)
	if err == nil {
		return dateVal
	}

	// If all conversions fail, return the original string
	return value
}

// convertToTimestampValue converts a string to a timestamp value
func ConvertToTimestampValue(value string) interface{} {
	timeVal, err := time.Parse(time.RFC3339, value)
	if err == nil {
		return timeVal
	}

	timeVal, err = time.Parse("2006-01-02 15:04:05", value)
	if err == nil {
		return timeVal
	}

	// If all conversions fail, return the original string
	return value
}
