package model

/*
* Gain calculates Information Gain for a categorical attribute
* How much information will gain if we were to split the class labels by this attribute
*
* input:
*  attributeValues - the attribute values that we are trying to splite the class labels by
*  classLabels - the class labels we are tyring to predict
*
* output:
*  an integer representing the information gain
 */
func Gain(attributeValues []int, classLabels []int) float64 {
	parentEntropy := Entropy(classLabels)
	total := float64(len(classLabels))

	// group class labels by attribute value
	groups := make(map[int][]int)
	for i, attrVal := range attributeValues {
		groups[attrVal] = append(groups[attrVal], classLabels[i])
	}

	// calculate weighted child entropy
	weightedChildEntropy := 0.0
	for _, group := range groups {
		p := float64(len(group)) / total
		weightedChildEntropy += p * Entropy(group)
	}

	return parentEntropy - weightedChildEntropy
}
