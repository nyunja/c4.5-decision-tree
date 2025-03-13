package main

import (
	"dt/cmd"
	"log"
)

// File where the trained model will be saved
const modelFile = "model.json"

// Main function
func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		log.Fatalf("Error: %v", err)
	}
}
