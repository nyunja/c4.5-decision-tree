package internal

/*
* GainRatio - calculates the gain ratio for an attribute with respect to class labels
*
* input:
*  a slice of integers representing the values of an attribute
*  a slice of integers representing the class labels corresponding to the attribute values
*
* output:
*  a float64 value representing the gain ratio of the attribute with respect to the class labels
 */

func GainRatio(attributeValues []int, classLabels []int) float64 {
	gain := Gain(attributeValues, classLabels)
	splitInfo := SplitInfo(attributeValues)

	// edge cases (no split, perfect split, or very small splitInfo)
	if splitInfo < 0.0001 || gain <= 0 {
		return 0.0
	}
	return gain / splitInfo
}
