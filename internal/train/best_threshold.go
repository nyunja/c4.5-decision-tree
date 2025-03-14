package model

import (
	"math"
	"sort"
)

func FindBestThreshold(values []float64, classes []int) (float64, float64) {
	// Sort values
	sort.Float64s(values)

	bestThreshold := 0.0
	bestGainRatio := -1.0

	for i := 0; i < len(values)-1; i++ {
		threshold := (values[i] + values[i+1]) / 2

		// Split dataset based on threshold
		leftClasses := []int{}
		rightClasses := []int{}
		for j, val := range values {
			if val <= threshold {
				leftClasses = append(leftClasses, classes[j])
			} else {
				rightClasses = append(rightClasses, classes[j])
			}
		}

		// Calculate Gain Ratio for this split
		parentEntropy := Entropy(classes)
		leftEntropy := Entropy(leftClasses)
		rightEntropy := Entropy(rightClasses)
		total := float64(len(classes))
		leftP := float64(len(leftClasses)) / total
		rightP := float64(len(rightClasses)) / total
		gain := parentEntropy - leftP*leftEntropy - rightP*rightEntropy

		splitInfo := -leftP*math.Log2(leftP) - rightP*math.Log2(rightP)
		if splitInfo <= 0 || gain <= 0 {
			continue
		}
		gainRatio := gain / splitInfo

		// Update best threshold if needed
		if gainRatio > bestGainRatio {
			bestGainRatio = gainRatio
			bestThreshold = threshold
		}
	}

	return bestThreshold, bestGainRatio
}
