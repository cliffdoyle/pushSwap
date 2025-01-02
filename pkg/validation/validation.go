package validation

import (
	"errors"
	"strconv"
)

// ParseInput takes a slice of strings, converts them to integers, and validates them
// It returns a slice of integers or an error if any input is invalid or duplicated
func ParseInput(input []string) ([]int, error) {
	// Create a slice to hold the converted integers
	numbers := make([]int, len(input))
	// Create a map to track seen numbers and detect duplicates
	seen := map[int]bool{}

	// Iterate over the input strings
	for i, str := range input {
		// Convert the current string into an integer
		num, err := strconv.Atoi(str)
		// Return an error if the string is not a valid number
		if err != nil {
			return nil, errors.New("invalid number: " + str)
		}
		// Check for duplicates
		if seen[num] {
			return nil, errors.New("duplicate number: " + str)
		}
		// Mark the number as seen
		seen[num] = true
		// Store the number in the numbers slice
		numbers[i] = num
	}
	// Return the slice of numbers with nil error
	return numbers, nil
}
