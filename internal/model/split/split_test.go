package split

import (
	"reflect"
	"testing"

	"github.com/nyunja/c4.5-decision-tree/internal/model/cache"
	test "github.com/nyunja/c4.5-decision-tree/internal/model/types"
)

func TestFindBestSplit(t *testing.T) {
	// Mock instance structure
	type Instance map[string]interface{}

	instances := []test.Instance{
		{"feature1": "A", "feature2": 5.0, "target": "Yes"},
		{"feature1": "B", "feature2": 10.0, "target": "No"},
		{"feature1": "A", "feature2": 15.0, "target": "Yes"},
		{"feature1": "B", "feature2": 20.0, "target": "No"},
	}

	// Define feature types
	featureTypes := map[string]string{
		"feature1": "categorical",
		"feature2": "numerical",
	}

	// Mock featureCache
	mockCache := &cache.FeatureCache{
		SortedValues: map[string][]float64{
			"feature2": {5.0, 10.0, 15.0, 20.0},
		},
		ValueCounts: map[string]map[string]int{
			"feature1": {"A": 2, "B": 2},
		},
	}

	tests := []struct {
		name             string
		instances        []test.Instance
		features         []string
		targetFeature    string
		featureTypes     map[string]string
		excludedFeatures map[string]bool
		cache            *cache.FeatureCache
		wantFeature      string
		wantValue        interface{}
		wantIsContinuous bool
		wantThreshold    float64
	}{
		{
			name:             "Best split is categorical",
			instances:        instances,
			features:         []string{"feature1", "feature2"},
			targetFeature:    "target",
			featureTypes:     featureTypes,
			excludedFeatures: map[string]bool{},
			cache:            mockCache,
			wantFeature:      "feature1",
			wantValue:        nil,
			wantIsContinuous: false,
			wantThreshold:    0,
		},
		{
			name:             "Best split is numerical",
			instances:        instances,
			features:         []string{"feature2"},
			targetFeature:    "target",
			featureTypes:     featureTypes,
			excludedFeatures: map[string]bool{},
			cache:            mockCache,
			wantFeature:      "feature2",
			wantValue:        nil,
			wantIsContinuous: true,
			wantThreshold:    7.5, // Expected midpoint split
		},
		{
			name:             "No instances",
			instances:        []test.Instance{},
			features:         []string{"feature1", "feature2"},
			targetFeature:    "target",
			featureTypes:     featureTypes,
			excludedFeatures: map[string]bool{},
			cache:            mockCache,
			wantFeature:      "",
			wantValue:        nil,
			wantIsContinuous: false,
			wantThreshold:    0,
		},
		{
			name:             "All features excluded",
			instances:        instances,
			features:         []string{"feature1", "feature2"},
			targetFeature:    "target",
			featureTypes:     featureTypes,
			excludedFeatures: map[string]bool{"feature1": true, "feature2": true},
			cache:            mockCache,
			wantFeature:      "",
			wantValue:        nil,
			wantIsContinuous: false,
			wantThreshold:    0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotFeature, gotValue, gotIsContinuous, gotThreshold := FindBestSplit(
				tt.instances, tt.features, tt.targetFeature, tt.featureTypes, tt.excludedFeatures, tt.cache,
			)

			if gotFeature != tt.wantFeature {
				t.Errorf("FindBestSplit() gotFeature = %v, want %v", gotFeature, tt.wantFeature)
			}
			if !reflect.DeepEqual(gotValue, tt.wantValue) {
				t.Errorf("FindBestSplit() gotValue = %v, want %v", gotValue, tt.wantValue)
			}
			if gotIsContinuous != tt.wantIsContinuous {
				t.Errorf("FindBestSplit() gotIsContinuous = %v, want %v", gotIsContinuous, tt.wantIsContinuous)
			}
			if gotThreshold != tt.wantThreshold {
				t.Errorf("FindBestSplit() gotThreshold = %v, want %v", gotThreshold, tt.wantThreshold)
			}
		})
	}
}
