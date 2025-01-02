package algorithm

import (
	"fmt"
	"push-swap/pkg/operations"
	"push-swap/pkg/stack"
)

// Sort sorts the stack 'a' using the push-swap algorithm and returns a list of operations.
func Sort(numbers []int) []string {
	// Create two new stacks: 'a' and 'b'
	a := stack.NewStack()
	b := stack.NewStack()

	// List to store the operations performed during the sorting process
	operationsList := []string{}

	// Initialize the stack 'a' with the input numbers
	a.InitializeStack(numbers)

	// Print the initial state of the stack 'a'
	a.PrintStack()

	// If the stack is already sorted or has fewer than 2 elements, return an empty list
	// of operations (no need to sort)
	if a.IsSorted() || a.Size() < 2 {
		a.PrintStack() // Print the final state of the stack (no changes made)
		return operationsList
	}

	// Base case: if the size of 'a' is 2, swap the top two elements (only needed if unsorted)
	if a.Size() == 2 {
		operations.Sa(a)                              // Swap the top two elements of stack 'a'
		operationsList = append(operationsList, "sa") // Record the operation
		a.PrintStack()                                // Print the final state of stack 'a'
		return operationsList
	}

	// Base case: if the size of 'a' is 3, use the SortThree function to sort it
	// This function handles sorting exactly 3 elements
	if a.Size() == 3 {
		operationsList = append(operationsList, SortThree(a)...) // Add the operations from SortThree
		a.PrintStack()                                           // Print the final state of stack 'a'
		return operationsList
	}

	// If the stack is not sorted and has more than 3 elements, proceed with the sorting logic
	if !a.IsSorted() {
		return []string{} // Placeholder for further sorting logic (not yet implemented)
	}

	// Print the final state of stack 'a' and stack 'b' (though no operations are done here)
	a.PrintStack()
	b.PrintStack()

	// Print the size of stack 'a' for debugging purposes
	fmt.Println(a.Size())

	// Return the list of operations performed (empty in this case, awaiting further logic)
	return operationsList
}

// SortThree sorts a stack containing exactly three elements.
// It returns a list of operations performed to achieve the sorted order.
func SortThree(a *stack.Stack) []string {
	operationList := []string{} // List to store the operations performed

	// Find the node with the maximum value in the stack
	biggestNode, _ := a.FindMaxNode()

	// Case 1: The largest element is at the top of the stack
	if biggestNode == a.Head {
		operations.Ra(a)                            // Rotate the stack (move top element to bottom)
		operationList = append(operationList, "ra") // Record the operation

		// Case 2: The largest element is in the second position
	} else if a.Head.Next == biggestNode {
		operations.Rra(a)                            // Reverse rotate the stack (move bottom element to top)
		operationList = append(operationList, "rra") // Record the operation
	}

	// Check if the top two elements are out of order
	if a.Head.Nbr > a.Head.Next.Nbr {
		operations.Sa(a)                            // Swap the top two elements
		operationList = append(operationList, "sa") // Record the operation
	}

	return operationList // Return the list of operations performed
}
