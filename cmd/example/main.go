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

	fmt.Println("CLI app structure ready")
}
