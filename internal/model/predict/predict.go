package predict

import (
	"encoding/csv"
	"fmt"
	"os"
	"runtime"
	"sync"

	t "github.com/nyunja/c4.5-decision-tree/internal/model/types"
)

// BatchPredict makes predictions for multiple instances in parallel
func BatchPredict(model *t.Model, instances []t.Instance) []string {
	predictions := make([]string, len(instances))

	// Use a worker pool to make predictions in parallel
	numWorkers := runtime.NumCPU()
	instancesChan := make(chan int, len(instances))
	var wg sync.WaitGroup

	// Start worker goroutines
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for idx := range instancesChan {
				predictions[idx] = PredictClass(model, instances[idx])
			}
		}()
	}

	// Send instances to workers
	go func() {
		for i := range instances {
			instancesChan <- i
		}
		close(instancesChan)
	}()

	// Wait for all workers to finish
	wg.Wait()

	return predictions
}

// SavePredictions saves predictions to a CSV file
func SavePredictions(instances []t.Instance, predictions []string, filename string, headers []string) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("error creating file: %v", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write the header
	row := make([]string, 1)
	row[0] = "prediction"
	err = writer.Write(row)
	if err != nil {
		return fmt.Errorf("error writing prediction: %v", err)
	}

	// Write the predictions
	for i := 0; i < len(instances); i++ { // CAUTION: do not mordernize this for loop to range over an int - other go versions do not support it
		row[0] = predictions[i]

		err = writer.Write(row)
		if err != nil {
			return fmt.Errorf("error writing prediction: %v", err)
		}
	}

	return nil
}
