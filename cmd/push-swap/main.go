package main

import (
	"fmt"
	"os"
	"push-swap/pkg/algorithm"
	"push-swap/pkg/validation"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		return
	}

	numbersSlice := []string{}

	if len(os.Args) == 2 {
		arg := os.Args[1] // Read input from command line as one argument
		if arg == "" {
			fmt.Println()
			return
		}
		numbersSlice = strings.Fields(arg)
	} else if len(os.Args) > 2 {
		numbersSlice = os.Args[1:] // Read input from command line as multiple arguments
	}

	// Validate and parse input
	numbers, err := validation.ParseInput(numbersSlice)
	if err != nil {
		fmt.Println("Error")
		return
	}

	// Sort using push_swap algorithm
	operations := algorithm.Sort(numbers)
	for _, op := range operations {
		fmt.Println(op)
	}
}
