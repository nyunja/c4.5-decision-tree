package util

import (
	"fmt"
	"os"
)

func ReadCSVFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error opening file: %v", err)
	}
	defer file.Close()
}
