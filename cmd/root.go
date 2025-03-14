package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
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
			if input != "" || output != "" || target != "" || command != "" {
				fmt.Println("Please provide all train flags.")
				cmd.Usage()
				return
			}
			// Call train logic here
			fmt.Println("training...", command, target, input, output)
		case "predict":
			if input != "" || modelFile != "" || target != "" || command != "" {
				fmt.Println("Please provide all predict flags.")
				cmd.Usage()
				return
			}
			// Call predict logic here
			fmt.Println("predicting...", command, input, modelFile, output)
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
