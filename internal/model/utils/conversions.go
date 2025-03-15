package utils

import (
	"errors"
	"strconv"
	"time"
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
	dateVal, err := time.Parse("2006-01-02", value)
	if err != nil {
		dateVal, err = time.Parse("01/02/2006", value)
		if err != nil {
			dateVal, err = time.Parse("02/01/2006", value)
			if err != nil {
				return nil, errors.New("could not convert string to date")
			}
		}
	}
	return &dateVal, nil
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