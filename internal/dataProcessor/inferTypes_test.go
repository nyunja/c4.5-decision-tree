package dataprocessor

import (
	"reflect"
	"testing"
	"time"
)

func TestInferColumnTypes(t *testing.T) {
	tests := []struct {
		name         string
		data         [][]string
		columnHeader []string
		expected     []ColumnData
		expectErr    bool
	}{
		{
			name:         "Empty dataset",
			data:         [][]string{},
			columnHeader: []string{"A", "B"},
			expected:     nil,
			expectErr:    true,
		},
		{
			name: "Numeric column",
			data: [][]string{
				{"1.0", "2.0"},
				{"3.0", "4.0"},
			},
			columnHeader: []string{"A", "B"},
			expected: []ColumnData{
				{Name: "A", Type: NumType},
				{Name: "B", Type: NumType},
			},
			expectErr: false,
		},
		{
			name: "Date column",
			data: [][]string{
				{"2023-01-01", "2023-01-02"},
				{"2023-01-03", "2023-01-04"},
			},
			columnHeader: []string{"A", "B"},
			expected: []ColumnData{
				{Name: "A", Type: DateType},
				{Name: "B", Type: DateType},
			},
			expectErr: false,
		},
		{
			name: "Timestamp column",
			data: [][]string{
				{"2023-01-01T00:00:00Z", "2023-01-02T00:00:00Z"},
				{"2023-01-03T00:00:00Z", "2023-01-04T00:00:00Z"},
			},
			columnHeader: []string{"A", "B"},
			expected: []ColumnData{
				{Name: "A", Type: TimestampType},
				{Name: "B", Type: TimestampType},
			},
			expectErr: false,
		},
		{
			name: "Mixed column",
			data: [][]string{
				{"1.0", "2023-01-01"},
				{"A", "2023-01-02"},
			},
			columnHeader: []string{"A", "B"},
			expected: []ColumnData{
				{Name: "A", Type: CategoryType},
				{Name: "B", Type: DateType},
			},
			expectErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, result, err := InferColumnTypes(tt.data, tt.columnHeader)
			if (err != nil) != tt.expectErr {
				t.Errorf("expected error: %v, got: %v", tt.expectErr, err)
			}
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("%s: expected: %v, got: %v", tt.name, tt.expected, result)
			}
		})
	}
}

func TestInferColumnTypes_ParseData(t *testing.T) {
	tests := []struct {
		name         string
		data         [][]string
		columnHeader []string
		expected     [][]any
		expectErr    bool
	}{
		{
			name: "Parse numeric data",
			data: [][]string{
				{"1.0", "2.0"},
				{"3.0", "4.0"},
			},
			columnHeader: []string{"A", "B"},
			expected: [][]any{
				{1.0, 2.0},
				{3.0, 4.0},
			},
			expectErr: false,
		},
		{
			name: "Parse date data",
			data: [][]string{
				{"2023-01-01", "2023-01-02"},
				{"2023-01-03", "2023-01-04"},
			},
			columnHeader: []string{"A", "B"},
			expected: [][]any{
				{parseDate("2023-01-01"), parseDate("2023-01-02")},
				{parseDate("2023-01-03"), parseDate("2023-01-04")},
			},
			expectErr: false,
		},
		{
			name: "Parse timestamp data",
			data: [][]string{
				{"2023-01-01T00:00:00Z", "2023-01-02T00:00:00Z"},
				{"2023-01-03T00:00:00Z", "2023-01-04T00:00:00Z"},
			},
			columnHeader: []string{"A", "B"},
			expected: [][]any{
				{parseTimestamp("2023-01-01T00:00:00Z"), parseTimestamp("2023-01-02T00:00:00Z")},
				{parseTimestamp("2023-01-03T00:00:00Z"), parseTimestamp("2023-01-04T00:00:00Z")},
			},
			expectErr: false,
		},
		{
			name: "Parse mixed data",
			data: [][]string{
				{"1.0", "2023-01-01"},
				{"A", "2023-01-02"},
			},
			columnHeader: []string{"A", "B"},
			expected: [][]any{
				{"1.0", parseDate("2023-01-01")},
				{"A", parseDate("2023-01-02")},
			},
			expectErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, _, err := InferColumnTypes(tt.data, tt.columnHeader)
			if (err != nil) != tt.expectErr {
				t.Errorf("expected error: %v, got: %v", tt.expectErr, err)
			}
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("%s: expected: %v, got: %v", tt.name, tt.expected, result)
			}
		})
	}
}

func parseDate(value string) time.Time {
	date, _ := time.Parse("2006-01-02", value)
	return date
}

func parseTimestamp(value string) time.Time {
	timestamp, _ := time.Parse(time.RFC3339, value)
	return timestamp
}
