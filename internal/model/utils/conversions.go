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
