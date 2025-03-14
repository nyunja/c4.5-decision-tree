package entropy

import (
	"fmt"
	t "github.com/nyunja/c45-decision-tree/internal/model/types"
	"github.com/nyunja/c45-decision-tree/internal/model/counter"
)

// CalculateEntropy calculates the entropy of a dataset
func CalculateEntropy(instances []t.Instance, targetFeature string) float64 {
	if len(instances) == 0 {
		return 0
	}

	counter := counter.NewClassCounter()
	for _, instance := range instances {
		targetVal := fmt.Sprintf("%v", instance[targetFeature])
		counter.Add(targetVal)
	}

	return counter.GetEntropy()
}
