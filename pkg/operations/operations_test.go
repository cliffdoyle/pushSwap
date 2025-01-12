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
