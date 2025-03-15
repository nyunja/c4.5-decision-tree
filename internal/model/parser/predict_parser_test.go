package parser

import (
	"os"
)

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
