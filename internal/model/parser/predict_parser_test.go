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
		// defines a set of test cases with different CSV contents, headers, chunk sizes, and target columns.
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
			name: "CSV without header",
			csvContent: `1,2021-01-01,foo
2,2021-01-02,bar
3,2021-01-03,baz`,
			hasHeader:    false,
			chunkSize:    10,
			targetColumn: "col2",
			expectedErr:  false,
		},
		{
			name: "CSV with missing values",
			csvContent: `col1,col2,col3
1,,foo
,2021-01-02,bar
3,2021-01-03,`,
			hasHeader:    true,
			chunkSize:    10,
			targetColumn: "col3",
			expectedErr:  false,
		},
		{
			name: "CSV with large dataset",
			csvContent: `col1,col2,col3
1,2021-01-01,foo
2,2021-01-02,bar
3,2021-01-03,baz
4,2021-01-04,qux
5,2021-01-05,quux`,
			hasHeader:    true,
			chunkSize:    3,
			targetColumn: "col3",
			expectedErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fileName, err := createTestCSV(tt.csvContent)
			assert.NoError(t, err)
			defer os.Remove(fileName)

			// checks for errors and validates the returned instances, headers, and feature types.
			instances, headers, featureTypes, err := PredictionCSVParser(fileName, tt.hasHeader, tt.chunkSize, tt.targetColumn)
			if tt.expectedErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, instances)
				assert.NotNil(t, headers)
				assert.NotNil(t, featureTypes)
			}
		})
	}
}
