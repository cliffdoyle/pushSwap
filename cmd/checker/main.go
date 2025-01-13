package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"push-swap/pkg/checkerlogic"
	"push-swap/pkg/validation"
)

// main is the entry point for the checker program.
// It validates input, reads instructions from stdin, and executes the checker logic.
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

	// Validate and parse the input numbers
	numbers, err := validation.ParseInput(numbersSlice)
	if err != nil {
		fmt.Println("Error")
		return
	}

	// Read instructions from stdin
	instructions, err := readInstructions()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error")
		return
	}

	// Execute the checker logic and print the result
	result := checkerlogic.ExecuteChecker(numbers, instructions)
	fmt.Println(result)
}

// readInstructions reads a series of instructions from stdin.
// Each instruction is expected to be on a new line.
func readInstructions() ([]string, error) {
	instructions := []string{}
	scanner := bufio.NewScanner(os.Stdin)
	previousEmpty := false

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			// If there's any consecutive empty line(s), return an error
			if previousEmpty {
				return nil, fmt.Errorf("consecutive empty lines are not allowed")
			}
			previousEmpty = true
		} else {
			instructions = append(instructions, scanner.Text())
			previousEmpty = false
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return instructions, nil
}
