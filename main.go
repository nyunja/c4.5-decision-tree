package main

import (
	"log"

	"github.com/nyunja/c45-decision-tree/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		log.Fatalln(err)
	}
}
