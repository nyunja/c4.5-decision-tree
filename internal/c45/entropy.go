package c45

import "math"

// Calculates entropy of given dataset
func Entropy(data [][]string, labelIndex int) float64 {
	counts := make(map[string]float64)
	total := float64(len(data))

	for _, row := range data {
        label := row[labelIndex]
        counts[label]++
    }
	entropy := 0.0
	for _, count := range counts {
        prob := count / total
        entropy -= prob * math.Log2(prob)
    }
	return entropy
}