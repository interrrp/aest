package main

import (
	"fmt"
	"os"
)

type args struct {
	mode       string
	input      []byte
	outputFile string
	key        []byte
}

func parseArgs() *args {
	if len(os.Args) < 5 {
		printUsageAndFail()
	}

	mode := os.Args[1]
	if mode != "encrypt" && mode != "decrypt" {
		printUsageAndFail()
	}

	input, err := os.ReadFile(os.Args[2])
	if err != nil {
		fmt.Println("error reading input file:", err)
		os.Exit(1)
	}

	outputFile := os.Args[3]
	key := []byte(os.Args[4])

	return &args{mode, input, outputFile, key}
}

func printUsageAndFail() {
	fmt.Println("Usage: aest <encrypt|decrypt> <input file> <output file> <key>")
	os.Exit(1)
}
