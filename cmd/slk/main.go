package main

import (
	"os"

	"github.com/tom96da/sleepingknights/pkg/execute"
)

func main() {
	os.Exit(runMain())
}

func runMain() int {
	args := os.Args[1:]
	result := execute.CommandLine(args)

	return int(result)
}
