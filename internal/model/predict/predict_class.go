package predict

import (
	"fmt"
	"strconv"
	"time"

	t "github.com/nyunja/c4.5-decision-tree/internal/model/types"
	ndp "github.com/nyunja/c4.5-decision-tree/internal/model/node"
)

// PredictClass predicts the class of an instance
func PredictClass(model *t.Model, instance t.Instance) string {
	node := model.Root
	for !node.IsLeaf {
		feature := node.Feature
		val, ok := instance[feature]
		if !ok || val == nil {
			// Handle missing value
			// Find the most common class among children
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

		if node.Continuous {
			var floatVal float64
			switch v := val.(type) {
			case float64:
				floatVal = v
			case int:
				floatVal = float64(v)
			case time.Time:
				floatVal = float64(v.Unix())
			default:
				strVal := fmt.Sprintf("%v", val)
				parsedVal, err := strconv.ParseFloat(strVal, 64)
				if err != nil {
					// Can't convert to float, use majority class
					return ndp.GetMajorityClassFromNode(node)
				}
				floatVal = parsedVal
			}

			if floatVal <= node.Threshold {
				node = node.Children[0] // Left child
			} else {
				node = node.Children[1] // Right child
			}
		} else {
			strVal := fmt.Sprintf("%v", val)
			found := false
			for _, child := range node.Children {
				childVal := fmt.Sprintf("%v", child.Value)
				if childVal == strVal {
					node = child
					found = true
					break
				}
			}
			if !found {
				// If value not found in any child, use the "unknown" branch or majority class
				for _, child := range node.Children {
					if fmt.Sprintf("%v", child.Value) == "unknown" {
						node = child
						found = true
						break
					}
				}
				if !found {
					return ndp.GetMajorityClassFromNode(node)
				}
			}
		}
	}
	return node.Class
}
