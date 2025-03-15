package parser

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStreamingCSVParser_NoHeader(t *testing.T) {
	csvContent := `1,Alice,30,2021-01-01
2,Bob,25,2021-02-01
3,Charlie,35,2021-03-01
`
	files, err := createTestCSV(csvContent)
	assert.NoError(t, err)
	defer os.Remove(files)

	instances, headers, featureTypes, err := StreamingCSVParser(files, false, 100, "col1")
	assert.NoError(t, err)
	assert.Equal(t, 3, len(instances))
	assert.Equal(t, []string{"col0", "col1", "col2", "col3"}, headers)
	assert.Equal(t, map[string]string{
		"col0": "numerical",
		"col1": "categorical",
		"col2": "numerical",
		"col3": "date",
	}, featureTypes)
}

func TestStreamingCSVParser_Emptyfiles(t *testing.T) {
	files, err := createTestCSV("")
	assert.NoError(t, err)
	defer os.Remove(files)

	instances, headers, featureTypes, err := StreamingCSVParser(files, true, 100, "age")
	assert.Error(t, err)
	assert.Nil(t, instances)
	assert.Nil(t, headers)
	assert.Nil(t, featureTypes)
}

func TestStreamingCSVParser_Invalidfiles(t *testing.T) {
	instances, headers, featureTypes, err := StreamingCSVParser("invalid_files.csv", true, 100, "age")
	assert.Error(t, err)
	assert.Nil(t, instances)
	assert.Nil(t, headers)
	assert.Nil(t, featureTypes)
}

func TestStreamingCSVParser(t *testing.T) {
	csvContent := `id,name,age,date
1,Alice,30.00,2021/01/01
2,Bob,25.00,2021/02/01
3,Charlie,35.00,2021/03/01
`
	files, err := createTestCSV(csvContent)
	assert.NoError(t, err)
	defer func() {
		if files != "" {
			os.Remove(files)
		}
	}()
	// defer os.Remove(files) // remove temporary file after reading from it

	instances, headers, featureTypes, err := StreamingCSVParser(files, true, 100, "date")
	assert.NoError(t, err)
	assert.Equal(t, 3, len(instances))
	assert.Equal(t, []string{"id", "name", "age", "date"}, headers)
	assert.Equal(t, map[string]string{
		"id":   "numerical",
		"name": "categorical",
		"age":  "numerical",
		"date": "date",
	}, featureTypes)
}
