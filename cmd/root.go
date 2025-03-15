package cmd

import (
	"fmt"
	"log"
	"os"
	"runtime"

	"github.com/spf13/cobra"

	m "github.com/nyunja/c45-decision-tree/internal/model/model"
	p "github.com/nyunja/c45-decision-tree/internal/model/parser"
	"github.com/nyunja/c45-decision-tree/internal/model/predict"
)

var (
	command   string
	target    string
	input     string
	output    string
	modelFile string
)

// Define the subcommands for train and predict commands
var RootCmd = &cobra.Command{
	Use:   "dt",
	Short: "C4.5 Decision Tree CLI",
	Run: func(cmd *cobra.Command, args []string) {
		switch command {
		case "train":
			if input == "" || output == "" || target == "" || command == "" {
				fmt.Println("Please provide all train flags.")
				cmd.Usage()
				return
			}
			// call train logic here
			fmt.Println("training...", command, target, input, output)

			// check if input file exists
			if _, err := os.Stat(input); os.IsNotExist(err) {
				fmt.Println("Error: Input file not found")
				os.Exit(1)
			}

			// parse the CSV file with streaming
			instances, headers, featureTypes, err := p.StreamingCSVParser(input, true, 10000, target)
			if err != nil {
				log.Fatalf("Error parsing CSV: %v", err)
			}
			fmt.Printf("Parsed %d instances with %d features\n", len(instances), len(headers))

			// Check if target column exists
			if _, ok := featureTypes[target]; !ok {
				fmt.Println("Error: Target column not found")
				os.Exit(1)
			}

			// Parse user-specified excluded columns
			excludeColumns := []string{}

			fmt.Printf("Columns excluded from training: %v\n", excludeColumns)

			// Train the model
			fmt.Println("Training model...")
			model, err := m.Train(instances, headers, target, featureTypes, excludeColumns, 20)
			if err != nil {
				log.Fatalf("Error training model: %v", err)
			}
			fmt.Println("Model trained successfully")

			// Save the model
			fmt.Println("Saving model...")
			err = m.SaveModel(model, output)
			if err != nil {
				log.Fatalf("Error saving model: %v", err)
			}

		case "predict":
			if output == "" || input == "" || modelFile == "" || command == "" {
				fmt.Println("Please provide all predict flags.")
				cmd.Usage()
				return
			}

			// check if input file exists
			if _, err := os.Stat(input); os.IsNotExist(err) {
				fmt.Println("Error: Input file not found")
				os.Exit(1)
			}

			// Load the model
			fmt.Println("Loading model...")
			model, err := m.LoadModel(modelFile)
			if err != nil {
				log.Fatalf("Error loading model: %v", err)
			}
			fmt.Println("Model loaded successfully")

			// parse the CSV file with streaming
			instances, headers, _, err := p.PredictionCSVParser(input, true, 10000, model.TargetName)
			if err != nil {
				log.Fatalf("Error parsing CSV: %v", err)
			}
			fmt.Printf("Parsed %d instances with %d features\n", len(instances), len(headers))

			// Make predictions
			fmt.Println("Making predictions...")
			predictions := predict.BatchPredict(model, instances)
			fmt.Println("Predictions made successfully")

			// Save predictions
			fmt.Println("Saving predictions...")
			err = predict.SavePredictions(instances, predictions, output, headers)
			if err != nil {
				log.Fatalf("Error saving predictions: %v", err)
			}

			fmt.Printf("Predictions successfully made and saved to %s\n", output)

		default:
			fmt.Println("Invalid command. Use -c train")
			cmd.Usage()
		}
	},
}

// Run the command
func init() {
	RootCmd.PersistentFlags().StringVarP(&command, "command", "c", "", "Specify command (train)")
	RootCmd.MarkPersistentFlagRequired("command")
	RootCmd.PersistentFlags().StringVarP(&target, "target", "t", "", "Specify target column")
	RootCmd.PersistentFlags().StringVarP(&input, "input", "i", "", "Input data file (CSV format)")
	RootCmd.PersistentFlags().StringVarP(&output, "output", "o", "", "Output model file")
	RootCmd.PersistentFlags().StringVarP(&modelFile, "model", "m", "", "Training model file")
}
