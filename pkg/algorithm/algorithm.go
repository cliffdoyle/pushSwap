package algorithm

import (
	"fmt"
	"math"
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

// SetTargetA sets target nodes for each node in stack 'a' by finding the largest
// number in stack 'b' that is smaller than the current node in 'a'.
// If no such number exists, it targets the maximum value in stack 'b'.
func SetTargetA(a, b *stack.Stack) {
	currentA := a.Head
	for currentA != nil {
		// Initialize with smallest possible integer to find largest valid match
		bestMatchValue := math.MinInt
		var targetNode *stack.StackNode

		// Search through stack B for the largest number smaller than currentA
		currentB := b.Head
		for currentB != nil {
			// Look for numbers in B that are smaller than current A
			// but larger than our current best match
			if currentB.Nbr < currentA.Nbr && currentB.Nbr > bestMatchValue {
				bestMatchValue = currentB.Nbr
				targetNode = currentB
			}
			currentB = currentB.Next
		}

		// If no valid target found, default to maximum value in stack B
		if targetNode == nil {
			maxNode, _ := b.FindMaxNode()
			currentA.Target_node = maxNode
		} else {
			currentA.Target_node = targetNode
		}
		currentA = currentA.Next
	}
}

// SetTargetB sets target nodes for each node in stack 'b' by finding the smallest
// number in stack 'a' that is larger than the current node in 'b'.
// If no such number exists, it targets the maximum value in stack 'a'.
func SetTargetB(a, b *stack.Stack) {
	currentB := b.Head
	for currentB != nil {
		// Initialize with largest possible integer to find smallest valid match
		bestMatchValue := math.MaxInt
		var targetNode *stack.StackNode

		// Search through stack A for the smallest number larger than currentB
		currentA := a.Head
		for currentA != nil {
			// Look for numbers in A that are larger than current B
			// but smaller than our current best match
			if currentA.Nbr > currentB.Nbr && currentA.Nbr < bestMatchValue {
				bestMatchValue = currentA.Nbr
				targetNode = currentA
			}
			currentA = currentA.Next
		}

		// If no valid target found, default to maximum value in stack A
		if targetNode == nil {
			maxNode, _ := a.FindMaxNode()
			currentB.Target_node = maxNode
		} else {
			currentB.Target_node = targetNode
		}
		currentB = currentB.Next
	}
}

// CostAnalysisA calculates the cost of pushing each node from Stack `a` to Stack `b`.
// The cost is determined based on the relative positions of the current node and its target node.
// - If both nodes are above the median, the function optimizes using shared "RR" operations.
// - If both nodes are below the median, the function optimizes using shared "RRR" operations.
// - If the nodes are on opposite sides of the median, individual rotations are used.
// The calculated cost is stored in the `Push_cost` field of each node in Stack `a`.
func CostAnalysisA(a, b *stack.Stack) {
	currentA := a.Head
	for currentA != nil {
		baseRotateCost := currentA.Index
		baseTargetRotateCost := currentA.Target_node.Index

		// If both nodes are below median, use reverse rotation costs instead
		if !currentA.Above_median {
			baseRotateCost = a.Size() - currentA.Index
		}
		if !currentA.Target_node.Above_median {
			baseTargetRotateCost = b.Size() - currentA.Target_node.Index
		}

		// Calculate the cost with and without synchronized operations
		var finalCost int

		// Case 1: Both nodes are above median - can use RR
		if currentA.Above_median && currentA.Target_node.Above_median {
			// Find the higher index to determine how many individual rotations needed
			maxIndex := max(currentA.Index, currentA.Target_node.Index)
			minIndex := min(currentA.Index, currentA.Target_node.Index)

			// Cost = shared rotations + remaining individual rotations
			finalCost = minIndex             // shared RR operations
			finalCost += maxIndex - minIndex // remaining individual rotations

			// Case 2: Both nodes are below median - can use RRR
		} else if !currentA.Above_median && !currentA.Target_node.Above_median {
			maxReverseRotations := max(a.Size()-currentA.Index, b.Size()-currentA.Target_node.Index)
			minReverseRotations := min(a.Size()-currentA.Index, b.Size()-currentA.Target_node.Index)

			// Cost = shared reverse rotations + remaining individual reverse rotations
			finalCost = minReverseRotations                        // shared RRR operations
			finalCost += maxReverseRotations - minReverseRotations // remaining individual rotations

			// Case 3: Nodes are on opposite sides of median - no synchronization possible
		} else {
			finalCost = baseRotateCost + baseTargetRotateCost
		}

		currentA.Push_cost = finalCost
		currentA = currentA.Next
	}
}
