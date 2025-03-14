package dataprocessor

import (
	"reflect"
	"testing"
	"time"
)

/*
test parseData() function
TestParseData tests the ParseData function to ensure it correctly parses
various types of data according to the provided metadata. The test cases
include:
- Valid data with mixed types (numerical, date, timestamp, and category)
- Invalid numerical value
- Invalid date value
- Invalid timestamp value
- Empty value
*/
func TestParseData(t *testing.T) {
	tests := []struct {
		name     string
		data     [][]string
		metadata []ColumnData
		want     [][]any
		wantErr  bool
	}{
		{
			name: "Valid data with mixed types",
			data: [][]string{
				{"1.23", "2023-01-01", "2023-01-01T12:00:00Z", "text"},
				{"4.56", "2023-02-02", "2023-02-02T12:00:00Z", "more text"},
			},
			metadata: []ColumnData{
				{Type: NumType},
				{Type: DateType},
				{Type: TimestampType},
				{Type: CategoryType},
			},
			want: [][]any{
				{1.23, mustParseDate("2006-01-02", "2023-01-01"), mustParseTimestamp("2023-01-01T12:00:00Z"), "text"},
				{4.56, mustParseDate("2006-01-02", "2023-02-02"), mustParseTimestamp("2023-02-02T12:00:00Z"), "more text"},
			},
			wantErr: false,
		},
		{
			name: "Invalid numerical value",
			data: [][]string{
				{"invalid", "2023-01-01", "2023-01-01T12:00:00Z", "text"},
			},
			metadata: []ColumnData{
				{Type: NumType},
				{Type: DateType},
				{Type: TimestampType},
				{Type: CategoryType},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Invalid date value",
			data: [][]string{
				{"1.23", "invalid", "2023-01-01T12:00:00Z", "text"},
			},
			metadata: []ColumnData{
				{Type: NumType},
				{Type: DateType},
				{Type: TimestampType},
				{Type: CategoryType},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Invalid timestamp value",
			data: [][]string{
				{"1.23", "2023-01-01", "invalid", "text"},
			},
			metadata: []ColumnData{
				{Type: NumType},
				{Type: DateType},
				{Type: TimestampType},
				{Type: CategoryType},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Empty value",
			data: [][]string{
				{"", "2023-01-01", "2023-01-01T12:00:00Z", "text"},
			},
			metadata: []ColumnData{
				{Type: NumType},
				{Type: DateType},
				{Type: TimestampType},
				{Type: CategoryType},
			},
			want: [][]any{
				{nil, mustParseDate("2006-01-02", "2023-01-01"), mustParseTimestamp("2023-01-01T12:00:00Z"), "text"},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got [][]any
			var err error
			for _, row := range tt.data {
				parsedRow := make([]any, len(tt.metadata))
				err = ParseData(row, tt.metadata, parsedRow)
				if err != nil {
					got = nil
					break
				}
				got = append(got, parsedRow)
			}
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseData() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func mustParseDate(layout, value string) time.Time {
	date, err := time.Parse(layout, value)
	if err != nil {
		panic(err)
	}
	return date
}

func mustParseTimestamp(value string) time.Time {
	timestamp, err := time.Parse(time.RFC3339, value)
	if err != nil {
		panic(err)
	}
	return timestamp
}
