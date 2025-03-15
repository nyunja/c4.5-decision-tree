package split

import (
	"fmt"
	"math"
	"runtime"
	"sync"
	"time"

	"github.com/nyunja/c4.5-decision-tree/internal/model/cache"
	"github.com/nyunja/c4.5-decision-tree/internal/model/counter"
	"github.com/nyunja/c4.5-decision-tree/internal/model/entropy"
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

// FindBestSplit finds the best feature and split point using the feature cache
func FindBestSplit(instances []t.Instance, features []string, targetFeature string, featureTypes map[string]string, excludedFeatures map[string]bool, cache *cache.FeatureCache) (string, interface{}, bool, float64) {
	if len(instances) == 0 || len(features) == 0 {
		return "", nil, false, 0
	}

	baseEntropy := entropy.CalculateEntropy(instances, targetFeature)

	// Use a worker pool to evaluate features in parallel
	numWorkers := runtime.NumCPU()
	featuresChan := make(chan string, len(features))
	resultsChan := make(chan SplitResult, len(features))

	// Start worker goroutines
	var wg sync.WaitGroup
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for feature := range featuresChan {
				// Skip excluded features and the target feature
				if excludedFeatures[feature] || feature == targetFeature {
					continue
				}

				featureType := featureTypes[feature]
				if featureType == "numerical" || featureType == "date" || featureType == "timestamp" {
					// Handle continuous attributes
					cache.Mu.RLock()
					sortedValues := cache.SortedValues[feature]
					cache.Mu.RUnlock()

					if len(sortedValues) <= 1 {
						resultsChan <- SplitResult{
							Feature:      feature,
							GainRatio:    0,
							IsContinuous: true,
						}
						continue
					}

					// Find the best split point
					bestGainRatio := 0.0
					bestThreshold := 0.0

					for i := 0; i < len(sortedValues)-1; i++ {
						threshold := (sortedValues[i] + sortedValues[i+1]) / 2

						// Count instances on each side of the threshold
						leftCounter := counter.NewClassCounter()
						rightCounter := counter.NewClassCounter()

						for _, instance := range instances {
							val := instance[feature]
							if val == nil {
								continue
							}

							var floatVal float64
							switch v := val.(type) {
							case float64:
								floatVal = v
							case int:
								floatVal = float64(v)
							case time.Time:
								floatVal = float64(v.Unix())
							default:
								continue
							}

							targetVal := fmt.Sprintf("%v", instance[targetFeature])
							if floatVal <= threshold {
								leftCounter.Add(targetVal)
							} else {
								rightCounter.Add(targetVal)
							}
						}

						// Calculate information gain
						leftProb := float64(leftCounter.Total) / float64(len(instances))
						rightProb := float64(rightCounter.Total) / float64(len(instances))

						infoGain := baseEntropy
						if leftCounter.Total > 0 {
							infoGain -= leftProb * leftCounter.GetEntropy()
						}
						if rightCounter.Total > 0 {
							infoGain -= rightProb * rightCounter.GetEntropy()
						}

						// Calculate split info for gain ratio
						splitInfo := 0.0
						if leftProb > 0 {
							splitInfo -= leftProb * math.Log2(leftProb)
						}
						if rightProb > 0 {
							splitInfo -= rightProb * math.Log2(rightProb)
						}

						// Calculate gain ratio
						gainRatio := 0.0
						if splitInfo > 0 {
							gainRatio = infoGain / splitInfo
						}

						if gainRatio > bestGainRatio {
							bestGainRatio = gainRatio
							bestThreshold = threshold
						}
					}

					resultsChan <- SplitResult{
						Feature:      feature,
						GainRatio:    bestGainRatio,
						IsContinuous: true,
						Threshold:    bestThreshold,
					}
				} else {
					// Handle categorical attributes
					cache.Mu.RLock()
					valueCounts := cache.ValueCounts[feature]
					cache.Mu.RUnlock()

					if len(valueCounts) == 0 {
						resultsChan <- SplitResult{
							Feature:      feature,
							GainRatio:    0,
							IsContinuous: false,
						}
						continue
					}

					// Calculate information gain
					infoGain := baseEntropy
					splitInfo := 0.0

					// Create counters for each value
					valueCounters := make(map[string]*counter.ClassCounter)
					for value := range valueCounts {
						valueCounters[value] = counter.NewClassCounter()
					}

					// Count target values for each feature value
					for _, instance := range instances {
						val := instance[feature]
						if val == nil {
							continue
						}

						strVal := fmt.Sprintf("%v", val)
						targetVal := fmt.Sprintf("%v", instance[targetFeature])

						if counter, ok := valueCounters[strVal]; ok {
							counter.Add(targetVal)
						}
					}

					// Calculate information gain and split info
					for _, counter := range valueCounters {
						prob := float64(counter.Total) / float64(len(instances))
						if prob > 0 {
							infoGain -= prob * counter.GetEntropy()
							splitInfo -= prob * math.Log2(prob)
						}
					}

					// Calculate gain ratio
					gainRatio := 0.0
					if splitInfo > 0 {
						gainRatio = infoGain / splitInfo
					}

					resultsChan <- SplitResult{
						Feature:      feature,
						GainRatio:    gainRatio,
						IsContinuous: false,
					}
				}
			}
		}()
	}

	// Send features to workers
	go func() {
		for _, feature := range features {
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
	bestResult := SplitResult{GainRatio: -1}
	for result := range resultsChan {
		if result.GainRatio > bestResult.GainRatio {
			bestResult = result
		}
	}

	if bestResult.GainRatio <= 0 {
		return "", nil, false, 0
	}

	return bestResult.Feature, bestResult.Value, bestResult.IsContinuous, bestResult.Threshold
}
