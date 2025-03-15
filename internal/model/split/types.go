// split/types.go
package split

import (
	"github.com/nyunja/c4.5-decision-tree/internal/model/cache"
	t "github.com/nyunja/c4.5-decision-tree/internal/model/types"
)

// SplitResult holds the result of a feature split evaluation
type SplitResult struct {
	Feature      string
	Value        interface{}
	GainRatio    float64
	IsContinuous bool
	Threshold    float64
}

// SplitContext holds the context data needed for split evaluation
type SplitContext struct {
	Instances        []t.Instance
	Features         []string
	TargetFeature    string
	FeatureTypes     map[string]string
	ExcludedFeatures map[string]bool
	Cache            *cache.FeatureCache
	BaseEntropy      float64
}
