package utils

import (
	"fmt"
	"strings"
	"time"

	t "github.com/nyunja/c4.5-decision-tree/internal/model/types"
)

// detectIDColumns identifies columns that are likely to be IDs or indexes
func DetectIDColumns(stats *t.DatasetStats, headers []string) []string {
	idColumns := []string{}

	for _, header := range headers {
		// Skip any empty columns
		if header == "" {
			continue
		}

		// Check if the column name suggests it's an ID
		lowerHeader := strings.ToLower(header)
		if lowerHeader == "id" || lowerHeader == "index" || lowerHeader == "key" ||
			strings.HasSuffix(lowerHeader, "id") || strings.HasSuffix(lowerHeader, "key") ||
			strings.HasPrefix(lowerHeader, "id") || lowerHeader == "day" {
			idColumns = append(idColumns, header)
			continue
		}

		colStats := stats.ColumnStats[header]

		// Check if numerical and all values are unique (high cardinality)
		if colStats.IsNumeric && len(colStats.UniqueValues) > 0 {
			// If we have close to as many unique values as rows, it's likely an ID
			uniqueRatio := float64(len(colStats.UniqueValues)) / float64(Min(1000, colStats.Count))
			if uniqueRatio > 0.9 {
				idColumns = append(idColumns, header)
			}
		}
	}

	return idColumns
}

// contains checks if a string is in a slice
func Contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

// min returns the minimum of two integers
func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// filterInstances filters instances based on a feature and value
func FilterInstances(instances []t.Instance, feature string, value string, isContinuous bool, threshold float64) []t.Instance {
	// Pre-allocate with a reasonable capacity
	filteredInstances := make([]t.Instance, 0, len(instances)/2)

	for _, instance := range instances {
		val, ok := instance[feature]
		if !ok || val == nil {
			continue
		}

		if isContinuous {
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
			if floatVal <= threshold {
				filteredInstances = append(filteredInstances, instance)
			}
		} else {
			if fmt.Sprintf("%v", val) == value {
				filteredInstances = append(filteredInstances, instance)
			}
		}
	}

	return filteredInstances
}

// convertRecordToInstance converts a CSV record to an Instance object
func ConvertRecordToInstance(record []string, headers []string, featureTypes map[string]string) t.Instance {
	instance := make(t.Instance, len(headers))

	for i, value := range record {
		header := headers[i]

		// Convert value based on feature type
		switch featureTypes[header] {
		case "numerical":
			convert, err := ConvertStringToNumerical(header)
			if err != nil {
				instance[header] = convert
			} else {
				instance[header] = value
			}
		case "date":
			convert, err := ConvertStringToDate(value)
			if err != nil {
				instance[header] = convert
			} else {
				instance[header] = value
			}
		case "timestamp":
			convert, err := ConvertStringToTimestamp(value)
			if err != nil {
				instance[header] = convert
			} else {
				instance[header] = value
			}
		default:
			instance[header] = value
		}
	}

	return instance
}
