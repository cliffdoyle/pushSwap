package main

import (
	"fmt"
	"os"
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
}
