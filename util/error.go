package util

import "log"

// Error log represents a structured error message
type ErrorLog struct {
	Error         string
	PossibleCause string
	SuggestedFix  string
}

// Predefined errors
var errorMessages = map[string]ErrorLog{
	"missing_input_file": {
		Error:         "Missing input file",
		PossibleCause: "Input CSV file path is incorrect or missing.",
		SuggestedFix:  "Check if the file exists and provide the correct path.",
	},
	"target_column_not_found": {
		Error:         "Target column not found",
		PossibleCause: "The column specified with -t is not in the dataset.",
		SuggestedFix:  "Verify the column name in the CSV file.",
	},
	"model_file_not_found": {
		Error:         "Model file not found",
		PossibleCause: "The specified model file does not exist.",
		SuggestedFix:  "Train a model first or check the file path.",
	},
	"output_path_missing": {
		Error:         "Output path not specified",
		PossibleCause: "The -o argument is missing.",
		SuggestedFix:  "Provide an output file path.",
	},
}

func LogError(errKey string) {
	if errLog, exists := errorMessages[errKey]; exists {
		log.Fatalf("ERROR: %s\nPossible Cause: %s\nSuggested Fix: %s\n", errLog.Error, errLog.PossibleCause, errLog.SuggestedFix)
	} else {
		log.Fatalln("ERROR: Unknown error occured")
	}
}
