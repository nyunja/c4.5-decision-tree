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

// Should create a new FeatureCache with empty TargetCounts map
func TestNewFeatureCache_EmptyTargetCountMap(t *testing.T) {
	cache := NewFeatureCache()
	if len(cache.TargetCounts) != 0 {
		t.Errorf("Expected empty TargetCounts, got %d items", len(cache.TargetCounts))
	}
}

// Should return a pointer to a FeatureCache struct
func TestNewFeatureCache_ReturnType(t *testing.T) {
	cache := NewFeatureCache()

	if cache == nil {
		t.Fatal("Expected a non-nil FeatureCache pointer, but got nil")
	}

	if cache.ValueCounts == nil {
		t.Error("Expected non-nil ValueCounts map")
	}

	if cache.SortedValues == nil {
		t.Error("Expected non-nil SortedValues map")
	}

	if cache.TargetCounts == nil {
		t.Error("Expected non-nil TargetCounts map")
	}
}

// Should initialize ValueCounts as a map of string to map of string to int
func TestNewFeatureCache_ValueCountsAsMap(t *testing.T) {
	cache := NewFeatureCache()

	if cache.ValueCounts == nil {
		t.Error("ValueCounts should be initialized")
	}

	if len(cache.ValueCounts) != 0 {
		t.Error("ValueCounts should be empty initially")
	}

	// Test that we can add to the ValueCounts map
	cache.ValueCounts["feature1"] = make(map[string]int)
	cache.ValueCounts["feature1"]["value1"] = 1

	if count := cache.ValueCounts["feature1"]["value1"]; count != 1 {
		t.Errorf("Expected count 1, got %d", count)
	}
}
