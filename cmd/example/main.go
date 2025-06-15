package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	lab2 "github.com/LeVasTiaN/KPI_Lab2"
)

var (
	inputExpression = flag.String("e", "", "Expression to compute")
	inputFile       = flag.String("f", "", "File with expression to compute")
	outputFile      = flag.String("o", "", "Output file for result")
)

func main() {
	flag.Parse()

	if (*inputExpression == "" && *inputFile == "") || (*inputExpression != "" && *inputFile != "") {
		fmt.Fprintf(os.Stderr, "Error: Must provide exactly one of -e (expression) or -f (file)\n")
		os.Exit(1)
	}

	var inputReader *strings.Reader
	var inputFileHandle *os.File
	var err error

	if *inputExpression != "" {
		inputReader = strings.NewReader(*inputExpression)
	} else {
		inputFileHandle, err = os.Open(*inputFile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error opening input file: %v\n", err)
			os.Exit(1)
		}
		defer inputFileHandle.Close()
	}

	var outputWriter *os.File
	if *outputFile != "" {
		outputWriter, err = os.Create(*outputFile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error creating output file: %v\n", err)
			os.Exit(1)
		}
		defer outputWriter.Close()
	} else {
		outputWriter = os.Stdout
	}

	var handler *lab2.ComputeHandler
	if inputReader != nil {
		handler = &lab2.ComputeHandler{
			Input:  inputReader,
			Output: outputWriter,
		}
	} else {
		handler = &lab2.ComputeHandler{
			Input:  inputFileHandle,
			Output: outputWriter,
		}
	}

	err = handler.Compute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error processing expression: %v\n", err)
		os.Exit(1)
	}
}
