package dataprocessor

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/nyunja/c45-decision-tree/internal"
)

func WritePredictionCSVFile(filename string, data *Dataset, predictions []internal.Prediction) error {
	outFile, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer outFile.Close()

	writer := csv.NewWriter(outFile)
	defer writer.Flush()

	// Write the new headers
	headers := append(data.Header, []string{"predictions", "confidence"}...)
	if err := writer.Write(headers); err != nil {
		return fmt.Errorf("failed to write headers: %v", err)
	}
	for i := 0; i < len(data.Data); i++ {
		row := append(convertToStringSlice(data.Data[i]), predictions[i].Class, fmt.Sprintf("%.2f", predictions[i].Confidence*100))
        if err := writer.Write(row); err != nil {
            return fmt.Errorf("failed to write row: %v", err)
        }
	}
	return nil
}

func convertToStringSlice(slice []any) []string {
	result := make([]string, len(slice))
    for i, v := range slice {
		if str, ok :=  v.(string) ; ok{
			result[i] = str
		} else {
			result[i] = fmt.Sprintf("%v", v)
		}
    }
    return result
}