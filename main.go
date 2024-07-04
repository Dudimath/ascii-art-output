package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"asciiart-output/ascii"
)

// formats is a map that maps output format flags to their corresponding filenames.
var formats = map[string]string{
	"standard":   "standard.txt",
	"thinkertoy": "thinkertoy.txt",
	"shadow":     "shadow.txt",
}

func main() {
	if len(os.Args) <= 1 || len(os.Args) > 4 {
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard")
		return
	}
	var inputString string
	var banner string
	var outputFile string
	if len(os.Args) == 2 {
		inputString = os.Args[1]
		banner = "standard"
	}

	providedOutputFileName := false
	if len(os.Args) == 3 {
		if strings.HasPrefix(os.Args[1], "--output=") {
			inputString = os.Args[2]
			banner = "standard"
			providedOutputFileName = true
		}
		if !strings.HasPrefix(os.Args[1], "--output=") && strings.Contains(os.Args[1], "=") {
			fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard")
			return
		} else {
			inputString = os.Args[1]
			banner = strings.ToLower(os.Args[2])
		}
	}

	if len(os.Args) == 4 {
		if strings.HasPrefix(os.Args[1], "--output=") {
			providedOutputFileName = true
		}
		if !strings.HasPrefix(os.Args[1], "--output=") && strings.Contains(os.Args[1], "=") {
			fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard")
			return
		}
		inputString = os.Args[2]
		banner = strings.ToLower(os.Args[3])
	}

	if providedOutputFileName {
		outputFile = os.Args[1][9:]
		if !strings.HasSuffix(outputFile, ".txt") {
			fmt.Println("Error: Output file must have a '.txt' extension.")
			return
		}
	} else {
		// write to default file
		outputFile = "output.txt"
	}

	if len(outputFile) <= 4 {
		fmt.Println("Error: The output file name must be at least 5 characters long.")
		return
	} else if strings.Count(outputFile, ".") > 1 {
		fmt.Println("Error: The output file name must not contain more than one period.")
		return
	}

	if outputFile == "standard.txt" || outputFile == "shadow.txt" || outputFile == "thinkertoy.txt" {
		fmt.Println("Error:This is a banner file and you are not to write to it")
		return
	}

	fileName, ok := formats[banner]
	if !ok {
		fmt.Println("Invalid banner specified.")
		return
	}

	asciMap, err := ascii.ReadASCIIMapFromFile(fileName)
	if err != nil {
		log.Fatalf("Error reading ASCII map: %v", err)
	}

	final, err := ascii.PrintArt(inputString, asciMap)
	if !providedOutputFileName {
		if err != nil {
			log.Printf("Error: %v", err)
		} else {
			fmt.Print(final)
		}
	}

	err = os.WriteFile(outputFile, []byte(final), 0o644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
}
