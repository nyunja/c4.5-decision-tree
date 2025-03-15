package parser

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"math"
	"math/rand/v2"
	"os"
	"strconv"
	"time"

	t "github.com/nyunja/c4.5-decision-tree/internal/model/types"
	"github.com/nyunja/c4.5-decision-tree/internal/model/utils"
)

// StreamingCSVParser efficiently parses a CSV file in chunks
func StreamingCSVParser(file string, hasHeader bool, chunkSize int, targetColumn string) ([]t.Instance, []string, map[string]string, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("error opening file: %v", err)
	}
	defer f.Close()

	// Use buffered reader for better performance
	reader := bufio.NewReaderSize(f, 1<<20) // 1MB buffer
	csvReader := csv.NewReader(reader)
	csvReader.ReuseRecord = true // Reuse record buffer for better performance

	var headers []string
	if hasHeader {
		record, err := csvReader.Read()
		if err != nil {
			return nil, nil, nil, fmt.Errorf("error reading CSV header: %v", err)
		}
		headers = make([]string, len(record))
		copy(headers, record)
	}

	// First pass: collect statistics about the data
	stats := &t.DatasetStats{
		ColumnStats: make(map[string]*t.ColumnStats),
	}

	// Initialize column stats
	for _, header := range headers {
		stats.ColumnStats[header] = &t.ColumnStats{
			Min:          math.MaxFloat64,
			Max:          -math.MaxFloat64,
			UniqueValues: make(map[string]int),
			IsNumeric:    true,
			IsDate:       true,
			IsTimestamp:  true,
		}
	}

	// Sample the data to determine column types
	sampleSize := 10000
	sampleCount := 0

	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, nil, nil, fmt.Errorf("error reading CSV record: %v", err)
		}

		if len(headers) == 0 {
			headers = make([]string, len(record))
			for i := range record {
				headers[i] = fmt.Sprintf("col%d", i)
			}
			// Initialize column stats after headers are known
			for _, header := range headers {
				stats.ColumnStats[header] = &t.ColumnStats{
					Min:          math.MaxFloat64,
					Max:          -math.MaxFloat64,
					UniqueValues: make(map[string]int),
					IsNumeric:    true,
					IsDate:       true,
					IsTimestamp:  true,
				}
			}
		}

		stats.RowCount++

		// Only sample a subset of rows for type detection
		if sampleCount < sampleSize {
			for i, value := range record {
				header := headers[i]
				colStats := stats.ColumnStats[header]

				// Skip empty values
				if value == "" {
					continue
				}

				// Update unique values (limit to 1000 unique values for memory efficiency)
				if len(colStats.UniqueValues) < 1000 {
					colStats.UniqueValues[value]++
				}

				// Check if numeric
				if colStats.IsNumeric {
					floatVal, err := strconv.ParseFloat(value, 64)
					if err != nil {
						colStats.IsNumeric = false
					} else {
						colStats.Sum += floatVal
						colStats.Count++
						if floatVal < colStats.Min {
							colStats.Min = floatVal
						}
						if floatVal > colStats.Max {
							colStats.Max = floatVal
						}
					}
				}

				// Check if date (only if not numeric)
				if !colStats.IsNumeric && colStats.IsDate {
					_, err := time.Parse("2006-01-02", value)
					if err != nil {
						_, err = time.Parse("01/02/2006", value)
						if err != nil {
							_, err = time.Parse("2006/02/01", value)
							if err != nil {
								colStats.IsDate = false
							}
						}
					}
				}

				// Check if timestamp (only if not numeric and not date)
				if !colStats.IsNumeric && !colStats.IsDate && colStats.IsTimestamp {
					_, err := time.Parse(time.RFC3339, value)
					if err != nil {
						_, err = time.Parse("2006-01-02 15:04:05", value)
						if err != nil {
							colStats.IsTimestamp = false
						}
					}
				}
			}
			sampleCount++
		}
	}

	// Determine column types based on statistics
	featureTypes := make(map[string]string)
	for header, colStats := range stats.ColumnStats {
		if colStats.IsNumeric {
			featureTypes[header] = "numerical"
		} else if colStats.IsDate {
			featureTypes[header] = "date"
		} else if colStats.IsTimestamp {
			featureTypes[header] = "timestamp"
		} else {
			featureTypes[header] = "categorical"
			colStats.IsCategorical = true
		}
	}

	// Detect ID columns
	idColumns := utils.DetectIDColumns(stats, headers)
	fmt.Printf("Detected ID columns: %v\n", idColumns)

	// Reset file pointer for second pass
	f.Seek(0, 0)
	reader = bufio.NewReaderSize(f, 1<<20)
	csvReader = csv.NewReader(reader)
	csvReader.ReuseRecord = true

	// Skip header if present
	if hasHeader {
		_, err := csvReader.Read()
		if err != nil {
			return nil, nil, nil, fmt.Errorf("error skipping CSV header: %v", err)
		}
	}

	// Determine if we should use sampling for very large datasets
	useSampling := stats.RowCount > 100000
	samplingRate := 1.0
	if useSampling {
		// Adjust sampling rate based on dataset size
		samplingRate = float64(100000) / float64(stats.RowCount)
		fmt.Printf("Using sampling rate of %.2f%% for large dataset (%d rows)\n", samplingRate*100, stats.RowCount)
	}

	// Second pass: read and convert data
	instances := make([]t.Instance, 0, utils.Min(stats.RowCount, 100000))
	rowCount := 0

	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, nil, nil, fmt.Errorf("error reading CSV record: %v", err)
		}

		rowCount++

		// Apply sampling if needed
		if useSampling && rand.Float64() > samplingRate {
			continue
		}

		instance := make(t.Instance, len(headers))
		for i, value := range record {
			header := headers[i]

			// Skip ID columns
			if utils.Contains(idColumns, header) {
				continue
			}

			// Convert value based on feature type
			switch featureTypes[header] {
			case "numerical":
				if value == "" {
					instance[header] = nil
				} else {
					floatVal, err := strconv.ParseFloat(value, 64)
					if err != nil {
						instance[header] = value
					} else {
						instance[header] = floatVal
					}
				}
			case "date":
				if value == "" {
					instance[header] = nil
				} else {
					dateVal, err := time.Parse("2006-01-02", value)
					if err != nil {
						dateVal, err = time.Parse("01/02/2006", value)
						if err != nil {
							dateVal, err = time.Parse("02/01/2006", value)
							if err != nil {
								instance[header] = value
								continue
							}
						}
					}
					instance[header] = dateVal
				}
			case "timestamp":
				if value == "" {
					instance[header] = nil
				} else {
					timeVal, err := time.Parse(time.RFC3339, value)
					if err != nil {
						timeVal, err = time.Parse("2006-01-02 15:04:05", value)
						if err != nil {
							instance[header] = value
							continue
						}
					}
					instance[header] = timeVal
				}
			default:
				instance[header] = value
			}
		}
		// Only include instances that have the target column
		if _, ok := instance[targetColumn]; ok {
			instances = append(instances, instance)
		}

		// Break if we've collected enough instances (for very large files)
		if len(instances) >= chunkSize {
			break
		}
	}

	fmt.Printf("Loaded %d instances from %d total rows\n", len(instances), rowCount)
	return instances, headers, featureTypes, nil
}
