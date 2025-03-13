package cmd

import (
	"dt/internal/c45"
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var PredictCmd = &cobra.Command{
	Use:   "predict",
	Short: "Predict labels using a trained C4.5 model",
	Run: func(cmd *cobra.Command, args []string) {
		modelFile, _ := cmd.Flags().GetString("model")
		inputData, _ := cmd.Flags().GetString("input")

		tree, err := c45.LoadModel(modelFile)
		if err != nil {
			log.Fatalf("Failed to load model: %v", err)
		}
		data, err := loadCSV(inputData)
		if err != nil {
			log.Fatalf("Failed to read input file: %v", err)
		}
		for _, row := range data {
			prediction := tree.Predict(row)
			fmt.Println("Prediction:", prediction)
		}
	},
}

func init() {
	PredictCmd.Flags().StringP("input", "i", "", "Path to input data file (CSV)")
	PredictCmd.Flags().StringP("target", "t", "", "Target column for classification")
	PredictCmd.Flags().StringP("model", "m", "models/model.json", "Path to trained model")
	PredictCmd.MarkFlagRequired("input")
	PredictCmd.MarkFlagRequired("target")
	PredictCmd.MarkFlagRequired("model")
	RootCmd.AddCommand(PredictCmd)
}
