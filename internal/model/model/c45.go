package model

import (
	"fmt"
	"time"

	"github.com/nyunja/c45-decision-tree/internal/model/cache"
	"github.com/nyunja/c45-decision-tree/internal/model/counter"
	ndp "github.com/nyunja/c45-decision-tree/internal/model/node"
	"github.com/nyunja/c45-decision-tree/internal/model/split"
	t "github.com/nyunja/c45-decision-tree/internal/model/types"
	"github.com/nyunja/c45-decision-tree/internal/model/utils"
)

// C45 implements the C4.5 algorithm with optimizations for large datasets
func C45(instances []t.Instance, features []string, targetFeature string, featureTypes map[string]string, excludedFeatures map[string]bool, minInstancesPerLeaf int, maxDepth int, cache *cache.FeatureCache) *t.Node {
	// Base case 1: If there are no instances, return a leaf node
	if len(instances) == 0 {
		return &t.Node{IsLeaf: true}
	}

	// Base case 2: If maximum depth reached, return a leaf node
	if maxDepth <= 0 {
		return &t.Node{
			IsLeaf: true,
			Class:  ndp.GetMajorityClass(instances, targetFeature),
		}
	}

	// Check if all instances have the same target value
	counter := counter.NewClassCounter()
	for _, instance := range instances {
		targetVal := fmt.Sprintf("%v", instance[targetFeature])
		counter.Add(targetVal)
	}

	// Base case 3: If all instances belong to the same class
	if len(counter.Counts) == 1 {
		return &t.Node{
			IsLeaf: true,
			Class:  counter.GetMajorityClass(),
		}
	}

	// Base case 4: If there are no features left or fewer than minInstancesPerLeaf
	if len(features) == 0 || len(instances) < minInstancesPerLeaf {
		return &t.Node{
			IsLeaf: true,
			Class:  counter.GetMajorityClass(),
		}
	}

	// Find the best feature to split on
	bestFeature, _, isContinuous, threshold := split.FindBestSplit(instances, features, targetFeature, featureTypes, excludedFeatures, cache)

	// If no good split found, return a leaf node
	if bestFeature == "" {
		return &t.Node{
			IsLeaf: true,
			Class:  counter.GetMajorityClass(),
		}
	}

	// Create a decision node
	node := &t.Node{
		Feature:    bestFeature,
		IsLeaf:     false,
		Continuous: isContinuous,
		Threshold:  threshold,
	}

	if isContinuous {
		// Split on continuous feature
		leftInstances := utils.FilterInstances(instances, bestFeature, "", true, threshold)
		rightInstances := make([]t.Instance, 0, len(instances)/2)

		for _, instance := range instances {
			if val, ok := instance[bestFeature]; ok && val != nil {
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
				if floatVal > threshold {
					rightInstances = append(rightInstances, instance)
				}
			}
		}

		// Create the child nodes
		leftNode := C45(leftInstances, features, targetFeature, featureTypes, excludedFeatures, minInstancesPerLeaf, maxDepth-1, cache)
		rightNode := C45(rightInstances, features, targetFeature, featureTypes, excludedFeatures, minInstancesPerLeaf, maxDepth-1, cache)

		// Add the children to the node
		node.Children = []*t.Node{leftNode, rightNode}
		node.Value = threshold
	} else {
		// Get all unique values for the feature
		featureValues := make(map[string]bool)
		for _, instance := range instances {
			val := fmt.Sprintf("%v", instance[bestFeature])
			featureValues[val] = true
		}

		// Create a child node for each value
		children := make([]*t.Node, 0, len(featureValues))
		for value := range featureValues {
			subsetInstances := utils.FilterInstances(instances, bestFeature, value, false, 0)
			if len(subsetInstances) > 0 {
				childNode := C45(subsetInstances, features, targetFeature, featureTypes, excludedFeatures, minInstancesPerLeaf, maxDepth-1, cache)
				childNode.Value = value
				children = append(children, childNode)
			}
		}

		// Handle missing values by adding a majority class child
		if len(children) < len(featureValues) {
			majorityClass := counter.GetMajorityClass()
			missingValueNode := &t.Node{
				IsLeaf: true,
				Class:  majorityClass,
				Value:  "unknown",
			}
			children = append(children, missingValueNode)
		}

		node.Children = children
	}

	return node
}
