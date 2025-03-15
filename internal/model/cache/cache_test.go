package cache

import "testing"

// Should create a new FeatureCache with empty ValueCounts map
func TestNewFeatureCache(t *testing.T) {
	cache := NewFeatureCache()

	if cache == nil {
		t.Error("NewFeatureCache returned nil")
	}

	if len(cache.ValueCounts) != 0 {
		t.Errorf("Expected empty ValueCounts, got %d items", len(cache.ValueCounts))
	}

	if len(cache.SortedValues) != 0 {
		t.Errorf("Expected empty SortedValues, got %d items", len(cache.SortedValues))
	}

	if len(cache.TargetCounts) != 0 {
		t.Errorf("Expected empty TargetCounts, got %d items", len(cache.TargetCounts))
	}
}

// Should create a new FeatureCache with empty SortedValues map
func TestNewFeatureCache_EmptySortedValueMap(t *testing.T) {
	cache := NewFeatureCache()

	if cache == nil {
		t.Fatal("NewFeatureCache returned nil")
	}

	if len(cache.SortedValues) != 0 {
		t.Errorf("Expected empty SortedValues map, got %d elements", len(cache.SortedValues))
	}
}
