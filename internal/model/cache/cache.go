package cache

import (
	"fmt"
	"runtime"
	"sort"
	"sync"
	"time"

	t "github.com/nyunja/c45-decision-tree/internal/model/types"
)

// FeatureCache caches computed values for features to avoid redundant calculations
type FeatureCache struct {
	// For categorical features
	ValueCounts map[string]map[string]int // feature -> value -> count

	// For continuous features
	SortedValues map[string][]float64 // feature -> sorted values

	// For target feature
	TargetCounts map[string]int // target value -> count

	// Mutex for thread safety
	Mu sync.RWMutex
}

// NewFeatureCache creates a new feature cache
func NewFeatureCache() *FeatureCache {
	return &FeatureCache{
		ValueCounts:  make(map[string]map[string]int),
		SortedValues: make(map[string][]float64),
		TargetCounts: make(map[string]int),
	}
}

// PrecomputeFeatureValues precomputes and caches values for all features
func (fc *FeatureCache) PrecomputeFeatureValues(instances []t.Instance, features []string, targetFeature string, featureTypes map[string]string) {
	// Count target values
	for _, instance := range instances {
		targetVal := fmt.Sprintf("%v", instance[targetFeature])
		fc.Mu.Lock()
		fc.TargetCounts[targetVal]++
		fc.Mu.Unlock()
	}

	// Process features in parallel
	var wg sync.WaitGroup
	featureChan := make(chan string, len(features))

	numWorkers := runtime.NumCPU()
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for feature := range featureChan {
				featureType := featureTypes[feature]

				if featureType == "numerical" || featureType == "date" || featureType == "timestamp" {
					// For continuous features, collect and sort values
					values := make([]float64, 0, len(instances))
					valueSet := make(map[float64]bool)

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

						if !valueSet[floatVal] {
							valueSet[floatVal] = true
							values = append(values, floatVal)
						}
					}

					// Sort values
					sort.Float64s(values)

					// Sample values if there are too many
					if len(values) > 100 {
						sampledValues := make([]float64, 100)
						step := float64(len(values)) / 100.0
						for i := 0; i < 100; i++ {
							index := int(float64(i) * step)
							if index >= len(values) {
								index = len(values) - 1
							}
							sampledValues[i] = values[index]
						}
						values = sampledValues
					}

					fc.Mu.Lock()
					fc.SortedValues[feature] = values
					fc.Mu.Unlock()
				} else {
					// For categorical features, count occurrences of each value
					valueCounts := make(map[string]int)

					for _, instance := range instances {
						val := instance[feature]
						if val == nil {
							continue
						}

						strVal := fmt.Sprintf("%v", val)
						valueCounts[strVal]++
					}

					fc.Mu.Lock()
					fc.ValueCounts[feature] = valueCounts
					fc.Mu.Unlock()
				}
			}
		}()
	}

	// Send features to workers
	for _, feature := range features {
		featureChan <- feature
	}
	close(featureChan)

	// Wait for all workers to finish
	wg.Wait()
}
