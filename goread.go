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

	filename := os.Args[1]

	if *standardInput {
		inputBytes, err := io.ReadAll(os.Stdin)
		if err != nil {
			fmt.Println("Error reading from stdin:", err)
			os.Exit(1)
		}

		fmt.Println(string(inputBytes))	
	}

	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}

	fmt.Println(string(data))
}			
