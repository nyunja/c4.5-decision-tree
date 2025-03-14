package model

import "math"

/*
* Entropy - calculates the entropy of a list of class labels
*
* input:
*  a collection with data items
*
* output:
*  a float64 value indicating the entropy level
 */
func Entropy(classes []int) float64 {
	instancesPerClass := make(map[int]int)
	for _, cls := range classes {
		instancesPerClass[cls]++
	}

	total := len(classes)
	if total == 0 {
		return 0
	}

	entropy := 0.0
	for _, numInstancesInClass := range instancesPerClass {
		p := float64(numInstancesInClass) / float64(total)
		entropy -= p * math.Log2(p)
	}
	return entropy
}
