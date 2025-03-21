package node

import (
	"fmt"
	"github.com/nyunja/c4.5-decision-tree/internal/model/counter"
	t "github.com/nyunja/c4.5-decision-tree/internal/model/types"
)

// GetMajorityClass returns the majority class in a set of instances
func GetMajorityClass(instances []t.Instance, targetFeature string) string {
	counter := counter.NewClassCounter()

	for _, instance := range instances {
		targetVal := fmt.Sprintf("%v", instance[targetFeature])
		counter.Add(targetVal)
	}

	return counter.GetMajorityClass()
}

// GetMajorityClassFromNode gets the majority class from a node's children
func GetMajorityClassFromNode(node *t.Node) string {
	classCounts := make(map[string]int)

	for _, child := range node.Children {
		if child.IsLeaf {
			classCounts[child.Class]++
		}
	}

	majorityClass := ""
	maxCount := 0

	for class, count := range classCounts {
		if count > maxCount {
			maxCount = count
			majorityClass = class
		}
	}

	return majorityClass
}
