package main

import (
	"log"
	"runtime"

	"github.com/nyunja/c45-decision-tree/cmd"
)

func main() {
	numCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(numCPU)

	if err := cmd.RootCmd.Execute(); err != nil {
		log.Fatalln(err)
	}
}
