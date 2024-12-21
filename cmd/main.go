package main

import (
	"fmt"
	"os"
	"push-swap/pkg/validation"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		return
	}

	arg := os.Args[1] // Read input from command line
	if arg == "" {
		fmt.Println()
		return
	}

	// Validate and parse input
	numbers, err := validation.ParseInput(strings.Fields(arg))
	if err != nil {
		fmt.Println("Error")
		return
	}
}
