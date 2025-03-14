package model

import (
	"fmt"
	"math"
	"strconv"

	"github.com/nyunja/c45-decision-tree/internal"
)

/*
* BuildTree - constructs a decision tree from a dataset
* input:
*  dataset: Matrix of feature values
*  classLabels: Array of class labels
*  attributes: Indices of attributes to consider
*  featureNames: Names of the features
* output:
*  A JSONTreeNode representing the decision tree
 */
func BuildTree(dataset [][]float64, classLabels []int, attributes []int, featureNames []string) *internal.JSONTreeNode {
	// Base cases
	if len(attributes) == 0 || allSameClass(classLabels) {
		return createLeafNode(classLabels)
	}

	bestAttr := -1
	bestGainRatio := -1.0
	bestThreshold := 0.0

	for _, attr := range attributes {
		var gainRatio float64
		var threshold float64

		if isContinuous(dataset, attr) {
			threshold, gainRatio = FindBestThreshold(getColumn(dataset, attr), classLabels)
		} else {
			attrValues := make([]int, len(dataset))
			for i, row := range dataset {
				attrValues[i] = int(row[attr])
			}
			gainRatio = GainRatio(attrValues, classLabels)
		}

		if gainRatio > bestGainRatio {
			bestGainRatio = gainRatio
			bestAttr = attr
			bestThreshold = threshold
		}
	}

	// If no good split found, create leaf node
	if bestGainRatio <= 0 {
		return createLeafNode(classLabels)
	}

	// Split dataset
	if isContinuous(dataset, bestAttr) {
		leftDataset, rightDataset, leftLabels, rightLabels := splitDataset(dataset, classLabels, bestAttr, bestThreshold)

		// Check if split produced empty datasets
		if len(leftDataset) == 0 || len(rightDataset) == 0 {
			return createLeafNode(classLabels)
		}

		node := &internal.JSONTreeNode{
			SplitFeature: featureNames[bestAttr],
			SplitType:    "continuous",
			Children:     make(map[string]*internal.JSONTreeNode),
		}

		// true represents <= threshold, false represents > threshold
		node.Children["<= "+fmt.Sprintf("%.2f", bestThreshold)] = BuildTree(leftDataset, leftLabels, attributes, featureNames)
		node.Children["> "+fmt.Sprintf("%.2f", bestThreshold)] = BuildTree(rightDataset, rightLabels, attributes, featureNames)
		return node
	} else {
		subsets := splitDatasetByAttribute(dataset, classLabels, bestAttr)

		// If any subset is empty, create a leaf node
		if len(subsets) <= 1 {
			return createLeafNode(classLabels)
		}

		node := &internal.JSONTreeNode{
			SplitFeature: featureNames[bestAttr],
			SplitType:    "categorical",
			Children:     make(map[string]*internal.JSONTreeNode),
		}

		newAttributes := removeAttribute(attributes, bestAttr)
		for value, subset := range subsets {
			node.Children[fmt.Sprintf("%d", value)] = BuildTree(subset.Dataset, subset.Labels, newAttributes, featureNames)
		}
		return node
	}
}

/*
* createLeafNode - creates a leaf node for the decision tree
* input:
*  labels: Array of class labels
* output:
*  A JSONTreeNode representing a leaf node with class distribution
 */
func createLeafNode(labels []int) *internal.JSONTreeNode {
	classDistribution := make(map[string]int)
	for _, label := range labels {
		classDistribution[fmt.Sprintf("%d", label)]++
	}

	return &internal.JSONTreeNode{
		SplitFeature:      "",
		SplitType:         "",
		Children:          nil,
		ClassDistribution: classDistribution,
	}
}

/*
* majorityClass - determines the majority class in a set of labels
* input:
*  labels: Array of class labels
* output:
*  The majority class label
 */
func MajorityClass(labels []int) int {
	counts := make(map[int]int)

	for _, label := range labels {
		counts[label]++
	}

	majorityClass := -1
	maxCount := -1

	for class, count := range counts {
		if count > maxCount {
			maxCount = count
			majorityClass = class
		}
	}

	return majorityClass
}

/*
* allSameClass - checks if all labels belong to the same class
* input:
*  labels: Array of class labels
* output:
*  True if all labels are the same, false otherwise
 */
func allSameClass(labels []int) bool {
	if len(labels) == 0 {
		return true
	}

	firstClass := labels[0]
	for _, label := range labels {
		if label != firstClass {
			return false
		}
	}

	return true
}

/*
* isContinuous - checks if an attribute is continuous
* input:
*  dataset: Matrix of feature values
*  attr: Index of the attribute to check
* output:
*  True if the attribute is continuous, false otherwise
 */
func isContinuous(dataset [][]float64, attr int) bool {
	// Check if column contains any non-integer values
	for _, row := range dataset {
		// Skip if out of bounds
		if attr >= len(row) {
			continue
		}

		// If value is not equal to its integer part, it's continuous
		if math.Floor(row[attr]) != row[attr] {
			return true
		}
	}

	// Find number of distinct values
	values := make(map[float64]bool)
	for _, row := range dataset {
		if attr < len(row) {
			values[row[attr]] = true
		}
	}

	// If more than half of rows have unique values, consider it continuous
	return len(values) > len(dataset)/2
}

/*
* getColumn - extracts a column from a dataset
* input:
*  dataset: Matrix of feature values
*  attr: Index of the column to extract
* output:
*  An array representing the extracted column
 */
func getColumn(dataset [][]float64, attr int) []float64 {
	column := make([]float64, len(dataset))

	for i, row := range dataset {
		if attr < len(row) {
			column[i] = row[attr]
		}
	}

	return column
}

/*
* getStringColumn - extracts a column from a string matrix
* input:
*  matrix: Matrix of string values
*  col: Index of the column to extract
* output:
*  An array representing the extracted column
 */
func GetStringColumn(matrix [][]string, col int) []string {
	column := make([]string, len(matrix))
	for i, row := range matrix {
		if col < len(row) {
			column[i] = row[col]
		}
	}
	return column
}

/*
* isNumeric - checks if a string can be parsed as a number
* input:
*  s: The string to check
* output:
*  True if the string is numeric, false otherwise
 */
func IsNumeric(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

/*
* splitDataset - splits a dataset based on a threshold
* input:
*  dataset: Matrix of feature values
*  labels: Array of class labels
*  attr: Index of the attribute to split on
*  threshold: Threshold value
* output:
*  Left and right datasets and corresponding labels
 */
func splitDataset(dataset [][]float64, labels []int, attr int, threshold float64) ([][]float64, [][]float64, []int, []int) {
	leftDataset := [][]float64{}
	rightDataset := [][]float64{}
	leftLabels := []int{}
	rightLabels := []int{}

	for i, row := range dataset {
		if attr < len(row) && row[attr] <= threshold {
			leftDataset = append(leftDataset, row)
			leftLabels = append(leftLabels, labels[i])
		} else {
			rightDataset = append(rightDataset, row)
			rightLabels = append(rightLabels, labels[i])
		}
	}

	return leftDataset, rightDataset, leftLabels, rightLabels
}

/*
* splitDatasetByAttribute - splits a dataset based on categorical attribute values
* input:
*  dataset: Matrix of feature values
*  labels: Array of class labels
*  attr: Index of the attribute to split on
* output:
*  A map of subsets, keyed by attribute values
 */
func splitDatasetByAttribute(dataset [][]float64, labels []int, attr int) map[interface{}]internal.Subset {
	subsets := make(map[interface{}]internal.Subset)

	for i, row := range dataset {
		if attr >= len(row) {
			continue
		}

		value := int(row[attr]) // Cast to int for categorical values
		subset, exists := subsets[value]

		if !exists {
			subset = internal.Subset{
				Dataset: [][]float64{},
				Labels:  []int{},
			}
		}

		subset.Dataset = append(subset.Dataset, row)
		subset.Labels = append(subset.Labels, labels[i])
		subsets[value] = subset
	}

	return subsets
}

/*
* removeAttribute - removes an attribute from a list of attributes
* input:
*  attributes: List of attribute indices
*  attr: Index of the attribute to remove
* output:
*  The updated list of attribute indices
 */
func removeAttribute(attributes []int, attr int) []int {
	result := []int{}

	for _, a := range attributes {
		if a != attr {
			result = append(result, a)
		}
	}

	return result
}
