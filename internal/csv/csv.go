package csv

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"math"
	"math/rand/v2"
	"os"
	"strconv"

	t "github.com/nyunja/c4.5-decision-tree/internal/model/types"
	"github.com/nyunja/c4.5-decision-tree/internal/model/utils"
)

// openCSVFile opens a CSV file and creates a buffered reader
func OpenCSVFile(file string) (*os.File, *csv.Reader, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, nil, fmt.Errorf("error opening file: %v", err)
	}

	// Use buffered reader for better performance
	reader := bufio.NewReaderSize(f, 1<<20) // 1MB buffer
	csvReader := csv.NewReader(reader)
	csvReader.ReuseRecord = true // Reuse record buffer for better performance

	return f, csvReader, nil
}

// readCSVHeaders reads headers from CSV file
func ReadCSVHeaders(csvReader *csv.Reader, hasHeader bool) ([]string, error) {
	var headers []string

	if hasHeader {
		record, err := csvReader.Read()
		if err != nil {
			return nil, fmt.Errorf("error reading CSV header: %v", err)
		}
		headers = make([]string, len(record))
		copy(headers, record)
	}

	return headers, nil
}

// initializeColumnStats creates initial statistics for each column
func InitializeColumnStats(headers []string) map[string]*t.ColumnStats {
	columnStats := make(map[string]*t.ColumnStats)

	for _, header := range headers {
		columnStats[header] = &t.ColumnStats{
			Min:          math.MaxFloat64,
			Max:          -math.MaxFloat64,
			UniqueValues: make(map[string]int),
			IsNumeric:    true,
			IsDate:       true,
			IsTimestamp:  true,
		}
	}

	return columnStats
}

// collectDatasetStatistics performs the first pass through the data to gather statistics
func CollectDatasetStatistics(f *os.File, headers []string, hasHeader bool) (*t.DatasetStats, error) {
	// Reset file to beginning
	f.Seek(0, 0)
	reader := bufio.NewReaderSize(f, 1<<20)
	csvReader := csv.NewReader(reader)
	csvReader.ReuseRecord = true

	// Skip header if present
	if hasHeader {
		_, err := csvReader.Read()
		if err != nil {
			return nil, fmt.Errorf("error skipping CSV header: %v", err)
		}
	}

	stats := &t.DatasetStats{
		ColumnStats: make(map[string]*t.ColumnStats),
	}

	// Initialize column stats
	stats.ColumnStats = InitializeColumnStats(headers)

	// Sample the data to determine column types
	sampleSize := 10000
	sampleCount := 0

	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("error reading CSV record: %v", err)
		}

		if len(headers) == 0 {
			headers = make([]string, len(record))
			for i := range record {
				headers[i] = fmt.Sprintf("col%d", i)
			}
			// Initialize column stats after headers are known
			stats.ColumnStats = InitializeColumnStats(headers)
		}

		stats.RowCount++

		// Only sample a subset of rows for type detection
		if sampleCount < sampleSize {
			UpdateColumnStatistics(record, headers, stats.ColumnStats)
			sampleCount++
		}
	}

	return stats, nil
}

// updateColumnStatistics analyzes a single row and updates column statistics
func UpdateColumnStatistics(record []string, headers []string, columnStats map[string]*t.ColumnStats) {
	for i, value := range record {
		header := headers[i]
		colStats := columnStats[header]

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
			colStats.IsDate = utils.IsDateValue(value)
		}

		// Check if timestamp (only if not numeric and not date)
		if !colStats.IsNumeric && !colStats.IsDate && colStats.IsTimestamp {
			colStats.IsTimestamp = utils.IsTimestampValue(value)
		}
	}
}

// determineColumnTypes analyzes statistics to determine the type of each column
func DetermineColumnTypes(stats *t.DatasetStats) map[string]string {
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

	return featureTypes
}

// loadInstances performs the second pass through the data to load instances
func LoadInstances(file string, headers []string, featureTypes map[string]string,
	targetColumn string, totalRows int, chunkSize int, hasHeader bool,
) ([]t.Instance, error) {
	// Open file again for second pass
	f, err := os.Open(file)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %v", err)
	}
	defer f.Close()

	reader := bufio.NewReaderSize(f, 1<<20)
	csvReader := csv.NewReader(reader)
	csvReader.ReuseRecord = true

	// Skip header if present
	if hasHeader {
		_, err := csvReader.Read()
		if err != nil {
			return nil, fmt.Errorf("error skipping CSV header: %v", err)
		}
	}

	// Determine if we should use sampling for very large datasets
	useSampling := totalRows > 100000
	samplingRate := 1.0
	if useSampling {
		// Adjust sampling rate based on dataset size
		samplingRate = float64(100000) / float64(totalRows)
		fmt.Printf("Using sampling rate of %.2f%% for large dataset (%d rows)\n", samplingRate*100, totalRows)
	}

	// Read and convert data
	instances := make([]t.Instance, 0, utils.Min(totalRows, 100000))
	rowCount := 0

	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("error reading CSV record: %v", err)
		}

		rowCount++

		// Apply sampling if needed
		if useSampling && rand.Float64() > samplingRate {
			continue
		}

		instance := utils.ConvertRecordToInstance(record, headers, featureTypes)

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
	return instances, nil
}

// loadPredictionInstances performs the second pass through the data to load instances for prediction
func LoadPredictionInstances(file string, headers []string, featureTypes map[string]string,
	targetColumn string, totalRows int, chunkSize int, hasHeader bool,
) ([]t.Instance, error) {
	// Open file again for second pass
	f, err := os.Open(file)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %v", err)
	}
	defer f.Close()

	reader := bufio.NewReaderSize(f, 1<<20)
	csvReader := csv.NewReader(reader)
	csvReader.ReuseRecord = true

	// Skip header if present
	if hasHeader {
		_, err := csvReader.Read()
		if err != nil {
			return nil, fmt.Errorf("error skipping CSV header: %v", err)
		}
	}

	// Determine if we should use sampling for very large datasets
	useSampling := totalRows > 100000
	samplingRate := 1.0
	if useSampling {
		// Adjust sampling rate based on dataset size
		samplingRate = float64(100000) / float64(totalRows)
		fmt.Printf("Using sampling rate of %.2f%% for large dataset (%d rows)\n", samplingRate*100, totalRows)
	}

	// Check if target column exists in dataset
	targetExists := utils.Contains(headers, targetColumn)
	fmt.Printf("Target column '%s' exists in dataset: %v\n", targetColumn, targetExists)

	// Read and convert data
	instances := make([]t.Instance, 0, utils.Min(totalRows, 100000))
	rowCount := 0

	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("error reading CSV record: %v", err)
		}

		rowCount++

		// Apply sampling if needed
		if useSampling && rand.Float64() > samplingRate {
			continue
		}

		instance := utils.ConvertPredictionRecordToInstance(record, headers, featureTypes)

		// For prediction, we don't require the target column to be present
		instances = append(instances, instance)

		// Break if we've collected enough instances (for very large files)
		if len(instances) >= chunkSize {
			break
		}
	}

	fmt.Printf("Loaded %d instances from %d total rows for prediction\n", len(instances), rowCount)
	return instances, nil
}
