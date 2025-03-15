package entropy

import (
	"fmt"
	"github.com/nyunja/c4.5-decision-tree/internal/model/counter"
	t "github.com/nyunja/c4.5-decision-tree/internal/model/types"
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
