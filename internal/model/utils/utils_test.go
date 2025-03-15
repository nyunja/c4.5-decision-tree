package utils

import (
	"testing"
	"time"

	test "github.com/nyunja/c4.5-decision-tree/internal/model/types"
	"github.com/stretchr/testify/assert"
)

func TestDetectIDColumns(t *testing.T) {
	stats := &test.DatasetStats{
		RowCount: 100,
		ColumnStats: map[string]*test.ColumnStats{
			"ID":          {IsNumeric: true, UniqueValues: map[string]int{"1": 1, "2": 1, "3": 1}, Count: 100},
			"index":       {IsNumeric: true, UniqueValues: map[string]int{"1": 1, "2": 1, "3": 1}, Count: 100},
			"key":         {IsNumeric: true, UniqueValues: map[string]int{"101": 1, "102": 1}, Count: 100},
			"name":        {IsNumeric: false, UniqueValues: map[string]int{"Alice": 1, "Bob": 1}, Count: 100},
			"random_num":  {IsNumeric: true, UniqueValues: map[string]int{"50": 1, "60": 1}, Count: 100},
			"transaction": {IsNumeric: true, UniqueValues: map[string]int{"1000": 1, "2000": 1}, Count: 100},
		},
	}

	headers := []string{"ID", "index", "key", "name", "random_num", "transaction"}

	idColumns := DetectIDColumns(stats, headers)

	expectedIDs := []string{"ID", "index", "key"}
	assert.ElementsMatch(t, expectedIDs, idColumns, "Detected ID columns do not match expected")
}

func TestContains(t *testing.T) {
	type args struct {
		slice []string
		item  string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Contains item in slice",
			args: args{slice: []string{"Alice", "Bob", "Charlie"}, item: "Alice"},
			want: true,
		},
		{
			name: "Does not contain item in slice",
			args: args{slice: []string{"Alice", "Bob", "Charlie"}, item: "Dave"},
			want: false,
		},
		{
			name: "Empty slice",
			args: args{slice: []string{}, item: "Alice"},
			want: false,
		},
		{
			name: "Nil slice",
			args: args{slice: nil, item: "Alice"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Contains(tt.args.slice, tt.args.item); got != tt.want {
				t.Errorf("Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMin(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Both positive, a < b", args{3, 5}, 3},
		{"Both positive, a > b", args{10, 7}, 7},
		{"Both negative, a < b", args{-5, -2}, -5},
		{"Both negative, a > b", args{-1, -8}, -8},
		{"Mixed positive and negative, a < b", args{-3, 4}, -3},
		{"Mixed positive and negative, a > b", args{6, -2}, -2},
		{"Both zero", args{0, 0}, 0},
		{"Zero and positive", args{0, 9}, 0},
		{"Zero and negative", args{0, -5}, -5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Min(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Min() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterInstances(t *testing.T) {
	instances := []test.Instance{
		{"feature1": "A", "feature2": 10, "feature3": time.Unix(1700000000, 0)},
		{"feature1": "B", "feature2": 20, "feature3": time.Unix(1800000000, 0)},
		{"feature1": "A", "feature2": 5, "feature3": time.Unix(1600000000, 0)},
		{"feature1": "C", "feature2": 30, "feature3": time.Unix(1900000000, 0)},
	}

	t.Run("Filter by categorical feature", func(t *testing.T) {
		filtered := FilterInstances(instances, "feature1", "A", false, 0)
		expected := []test.Instance{
			{"feature1": "A", "feature2": 10, "feature3": time.Unix(1700000000, 0)},
			{"feature1": "A", "feature2": 5, "feature3": time.Unix(1600000000, 0)},
		}
		assert.ElementsMatch(t, expected, filtered)
	})

	t.Run("Filter by numerical feature", func(t *testing.T) {
		filtered :=FilterInstances(instances, "feature2", "", true, 15)
		expected := []test.Instance{
			{"feature1": "A", "feature2": 10, "feature3": time.Unix(1700000000, 0)},
			{"feature1": "A", "feature2": 5, "feature3": time.Unix(1600000000, 0)},
		}
		assert.ElementsMatch(t, expected, filtered)
	})

	t.Run("No matching categorical value", func(t *testing.T) {
		filtered := FilterInstances(instances, "feature1", "D", false, 0)
		assert.Empty(t, filtered)
	})

	t.Run("No matching numerical value", func(t *testing.T) {
		filtered := FilterInstances(instances, "feature2", "", true, 1)
		assert.Empty(t, filtered)
	})
}
