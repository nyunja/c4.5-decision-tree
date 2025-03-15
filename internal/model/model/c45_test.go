package model

import (
	"testing"

	"github.com/nyunja/c4.5-decision-tree/internal/model/cache"
	t "github.com/nyunja/c4.5-decision-tree/internal/model/types"
	"github.com/stretchr/testify/assert"
)

func TestC45_EmptyDataset(tc *testing.T) {
	cache := &cache.FeatureCache{}
	instances := []t.Instance{}
	features := []string{"age", "income"}
	featureTypes := map[string]string{"age": "numerical", "income": "numerical"}
	tree := C45(instances, features, "category", featureTypes, map[string]bool{}, 1, 3, cache)

	assert.NotNil(tc, tree)
	assert.True(tc, tree.IsLeaf)
}

func TestC45_MaxDepthReached(tc *testing.T) {
	cache := &cache.FeatureCache{}
	instances := []t.Instance{
		{"age": 25, "category": "A"},
		{"age": 30, "category": "B"},
	}
	features := []string{"age"}
	featureTypes := map[string]string{"age": "numerical"}
	tree := C45(instances, features, "category", featureTypes, map[string]bool{}, 1, 0, cache)

	assert.NotNil(tc, tree)
	assert.True(tc, tree.IsLeaf)
}

func TestC45_SingleClass(tc *testing.T) {
	cache := &cache.FeatureCache{}
	instances := []t.Instance{
		{"age": 25, "category": "A"},
		{"age": 30, "category": "A"},
	}
	features := []string{"age"}
	featureTypes := map[string]string{"age": "numerical"}
	tree := C45(instances, features, "category", featureTypes, map[string]bool{}, 1, 3, cache)

	assert.NotNil(tc, tree)
	assert.True(tc, tree.IsLeaf)
	assert.Equal(tc, "A", tree.Class)
}

func TestC45_CategoricalSplit(tc *testing.T) {
	cache := &cache.FeatureCache{}
	instances := []t.Instance{
		{"": "red", "category": "A"},
		{"": "blue", "category": "B"},
		{"": "green", "category": "A"},
	}
	features := []string{"color"}
	featureTypes := map[string]string{"": "categorical"}
	tree := C45(instances, features, "category", featureTypes, map[string]bool{}, 1, 3, cache)

	assert.NotNil(tc, tree)
	assert.Equal(tc, "", tree.Feature)
}
