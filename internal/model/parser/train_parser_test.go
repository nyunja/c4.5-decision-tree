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
