package main

import (
	"fmt"
	"log"
)

func kataSolution(input string) string {
	// Write your code here
	return input
}

func main() {
	// Read the input from standard input
	var input string
	_, err := fmt.Scanln(&input)
	if err != nil {
		log.Fatal(err)
	}

	// Print the solution to standard output
	_, err = fmt.Print(kataSolution(input))
	if err != nil {
		log.Fatal(err)
	}
}
