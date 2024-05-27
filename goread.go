package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	standardInput := flag.Bool("i", false, "read from stdin")
	flag.Parse()

	if *standardInput {
		inputBytes, err := io.ReadAll(os.Stdin)
		if err != nil {
			fmt.Println("Error reading from stdin:", err)
			os.Exit(1)
		}

		fmt.Println(string(inputBytes))	
	}

	if flag.NArg() < 1 {
		// fmt.Println("Usage: goread [filename]")
		os.Exit(1)
	}

	filename := os.Args[1]

	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}

	fmt.Println(string(data))
}			
