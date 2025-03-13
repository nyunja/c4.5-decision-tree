package internal

import "math"

/*
* SplitInfo - calculates the split information for an attribute
*
* input:
*  a slice of integers representing the values of an attribute
*
* output:
*  a float64 value representing the split information (entropy) for the attribute
 */

func SplitInfo(attributeValues []int) float64 {
	counts := make(map[int]int)
	for _, val := range attributeValues {
		counts[val]++
	}

	total := float64(len(attributeValues))
	splitEntropy := 0.0
	for _, cnt := range counts {
		p := float64(cnt) / total
		splitEntropy -= p * math.Log2(p)
	}

	return splitEntropy
}
