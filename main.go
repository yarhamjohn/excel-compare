package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/yarhamjohn/excel-compare/flatten"
)

func main() {
	analyzeCommand := flag.NewFlagSet("analyze", flag.ExitOnError)
	flattenCommand := flag.NewFlagSet("flatten", flag.ExitOnError)

	// Analyze subcommand flag pointers
	analyzeFilePath := analyzeCommand.String("filePath", "", "Excel file to analyse. (Required)")

	// Flatten subcommand flag pointers
	flattenFilePath := flattenCommand.String("filePath", "", "Excel file to analyse. (Required)")
	delimiter := flattenCommand.String("delimiter", ",", "The field delimiter to use {,|^}. (Required)")

	// Checks a subcommand has been provided
	// os.Args[0] is the main command, os.Args[1] the subcommand
	if len(os.Args) < 2 {
		fmt.Println("One of the subcommands 'analyze' or 'flatten' is required.")
		os.Exit(1)
	}

	// Parse the flags
	switch os.Args[1] {
	case "analyze":
		analyzeCommand.Parse(os.Args[2:])
	case "flatten":
		flattenCommand.Parse(os.Args[2:])
	default:
		flag.PrintDefaults()
		os.Exit(1)
	}

	// Check which subcommand was parsed
	if analyzeCommand.Parsed() {
		if *analyzeFilePath == "" {
			analyzeCommand.PrintDefaults()
			os.Exit(1)
		}

		fmt.Printf("File path: %s", *analyzeFilePath)
	}

	if flattenCommand.Parsed() {
		if *flattenFilePath == "" {
			flattenCommand.PrintDefaults()
			os.Exit(1)
		}

		delimiterChoices := map[string]bool{",": true, "^": true, "|": true}
		if _, validChoice := delimiterChoices[*delimiter]; !validChoice {
			flattenCommand.PrintDefaults()
			os.Exit(1)
		}

		files := flatten.Flatten(*flattenFilePath, *delimiter)
		for _, f := range files {
			fmt.Printf("File created: %s\n", f.Path)
		}
		fmt.Printf("The file: %s has been flattened using the delimiter: %s\n", *flattenFilePath, *delimiter)
	}
}
