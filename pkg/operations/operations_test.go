package operations

import (
	"push-swap/pkg/stack"
	"testing"
)

func TestPush(t *testing.T) {
	// Create source and destination stacks
	source := stack.NewStack()
	destination := stack.NewStack()

	// Initialize the source stack with some values
	source.InitializeStack([]int{10, 20, 30})

	// Perform Push operation
	err := Push(source, destination)
	if err != nil {
		t.Fatalf("unexpected error during Push: %v", err)
	}

	// Check the top of the destination stack
	if destination.Head == nil || destination.Head.Nbr != 10 {
		t.Errorf("expected top of destination stack to be 10, got: %v", destination.Head)
	}

	// Check the new head of the source stack
	if source.Head == nil || source.Head.Nbr != 20 {
		t.Errorf("expected new top of source stack to be 20, got: %v", source.Head)
	}
}

func TestPushEmptySource(t *testing.T) {
	// Create source and destination stacks
	source := stack.NewStack()
	destination := stack.NewStack()

	// Perform Push operation on an empty source stack
	err := Push(source, destination)
	if err == nil {
		t.Fatal("expected an error when pushing from an empty source stack, but got none")
	}

	// Verify the error message
	expectedError := "source stack is empty, pop operation failed"
	if err.Error() != expectedError {
		t.Errorf("expected error message: %q, got: %q", expectedError, err.Error())
	}
}

func TestPa(t *testing.T) {
	// Create stacks A and B
	a := stack.NewStack()
	b := stack.NewStack()

	// Initialize stack B with some values
	b.InitializeStack([]int{50, 60, 70})

	// Perform Pa operation (Push from B to A)
	err := Pa(b, a)
	if err != nil {
		t.Fatalf("unexpected error during Pa: %v", err)
	}

	// Check the top of stack A
	if a.Head == nil || a.Head.Nbr != 50 {
		t.Errorf("expected top of stack A to be 50, got: %v", a.Head)
	}

	// Check the new head of stack B
	if b.Head == nil || b.Head.Nbr != 60 {
		t.Errorf("expected new top of stack B to be 60, got: %v", b.Head)
	}
}

func TestPb(t *testing.T) {
	// Create stacks A and B
	a := stack.NewStack()
	b := stack.NewStack()

	// Initialize stack A with some values
	a.InitializeStack([]int{80, 90, 100})

	// Perform Pb operation (Push from A to B)
	err := Pb(a, b)
	if err != nil {
		t.Fatalf("unexpected error during Pb: %v", err)
	}

	// Check the top of stack B
	if b.Head == nil || b.Head.Nbr != 80 {
		t.Errorf("expected top of stack B to be 80, got: %v", b.Head)
	}

	// Check the new head of stack A
	if a.Head == nil || a.Head.Nbr != 90 {
		t.Errorf("expected new top of stack A to be 90, got: %v", a.Head)
	}
}

func TestSa(t *testing.T) {
	// Create stack A and initialize it with values
	a := stack.NewStack()
	a.InitializeStack([]int{10, 20, 30})

	// Perform Sa operation (swap top two elements of stack A)
	Sa(a)

	// Check the top two elements of stack A after swap
	if a.Head == nil || a.Head.Nbr != 20 {
		t.Errorf("expected top of stack A to be 20, got: %v", a.Head)
	}
	if a.Head.Next == nil || a.Head.Next.Nbr != 10 {
		t.Errorf("expected second element of stack A to be 10, got: %v", a.Head.Next)
	}
}

func TestSb(t *testing.T) {
	// Create stack B and initialize it with values
	b := stack.NewStack()
	b.InitializeStack([]int{40, 50, 60})

	// Perform Sb operation (swap top two elements of stack B)
	Sb(b)

	// Check the top two elements of stack B after swap
	if b.Head == nil || b.Head.Nbr != 50 {
		t.Errorf("expected top of stack B to be 50, got: %v", b.Head)
	}
	if b.Head.Next == nil || b.Head.Next.Nbr != 40 {
		t.Errorf("expected second element of stack B to be 40, got: %v", b.Head.Next)
	}
}

func TestSs(t *testing.T) {
	// Create stacks A and B and initialize them with values
	a := stack.NewStack()
	b := stack.NewStack()
	a.InitializeStack([]int{70, 80, 90})
	b.InitializeStack([]int{100, 110, 120})

	// Perform Ss operation (swap top two elements of both stacks)
	Ss(a, b)

	// Check the top two elements of stack A after swap
	if a.Head == nil || a.Head.Nbr != 80 {
		t.Errorf("expected top of stack A to be 80, got: %v", a.Head)
	}
	if a.Head.Next == nil || a.Head.Next.Nbr != 70 {
		t.Errorf("expected second element of stack A to be 70, got: %v", a.Head.Next)
	}

	// Check the top two elements of stack B after swap
	if b.Head == nil || b.Head.Nbr != 110 {
		t.Errorf("expected top of stack B to be 110, got: %v", b.Head)
	}
	if b.Head.Next == nil || b.Head.Next.Nbr != 100 {
		t.Errorf("expected second element of stack B to be 100, got: %v", b.Head.Next)
	}
}

func TestRa(t *testing.T) {
	// Create stack A and initialize it with values
	a := stack.NewStack()
	a.InitializeStack([]int{1, 2, 3, 4, 5})

	// Perform Ra operation (rotate stack A)
	Ra(a)

	// Verify the new order of stack A
	expected := []int{2, 3, 4, 5, 1}
	actual := a.ToSlice() // Assuming `ToSlice()` converts the stack to a slice for easy comparison
	if !equalSlices(expected, actual) {
		t.Errorf("expected stack A to be %v, got %v", expected, actual)
	}
}

func TestRb(t *testing.T) {
	// Create stack B and initialize it with values
	b := stack.NewStack()
	b.InitializeStack([]int{10, 20, 30, 40, 50})

	// Perform Rb operation (rotate stack B)
	Rb(b)

	// Verify the new order of stack B
	expected := []int{20, 30, 40, 50, 10}
	actual := b.ToSlice()
	if !equalSlices(expected, actual) {
		t.Errorf("expected stack B to be %v, got %v", expected, actual)
	}
}

func TestRr(t *testing.T) {
	// Create stacks A and B and initialize them with values
	a := stack.NewStack()
	b := stack.NewStack()
	a.InitializeStack([]int{1, 2, 3})
	b.InitializeStack([]int{4, 5, 6})

	// Perform Rr operation (rotate both stacks A and B)
	Rr(a, b)

	// Verify the new order of stack A
	expectedA := []int{2, 3, 1}
	actualA := a.ToSlice()
	if !equalSlices(expectedA, actualA) {
		t.Errorf("expected stack A to be %v, got %v", expectedA, actualA)
	}

	// Verify the new order of stack B
	expectedB := []int{5, 6, 4}
	actualB := b.ToSlice()
	if !equalSlices(expectedB, actualB) {
		t.Errorf("expected stack B to be %v, got %v", expectedB, actualB)
	}
}

func TestRra(t *testing.T) {
	a := stack.NewStack()
	a.InitializeStack([]int{1, 2, 3, 4, 5}) // Stack: [5, 4, 3, 2, 1]

	Rra(a)

	expected := []int{5, 1, 2, 3, 4}
	result := a.ToSlice()

	if !equalSlices(result, expected) {
		t.Errorf("Rra failed, got: %v, expected: %v", result, expected)
	}
}

func TestRrb(t *testing.T) {
	b := stack.NewStack()
	b.InitializeStack([]int{4, 5, 6}) // Stack: [6, 5, 4]

	Rrb(b)

	expected := []int{6, 4, 5}
	result := b.ToSlice()

	if !equalSlices(result, expected) {
		t.Errorf("Rrb failed, got: %v, expected: %v", result, expected)
	}
}

func TestRrr(t *testing.T) {
	a := stack.NewStack()
	b := stack.NewStack()
	a.InitializeStack([]int{1, 2, 3}) // Stack a: [3, 2, 1]
	b.InitializeStack([]int{4, 5, 6}) // Stack b: [6, 5, 4]

	Rrr(a, b)

	expectedA := []int{3, 1, 2}
	expectedB := []int{6, 4, 5}
	resultA := a.ToSlice()
	resultB := b.ToSlice()

	if !equalSlices(resultA, expectedA) {
		t.Errorf("Rrr failed for stack a, got: %v, expected: %v", resultA, expectedA)
	}
	if !equalSlices(resultB, expectedB) {
		t.Errorf("Rrr failed for stack b, got: %v, expected: %v", resultB, expectedB)
	}
}

// Helper function to compare two slices for equality
func equalSlices(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
