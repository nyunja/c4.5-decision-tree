package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var command string
var inputFile string
var targetColumn string
var outputModel string
var modelFile string
// var dataFile string

var RootCmd = &cobra.Command{
	Use:     "dt",
	Version: "1.0.0",
	Short:   "C4.5 Decision Tree CLI",
	Long:    "A command-line tool for training and using a C4.5 decision tree model.",
	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Println("Invalid command. Use 'dt -c train' or 'dt -c predict'.")

		switch command {
		case "train":
			fmt.Println("Training")
			if inputFile == "" || targetColumn == "" || outputModel == "" {
				fmt.Println("Error: Missing required parameters for training.")
				cmd.Usage()
				return
			}
			TrainModel(inputFile, targetColumn, outputModel)
			// TrainCmd.Run(cmd, args)
		case "predict":
			PredictCmd.Run(cmd, args)
		default:
			fmt.Println("Invalid command. Use 'dt -c train' or 'dt -c predict'.")
			cmd.Usage()
		}
	},
}

func init() {
	// Define -c flag for specifying commands
	RootCmd.PersistentFlags().StringVarP(&command, "command", "c", "", "Specify command (train, predict)")

	// Flags for train
	RootCmd.PersistentFlags().StringVarP(&inputFile, "input", "i", "", "Input data file (CSV format)")
	RootCmd.MarkPersistentFlagRequired("input")
	RootCmd.PersistentFlags().StringVarP(&targetColumn, "target", "t", "", "Target column for training")
	RootCmd.PersistentFlags().StringVarP(&outputModel, "output", "o", "", "Output model file")

	// Flags for predict
	RootCmd.PersistentFlags().StringVarP(&modelFile, "model", "m", "", "Trained model file")
}
