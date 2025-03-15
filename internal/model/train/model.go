package train

import (
	"fmt"

	"github.com/nyunja/c4.5-decision-tree/internal/model/cache"
	t "github.com/nyunja/c4.5-decision-tree/internal/model/types"
)

// TrainModel trains a C4.5 decision tree model with optimizations for large datasets
func Train(instances []t.Instance, headers []string, targetFeature string, featureTypes map[string]string, excludeColumns []string, maxDepth int) (*t.Model, error) {
	// Validate inputs
	if len(instances) == 0 {
		return nil, fmt.Errorf("no instances provided for training")
	}
	if _, ok := featureTypes[targetFeature]; !ok {
		return nil, fmt.Errorf("target feature '%s' not found in feature types", targetFeature)
	}

	// Create a map of excluded features for faster lookup
	excludedFeatures := make(map[string]bool, len(excludeColumns))
	for _, feature := range excludeColumns {
		excludedFeatures[feature] = true
	}

	// Remove target feature from the list of features used for splitting
	features := make([]string, 0, len(headers))
	for _, feature := range headers {
		if feature != targetFeature && !excludedFeatures[feature] {
			features = append(features, feature)
		}
	}

	// Precompute feature values
	cache := cache.NewFeatureCache()
	cache.PrecomputeFeatureValues(instances, features, targetFeature, featureTypes)

	// Train the decision tree
	root := C45(instances, features, targetFeature, featureTypes, excludedFeatures, 5, maxDepth, cache) // minInstancesPerLeaf = 5

	// Create and return the model
	model := &t.Model{
		Root:         root,
		FeatureTypes: featureTypes,
		FeatureNames: headers,
		TargetName:   targetFeature,
	}

	return model, nil
}
