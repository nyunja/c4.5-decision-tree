package model

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

func Train(inputFile string, outputFile string, targetColumn string) error {
	// Read CSV file
	file, err := os.Open(inputFile)
	if err != nil {
		return err
	}
	defer file.Close()

	csvReader := csv.NewReader(file)
	records, err := csvReader.ReadAll()
	if err != nil {
		return err
	}

	// Get header row
	headers := records[0]

	// Find index of target column
	var targetIndex int
	for i, header := range headers {
		if header == targetColumn {
			targetIndex = i
			break
		}
	}

	// Extract features and class labels
	features := make([][]string, len(records)-1)
	classLabels := []string{}
	for i, record := range records[1:] {
		features[i] = record[:targetIndex]
		features[i] = append(features[i], record[targetIndex+1:]...)
		classLabels = append(classLabels, record[targetIndex])
	}

	// Convert string features to ordinal integers
	ordinalFeatures := make([][]float64, len(classLabels))
	for i, featureRow := range features {
		ordinalRow := []float64{}
		for j, feature := range featureRow {
			if IsNumeric(feature) {
				num, err := strconv.ParseFloat(feature, 64)
				if err != nil {
					return err
				}
				ordinalRow = append(ordinalRow, num)
			} else {
				_, ordinals := StringArrayToOrdinal(GetStringColumn(features, j))
				ordinalRow = append(ordinalRow, float64(ordinals[i]))
			}
		}
		ordinalFeatures[i] = ordinalRow
	}

	// Convert class labels to ordinal integers
	_, classOrdinals := StringArrayToOrdinal(classLabels)

	// Available attributes (column indices)
	attributes := make([]int, len(headers)-2) // Exclude Loan_ID and target column
	for i := range attributes {
		attributes[i] = i
	}

	// Feature names
	featureNames := headers[:targetIndex]
	featureNames = append(featureNames, headers[targetIndex+1:]...)

	tree := BuildTree(ordinalFeatures, classOrdinals, attributes, featureNames)

	// Print the tree structure
	jsonTree, err := json.MarshalIndent(tree, "", "  ")
	if err != nil {
		return err
	}
	fmt.Println(string(jsonTree))

	// Write to file
	fileOut, err := os.OpenFile(outputFile, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0o644)
	if err != nil {
		return err
	}
	defer fileOut.Close() // Close the file when done

	// Directly encode the tree to the file with indentation
	encoder := json.NewEncoder(fileOut)
	encoder.SetIndent("", "  ") // Set indentation for pretty printing
	err = encoder.Encode(tree)
	if err != nil {
		return err
	}

	return nil
}
