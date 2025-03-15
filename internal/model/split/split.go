// split/split.go
package split

import (
	"runtime"
	"sync"

	"github.com/nyunja/c4.5-decision-tree/internal/model/cache"
	"github.com/nyunja/c4.5-decision-tree/internal/model/entropy"
	t "github.com/nyunja/c4.5-decision-tree/internal/model/types"
)

// FindBestSplit finds the best feature and split point using the feature cache
func FindBestSplit(instances []t.Instance, features []string, targetFeature string,
	featureTypes map[string]string, excludedFeatures map[string]bool,
	cache *cache.FeatureCache,
) (string, interface{}, bool, float64) {
	if len(instances) == 0 || len(features) == 0 {
		return "", nil, false, 0
	}

	context := CreateSplitContext(instances, features, targetFeature, featureTypes, excludedFeatures, cache)

	// Start the parallel evaluation process
	result := EvaluateFeaturesInParallel(context)

	if result.GainRatio <= 0 {
		return "", nil, false, 0
	}

	return result.Feature, result.Value, result.IsContinuous, result.Threshold
}

// CreateSplitContext prepares the context needed for split evaluation
func CreateSplitContext(instances []t.Instance, features []string, targetFeature string,
	featureTypes map[string]string, excludedFeatures map[string]bool,
	cache *cache.FeatureCache,
) SplitContext {
	baseEntropy := entropy.CalculateEntropy(instances, targetFeature)

	return SplitContext{
		Instances:        instances,
		Features:         features,
		TargetFeature:    targetFeature,
		FeatureTypes:     featureTypes,
		ExcludedFeatures: excludedFeatures,
		Cache:            cache,
		BaseEntropy:      baseEntropy,
	}
}

// EvaluateFeaturesInParallel evaluates all features in parallel using worker pool
func EvaluateFeaturesInParallel(context SplitContext) SplitResult {
	numWorkers := runtime.NumCPU()
	featuresChan := make(chan string, len(context.Features))
	resultsChan := make(chan SplitResult, len(context.Features))

	// Start worker goroutines
	var wg sync.WaitGroup
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go FeatureEvaluationWorker(&wg, featuresChan, resultsChan, context)
	}

	// Send features to workers
	go func() {
		for _, feature := range context.Features {
			featuresChan <- feature
		}
		close(featuresChan)
	}()

	// Collect results
	go func() {
		wg.Wait()
		close(resultsChan)
	}()

	// Find the best split
	return FindBestResult(resultsChan)
}

// FindBestResult collects all results and finds the best one
func FindBestResult(resultsChan <-chan SplitResult) SplitResult {
	bestResult := SplitResult{GainRatio: -1}
	for result := range resultsChan {
		if result.GainRatio > bestResult.GainRatio {
			bestResult = result
		}
	}
	return bestResult
}

// FeatureEvaluationWorker evaluates features from the channel
func FeatureEvaluationWorker(wg *sync.WaitGroup, featuresChan <-chan string,
	resultsChan chan<- SplitResult, context SplitContext,
) {
	defer wg.Done()

	for feature := range featuresChan {
		// Skip excluded features and the target feature
		if context.ExcludedFeatures[feature] || feature == context.TargetFeature {
			continue
		}

		featureType := context.FeatureTypes[feature]
		if IsContinuousFeature(featureType) {
			resultsChan <- EvaluateContinuousFeature(feature, context)
		} else {
			resultsChan <- EvaluateCategoricalFeature(feature, context)
		}
	}
}

// IsContinuousFeature checks if a feature type is continuous
func IsContinuousFeature(featureType string) bool {
	return featureType == "numerical" || featureType == "date" || featureType == "timestamp"
}
