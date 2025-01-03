package main

import (
	"bufio"
	"fmt"
	"os"
	"push-swap/pkg/checkerlogic"
	"push-swap/pkg/validation"
	"strings"
)

// main is the entry point for the checker program.
// It validates input, reads instructions from stdin, and executes the checker logic.
func main() {
	if len(os.Args) != 2 {
		return
	}

	// Parse the input string from command-line arguments
	arg := os.Args[1]
	if len(arg) == 0 {
		fmt.Println()
		return
	}

	// Validate and parse the input numbers
	numbers, err := validation.ParseInput(strings.Fields(arg))
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
	for scanner.Scan() {
		instructions = append(instructions, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return instructions, nil
}
