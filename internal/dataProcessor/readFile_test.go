package dataprocessor

import (
	"errors"
	"os"
	"reflect"
	"testing"
)

// Mock CSV data
var sampleCSV = `name,age,dob,timestamp
Joe,44,1981-03-13,2025-03-13T10:00:00Z
Nick,20,2006-01-12,2026-03-13T12:36:44Z
Charles,,1995-11-10,`

// create a temporary CSV file
func CreateTempCSV(content string) (string, error) {
	tempFile, err := os.CreateTemp("", "test_*.csv")
	if err != nil {
		return "", errors.New(err.Error())
	}
	defer tempFile.Close()

	// write the contents into the file
	_, err = tempFile.WriteString(content)
	return tempFile.Name(), err
}

// test CreateTempCSV()
func TestLoadCSV(t *testing.T) {
	tempFile, err := CreateTempCSV(sampleCSV)
	if err != nil {
		t.Fatalf("Failed to create temporary CSV: %v", err)
	}
	defer os.Remove(tempFile)

	// read the contents of the CSV file
	data, err := ReadCSVFile(tempFile)
	if err != nil {
		t.Fatalf("Error reading CSV file: %v", err)
	}

	expectedColumnHeader := []string{"name", "age", "dob", "timestamp"}
	if !reflect.DeepEqual(data.Header, expectedColumnHeader) {
		t.Errorf("Expected header %v, got %v", expectedColumnHeader, data.Header)
	}
}
