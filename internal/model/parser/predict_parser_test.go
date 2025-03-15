package parser

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

// create a temporary CSV file to write out prdiction content to
func createTestCSV(content string) (string, error) {
	file, err := os.CreateTemp("", "test*.csv")
	if err != nil {
		return "", err
	}
	defer file.Close()

	_, err = file.WriteString(content)
	if err != nil {
		return "", err
	}

	return file.Name(), nil
}

/*
TestPredictionCSVParser tests the PredictionCSVParser function with various CSV inputs.
The test cases include:
- Basic CSV with header
- CSV without header
- CSV with missing values
- CSV with a large dataset
*/
func TestPredictionCSVParser(t *testing.T) {
	tests := []struct {
		name         string
		csvContent   string
		hasHeader    bool
		chunkSize    int
		targetColumn string
		expectedErr  bool
	}{
		{
			name: "Basic CSV with header",
			csvContent: `col1,col2,col3
1,2021-01-01,foo
2,2021-01-02,bar
3,2021-01-03,baz`,
			hasHeader:    true,
			chunkSize:    10,
			targetColumn: "col3",
			expectedErr:  false,
		},
		{
			name: "CSV with missing values (handles empty values)",
			csvContent: `col1,col2,col3
1,,foo
2,2021-01-02,
3,2021-01-03,baz`,
			hasHeader:    true,
			chunkSize:    10,
			targetColumn: "col3",
			expectedErr:  false,
		},
		{
			name: "CSV with large dataset (ensures chunk processing works)",
			csvContent: `col1,col2,col3
1,2021-01-01,foo
2,2021-01-02,bar
3,2021-01-03,baz
4,2021-01-04,qux
5,2021-01-05,quux
6,2021-01-06,corge
7,2021-01-07,grault
8,2021-01-08,garply
9,2021-01-09,waldo
10,2021-01-10,fred`,
			hasHeader:    true,
			chunkSize:    3, // Tests chunk processing
			targetColumn: "col3",
			expectedErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create test CSV file
			fileName, err := createTestCSV(tt.csvContent)
			assert.NoError(t, err, "Failed to create test CSV file")
			defer os.Remove(fileName) // Cleanup after test

			// Read data using PredictionCSVParser
			instances, headers, featureTypes, err := PredictionCSVParser(fileName, tt.hasHeader, tt.chunkSize, tt.targetColumn)

			// Check if an error was expected
			if tt.expectedErr {
				assert.Error(t, err, "Expected error but got nil")
				return
			}

			assert.NoError(t, err, "Unexpected error during parsing")
			assert.NotNil(t, instances, "Instances should not be nil")
			assert.NotNil(t, headers, "Headers should not be nil")
			assert.NotNil(t, featureTypes, "Feature types should not be nil")

			// Ensure headers match expectations
			if tt.hasHeader {
				expectedHeaderCount := len(headers)
				assert.Greater(t, expectedHeaderCount, 0, "Headers should not be empty")
			}

			// Ensure instances are parsed correctly
			assert.Greater(t, len(instances), 0, "Instances should be populated")
		})
	}
}
