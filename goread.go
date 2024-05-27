package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	standardInput := flag.Bool("i", false, "read from stdin")
	numberedOutput := flag.Bool("n", false, "number the lines in the output")

	// Parse the flags
	flag.Parse()

	files := flag.Args()

	// if no files provided and no stdin, exit
	if len(files) < 1 && !*standardInput && !*numberedOutput {	
		fmt.Println("No input files provided.")
		os.Exit(1)
	}

	if *standardInput {
		if *numberedOutput {
			inputBytes, err := io.ReadAll(os.Stdin)
			if err != nil {
				fmt.Println("Error reading from stdin:", err)
				os.Exit(1)
			}
			// divide the file into lines
			text := string(inputBytes)
			lines := strings.Split(text, "\n")
			for i, line := range lines[0 : len(lines)-1] {
				fmt.Printf("%d: %s\n", i+1, line)
			}	
			os.Exit(0)
		}
		inputBytes, err := io.ReadAll(os.Stdin)
		if err != nil {
			fmt.Println("Error reading from stdin:", err)
			os.Exit(1)
		}

		fmt.Println(string(inputBytes))	
	}

	// Print output of the files with each lines numbered

	if *numberedOutput {
		files := flag.Args()
		for _, file := range files {
			inputBytes, err := os.ReadFile(file)
			if err != nil {
				fmt.Printf("Error reading file %s: %v\n (numbered output error)", file, err)
				continue
			}
			// divide the file into lines
			text := string(inputBytes)
			lines := strings.Split(text, "\n")
			for i, line := range lines[0 : len(lines)-1] {
				fmt.Printf("%d: %s\n", i+1, line)
			}	
		}
	
			os.Exit(0)
	}
		
		
	if flag.NArg() < 1 {
		os.Exit(1)
	}

	// concatenate the contents of all provided files
	for _, file := range files { // start from 1 because 0 is the name of the program
		data, err := os.ReadFile(file)
		if err != nil {
			fmt.Printf("Error reading file %s: %v\n (concatenate error)", file, err)
			continue
		}
		fmt.Print(string(data))
	}

}			
