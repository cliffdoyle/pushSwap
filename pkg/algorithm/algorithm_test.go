package algorithm

import (
	"push-swap/pkg/stack"
	"reflect"
	"testing"
)

func TestSort(t *testing.T) {
	tests := []struct {
		name           string
		input          []int
		expectedOutput []string
	}{
		{
			name:           "Already sorted stack",
			input:          []int{1, 2, 3, 4, 5},
			expectedOutput: []string{}, // No operations needed
		},
		{
			name:           "Single element stack",
			input:          []int{42},
			expectedOutput: []string{}, // No operations needed
		},
		{
			name:           "Two elements unsorted",
			input:          []int{2, 1},
			expectedOutput: []string{"sa"}, // Swap the two elements
		},
		{
			name:           "Three elements unsorted",
			input:          []int{3, 1, 2},
			expectedOutput: []string{"ra"},
		},
		{
			name:           "Five elements unsorted",
			input:          []int{3, 5, 1, 4, 2},
			expectedOutput: []string{"pb", "pb", "rra", "sa", "rra", "pa", "sa", "pa", "rra", "rra"},
		},
		{
			name:           "Empty stack",
			input:          []int{},
			expectedOutput: []string{}, // No operations needed
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := Sort(tt.input)

			// Check for equality with the expected output
			if !reflect.DeepEqual(output, tt.expectedOutput) {
				t.Errorf("Sort(%v) = %v; want %v", tt.input, output, tt.expectedOutput)
			}
		})
	}
}

func TestSetTargetA(t *testing.T) {
	tests := []struct {
		name          string
		stackA        []int
		stackB        []int
		expectedMatch map[int]int // Maps node values in A to their target node values in B
	}{
		{
			name:          "Normal case with valid matches",
			stackA:        []int{5, 10, 15},
			stackB:        []int{1, 7, 12},
			expectedMatch: map[int]int{5: 1, 10: 7, 15: 12},
		},
		{
			name:          "No smaller values in B",
			stackA:        []int{2, 4, 6},
			stackB:        []int{8, 10, 12},
			expectedMatch: map[int]int{2: 12, 4: 12, 6: 12}, // All target the max in B
		},
		{
			name:          "Mixed case with some matches and some max targets",
			stackA:        []int{3, 8, 14},
			stackB:        []int{5, 10, 20},
			expectedMatch: map[int]int{3: 20, 8: 5, 14: 10},
		},
		{
			name:          "Empty stack B",
			stackA:        []int{1, 2, 3},
			stackB:        []int{},
			expectedMatch: map[int]int{1: 0, 2: 0, 3: 0}, // No target nodes
		},
		{
			name:          "Empty stack A",
			stackA:        []int{},
			stackB:        []int{1, 2, 3},
			expectedMatch: map[int]int{}, // No nodes in A
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Initialize stacks
			a := stack.NewStack()
			a.InitializeStack(tt.stackA)

			b := stack.NewStack()
			b.InitializeStack(tt.stackB)

			// Call SetTargetA
			SetTargetA(a, b)

			// Verify that each node in A has the expected target node
			currentA := a.Head
			for currentA != nil {
				expectedTarget := tt.expectedMatch[currentA.Nbr]
				if expectedTarget == 0 {
					// Check if the Target_node is nil
					if currentA.Target_node != nil {
						t.Errorf("Expected no target for node %d, but got %v", currentA.Nbr, currentA.Target_node)
					}
				} else {
					// Check if the Target_node is not nil
					if currentA.Target_node == nil {
						t.Errorf("Expected a target for node %d, but got nil", currentA.Nbr)
					} else {
						// Compare the expected target number with the actual target number
						if !reflect.DeepEqual(expectedTarget, currentA.Target_node.Nbr) {
							t.Errorf("Mismatch for node %d: expected %d, got %d", currentA.Nbr, expectedTarget, currentA.Target_node.Nbr)
						}
					}
				}
				currentA = currentA.Next
			}
		})
	}
}

func TestCostAnalysisA(t *testing.T) {
	tests := []struct {
		name           string
		stackA         []int
		stackB         []int
		expectedCosts  map[int]int // Maps node values in A to their expected Push_cost
		aboveMedianA   map[int]bool // Maps node values in A to their Above_median status
		aboveMedianB   map[int]bool // Maps node values in B to their Above_median status
	}{
		{
			name:          "Both above median with RR optimization",
			stackA:        []int{2, 4, 6},
			stackB:        []int{1, 3, 5},
			expectedCosts: map[int]int{2: 0, 4: 0, 6: 0},
			aboveMedianA:  map[int]bool{2: true, 4: true, 6: true},
			aboveMedianB:  map[int]bool{1: true, 3: true, 5: true},
		},
		{
			name:          "Both below median with RRR optimization",
			stackA:        []int{6, 4, 2},
			stackB:        []int{5, 3, 1},
			expectedCosts: map[int]int{6: 3, 4: 3, 2: 3},
			aboveMedianA:  map[int]bool{6: false, 4: false, 2: false},
			aboveMedianB:  map[int]bool{5: false, 3: false, 1: false},
		},
		{
			name:          "Mixed case with opposite sides of the median",
			stackA:        []int{1, 2, 3},
			stackB:        []int{4, 5, 6},
			expectedCosts: map[int]int{1: 3, 2: 3, 3: 3},
			aboveMedianA:  map[int]bool{1: true, 2: false, 3: false},
			aboveMedianB:  map[int]bool{4: false, 5: true, 6: true},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Initialize stacks
			a := stack.NewStack()
			a.InitializeStack(tt.stackA)

			b := stack.NewStack()
			b.InitializeStack(tt.stackB)

			// Assign Above_median based on test data
			for node := a.Head; node != nil; node = node.Next {
				node.Above_median = tt.aboveMedianA[node.Nbr]
			}
			for node := b.Head; node != nil; node = node.Next {
				node.Above_median = tt.aboveMedianB[node.Nbr]
			}

			// Assign Target_node for stack A based on matching indices in stack B
			for node := a.Head; node != nil; node = node.Next {
				node.Target_node = b.FindByIndex(node.Index % b.Size())
			}

			// Call CostAnalysisA
			CostAnalysisA(a, b)

			// Verify the Push_cost for each node in A
			actualCosts := map[int]int{}
			for node := a.Head; node != nil; node = node.Next {
				actualCosts[node.Nbr] = node.Push_cost
			}

			// Compare expected and actual costs using reflect.DeepEqual
			if !reflect.DeepEqual(actualCosts, tt.expectedCosts) {
				t.Errorf("Push_cost mismatch: expected %v, got %v", tt.expectedCosts, actualCosts)
			}
		})
	}
}

// TestSetCheapestA tests the SetCheapestA function.
func TestSetCheapestA(t *testing.T) {
	// Create a stack and add nodes with varying Push_cost
	a := stack.NewStack()
	a.InitializeStack([]int{10, 20, 30})
	a.Head.Push_cost = 5
	a.Head.Next.Push_cost = 3
	a.Head.Next.Next.Push_cost = 8

	// Call the function to set the cheapest node
	SetCheapestA(a)

	// Verify that the node with the lowest Push_cost is marked as "Cheapest"
	current := a.Head
	foundCheapest := false
	for current != nil {
		if current.Cheapest {
			if foundCheapest {
				t.Errorf("More than one node marked as cheapest")
			}
			foundCheapest = true
			if current.Push_cost != 3 {
				t.Errorf("Cheapest node not correctly identified. Expected Push_cost 3, got %d", current.Push_cost)
			}
		}
		current = current.Next
	}

	// Ensure at least one node was marked as cheapest
	if !foundCheapest {
		t.Errorf("No node marked as cheapest")
	}
}

