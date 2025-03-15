// split/continuous.go
package split

import (
	"fmt"
	"time"

	"github.com/nyunja/c4.5-decision-tree/internal/model/counter"
	"github.com/nyunja/c4.5-decision-tree/internal/model/entropy"
)

// EvaluateContinuousFeature evaluates a continuous feature for the best split point
func EvaluateContinuousFeature(feature string, context SplitContext) SplitResult {
	context.Cache.Mu.RLock()
	sortedValues := context.Cache.SortedValues[feature]
	context.Cache.Mu.RUnlock()

	if len(sortedValues) <= 1 {
		return SplitResult{
			Feature:      feature,
			GainRatio:    0,
			IsContinuous: true,
		}
	}

	// Find the best split point
	bestGainRatio := 0.0
	bestThreshold := 0.0

	for i := 0; i < len(sortedValues)-1; i++ {
		threshold := (sortedValues[i] + sortedValues[i+1]) / 2
		gainRatio, leftCount, rightCount := EvaluateThreshold(feature, threshold, context)

		if gainRatio > bestGainRatio && IsSplitValid(leftCount, rightCount) {
			bestGainRatio = gainRatio
			bestThreshold = threshold
		}
	}

	return SplitResult{
		Feature:      feature,
		GainRatio:    bestGainRatio,
		IsContinuous: true,
		Threshold:    bestThreshold,
	}
}

// EvaluateThreshold evaluates a specific threshold for a continuous feature
func EvaluateThreshold(feature string, threshold float64, context SplitContext) (float64, int, int) {
	// Count instances on each side of the threshold
	leftCounter := counter.NewClassCounter()
	rightCounter := counter.NewClassCounter()

	for _, instance := range context.Instances {
		val := instance[feature]
		if val == nil {
			continue
		}

		floatVal, ok := ExtractNumericValue(val)
		if !ok {
			continue
		}

		targetVal := fmt.Sprintf("%v", instance[context.TargetFeature])
		if floatVal <= threshold {
			leftCounter.Add(targetVal)
		} else {
			rightCounter.Add(targetVal)
		}
	}

	// Calculate information gain
	leftProb := float64(leftCounter.Total) / float64(len(context.Instances))
	rightProb := float64(rightCounter.Total) / float64(len(context.Instances))

	infoGain := context.BaseEntropy
	if leftCounter.Total > 0 {
		infoGain -= leftProb * leftCounter.GetEntropy()
	}
	if rightCounter.Total > 0 {
		infoGain -= rightProb * rightCounter.GetEntropy()
	}

	// Calculate the gain ratio
	gainRatio := entropy.GainRatio(leftProb, rightProb, infoGain)

	return gainRatio, leftCounter.Total, rightCounter.Total
}

// ExtractNumericValue converts various types to float64 for comparison
func ExtractNumericValue(val interface{}) (float64, bool) {
	switch v := val.(type) {
	case float64:
		return v, true
	case int:
		return float64(v), true
	case time.Time:
		return float64(v.Unix()), true
	default:
		return 0, false
	}
}

// IsSplitValid checks if a split is valid based on left and right counts
func IsSplitValid(leftCount, rightCount int) bool {
	// You can add minimum size constraints here if needed
	return leftCount > 0 && rightCount > 0
}
