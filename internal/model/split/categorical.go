// split/categorical.go
package split

import (
	"fmt"

	"github.com/nyunja/c4.5-decision-tree/internal/model/counter"
	"github.com/nyunja/c4.5-decision-tree/internal/model/entropy"
)

// EvaluateCategoricalFeature evaluates a categorical feature
func EvaluateCategoricalFeature(feature string, context SplitContext) SplitResult {
	context.Cache.Mu.RLock()
	valueCounts := context.Cache.ValueCounts[feature]
	context.Cache.Mu.RUnlock()

	if len(valueCounts) == 0 {
		return SplitResult{
			Feature:      feature,
			GainRatio:    0,
			IsContinuous: false,
		}
	}

	// Calculate information gain
	infoGain := context.BaseEntropy
	splitInfo := 0.0

	// Create counters for each value
	valueCounters := CreateValueCounters(feature, context, valueCounts)

	// Calculate information gain and split info
	for _, counter := range valueCounters {
		infoGain, splitInfo = entropy.GainInfoAndSplitInfo(counter, context.Instances, infoGain, splitInfo)
	}

	// Calculate gain ratio
	gainRatio := 0.0
	if splitInfo > 0 {
		gainRatio = infoGain / splitInfo
	}

	return SplitResult{
		Feature:      feature,
		GainRatio:    gainRatio,
		IsContinuous: false,
	}
}

// CreateValueCounters creates class counters for each value of a categorical feature
func CreateValueCounters(feature string, context SplitContext,
	valueCounts map[string]int,
) map[string]*counter.ClassCounter {
	// Create counters for each value
	valueCounters := make(map[string]*counter.ClassCounter)
	for value := range valueCounts {
		valueCounters[value] = counter.NewClassCounter()
	}

	// Count target values for each feature value
	for _, instance := range context.Instances {
		val := instance[feature]
		if val == nil {
			continue
		}

		strVal := fmt.Sprintf("%v", val)
		targetVal := fmt.Sprintf("%v", instance[context.TargetFeature])

		if counter, ok := valueCounters[strVal]; ok {
			counter.Add(targetVal)
		}
	}

	return valueCounters
}
