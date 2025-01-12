package stack

import (
	"bytes"
	"os"
	"testing"
)

func TestNewStack(t *testing.T) {
	t.Run("NewStack creates an empty stack", func(t *testing.T) {
		s := NewStack()
		if s.Head != nil {
			t.Errorf("expected stack head to be nil, got %v", s.Head)
		}
	})
}

func TestPush(t *testing.T) {
	t.Run("Push adds a single element to an empty stack", func(t *testing.T) {
		s := NewStack()
		s.Push(42)

		if s.Head == nil {
			t.Fatalf("expected stack head to be non-nil")
		}

		if s.Head.Nbr != 42 {
			t.Errorf("expected head value to be 42, got %d", s.Head.Nbr)
		}

		if s.Head.Next != nil {
			t.Errorf("expected head's next to be nil, got %v", s.Head.Next)
		}
	})

	t.Run("Push adds multiple elements and maintains correct order", func(t *testing.T) {
		s := NewStack()
		s.Push(10)
		s.Push(20)

		if s.Head.Nbr != 20 {
			t.Errorf("expected head value to be 20, got %d", s.Head.Nbr)
		}

		if s.Head.Next == nil || s.Head.Next.Nbr != 10 {
			t.Errorf("expected next value to be 10, got %v", s.Head.Next)
		}
	})
}

func TestPop(t *testing.T) {
	t.Run("Pop removes and returns the top element from the stack", func(t *testing.T) {
		s := NewStack()
		s.Push(10)
		s.Push(20)

		value, ok := s.Pop()
		if !ok {
			t.Errorf("expected Pop to succeed, but got ok = false")
		}

		if value != 20 {
			t.Errorf("expected value to be 20, got %d", value)
		}

		if s.Head == nil || s.Head.Nbr != 10 {
			t.Errorf("expected new head value to be 10, got %v", s.Head)
		}
	})

	t.Run("Pop returns false when called on an empty stack", func(t *testing.T) {
		s := NewStack()

		value, ok := s.Pop()
		if ok {
			t.Errorf("expected Pop to fail, but got ok = true")
		}

		if value != 0 {
			t.Errorf("expected value to be 0, got %d", value)
		}
	})
}

func TestAppendNode(t *testing.T) {
	t.Run("AppendNode adds a node to an empty stack", func(t *testing.T) {
		s := NewStack()
		s.AppendNode(10)

		if s.Head == nil {
			t.Fatalf("expected stack head to be non-nil")
		}

		if s.Head.Nbr != 10 {
			t.Errorf("expected head value to be 10, got %d", s.Head.Nbr)
		}

		if s.Head.Next != nil {
			t.Errorf("expected head's next to be nil, got %v", s.Head.Next)
		}
	})

	t.Run("AppendNode adds a node to the end of a non-empty stack", func(t *testing.T) {
		s := NewStack()
		s.AppendNode(10)
		s.AppendNode(20)

		if s.Head.Next == nil || s.Head.Next.Nbr != 20 {
			t.Errorf("expected last node value to be 20, got %v", s.Head.Next)
		}
	})
}

func TestInitializeStack(t *testing.T) {
	t.Run("InitializeStack initializes stack with multiple values", func(t *testing.T) {
		s := NewStack()
		s.InitializeStack([]int{1, 2, 3})

		if s.Head == nil {
			t.Fatalf("expected stack head to be non-nil")
		}

		// Verify the values in the stack
		expectedValues := []int{1, 2, 3}
		current := s.Head
		for i, expected := range expectedValues {
			if current == nil {
				t.Fatalf("expected %d nodes, but stack ended early", len(expectedValues))
			}
			if current.Nbr != expected {
				t.Errorf("expected node %d to have value %d, got %d", i, expected, current.Nbr)
			}
			current = current.Next
		}

		if current != nil {
			t.Errorf("expected stack to have %d nodes, but it has more", len(expectedValues))
		}
	})
}

func TestIsSorted(t *testing.T) {
	t.Run("IsSorted returns true for an empty stack", func(t *testing.T) {
		s := NewStack()
		if !s.IsSorted() {
			t.Errorf("expected IsSorted to return true for an empty stack")
		}
	})

	t.Run("IsSorted returns true for a sorted stack", func(t *testing.T) {
		s := NewStack()
		s.InitializeStack([]int{1, 2, 3, 4, 5})
		if !s.IsSorted() {
			t.Errorf("expected IsSorted to return true for a sorted stack")
		}
	})

	t.Run("IsSorted returns false for an unsorted stack", func(t *testing.T) {
		s := NewStack()
		s.InitializeStack([]int{1, 3, 2, 4})
		if s.IsSorted() {
			t.Errorf("expected IsSorted to return false for an unsorted stack")
		}
	})
}

func TestSize(t *testing.T) {
	t.Run("Size returns 0 for an empty stack", func(t *testing.T) {
		s := NewStack()
		if s.Size() != 0 {
			t.Errorf("expected size to be 0, got %d", s.Size())
		}
	})

	t.Run("Size returns correct size for a non-empty stack", func(t *testing.T) {
		s := NewStack()
		s.InitializeStack([]int{1, 2, 3})
		if s.Size() != 3 {
			t.Errorf("expected size to be 3, got %d", s.Size())
		}
	})
}

func TestPrintStack(t *testing.T) {
	t.Run("PrintStack prints the stack values from head to tail", func(t *testing.T) {
		s := NewStack()
		s.InitializeStack([]int{10, 20, 30, 40, 50})

		// Redirect stdout to a buffer
		oldStdout := os.Stdout
		r, w, _ := os.Pipe()
		os.Stdout = w

		s.PrintStack()

		// Restore stdout and read the captured output
		w.Close()
		os.Stdout = oldStdout
		var buf bytes.Buffer
		_, _ = buf.ReadFrom(r)

		expectedOutput := "10 20 30 40 50 \n"
		actualOutput := buf.String()

		if actualOutput != expectedOutput {
			t.Errorf("expected output: %q, got: %q", expectedOutput, actualOutput)
		}
	})
}

func TestSwap(t *testing.T) {
	t.Run("Swap does nothing on an empty stack", func(t *testing.T) {
		s := NewStack()
		s.Swap()

		if s.Head != nil {
			t.Errorf("expected head to be nil, got %v", s.Head)
		}
	})

	t.Run("Swap does nothing on a stack with one element", func(t *testing.T) {
		s := NewStack()
		s.AppendNode(10)
		s.Swap()

		if s.Head.Nbr != 10 {
			t.Errorf("expected head value to remain 10, got %d", s.Head.Nbr)
		}
		if s.Head.Next != nil {
			t.Errorf("expected head.Next to remain nil")
		}
	})

	t.Run("Swap swaps the first two elements", func(t *testing.T) {
		s := NewStack()
		s.AppendNode(10)
		s.AppendNode(20)
		s.Swap()

		if s.Head.Nbr != 20 {
			t.Errorf("expected head value to be 20, got %d", s.Head.Nbr)
		}
		if s.Head.Next == nil || s.Head.Next.Nbr != 10 {
			t.Errorf("expected second element value to be 10, got %v", s.Head.Next)
		}
	})
}

func TestRotate(t *testing.T) {
	t.Run("Rotate does nothing on an empty stack", func(t *testing.T) {
		s := NewStack()
		s.Rotate()

		if s.Head != nil {
			t.Errorf("expected head to be nil, got %v", s.Head)
		}
	})

	t.Run("Rotate does nothing on a stack with one element", func(t *testing.T) {
		s := NewStack()
		s.AppendNode(10)
		s.Rotate()

		if s.Head.Nbr != 10 {
			t.Errorf("expected head value to remain 10, got %d", s.Head.Nbr)
		}
		if s.Head.Next != nil {
			t.Errorf("expected head.Next to remain nil")
		}
	})

	t.Run("Rotate shifts the first element to the last", func(t *testing.T) {
		s := NewStack()
		s.InitializeStack([]int{10, 20, 30})
		s.Rotate()

		if s.Head.Nbr != 20 {
			t.Errorf("expected head value to be 20, got %d", s.Head.Nbr)
		}

		// Check the new ordering
		values := []int{20, 30, 10}
		current := s.Head
		for i, expected := range values {
			if current == nil {
				t.Fatalf("expected %d nodes, but stack ended early", len(values))
			}
			if current.Nbr != expected {
				t.Errorf("expected node %d to have value %d, got %d", i, expected, current.Nbr)
			}
			current = current.Next
		}
	})
}

func TestReverseRotate(t *testing.T) {
	t.Run("ReverseRotate does nothing on an empty stack", func(t *testing.T) {
		s := NewStack()
		s.ReverseRotate()

		if s.Head != nil {
			t.Errorf("expected head to be nil, got %v", s.Head)
		}
	})

	t.Run("ReverseRotate does nothing on a stack with one element", func(t *testing.T) {
		s := NewStack()
		s.AppendNode(10)
		s.ReverseRotate()

		if s.Head.Nbr != 10 {
			t.Errorf("expected head value to remain 10, got %d", s.Head.Nbr)
		}
		if s.Head.Next != nil {
			t.Errorf("expected head.Next to remain nil")
		}
	})

	t.Run("ReverseRotate shifts the last element to the front", func(t *testing.T) {
		s := NewStack()
		s.InitializeStack([]int{10, 20, 30})
		s.ReverseRotate()

		if s.Head.Nbr != 30 {
			t.Errorf("expected head value to be 30, got %d", s.Head.Nbr)
		}

		// Check the new ordering
		values := []int{30, 10, 20}
		current := s.Head
		for i, expected := range values {
			if current == nil {
				t.Fatalf("expected %d nodes, but stack ended early", len(values))
			}
			if current.Nbr != expected {
				t.Errorf("expected node %d to have value %d, got %d", i, expected, current.Nbr)
			}
			current = current.Next
		}
	})
}

func TestFindMinNode(t *testing.T) {
	t.Run("FindMinNode returns an error if the stack is empty", func(t *testing.T) {
		s := NewStack()
		_, err := s.FindMinNode()
		if err == nil {
			t.Errorf("expected error for empty stack, got nil")
		}
	})

	t.Run("FindMinNode returns the minimum value node", func(t *testing.T) {
		s := NewStack()
		s.InitializeStack([]int{10, 20, 5, 15})
		node, err := s.FindMinNode()

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if node.Nbr != 5 {
			t.Errorf("expected min value to be 5, got %d", node.Nbr)
		}
	})
}

func TestFindMaxNode(t *testing.T) {
	t.Run("FindMaxNode returns an error if the stack is empty", func(t *testing.T) {
		s := NewStack()
		_, err := s.FindMaxNode()
		if err == nil {
			t.Errorf("expected error for empty stack, got nil")
		}
	})

	t.Run("FindMaxNode returns the maximum value node", func(t *testing.T) {
		s := NewStack()
		s.InitializeStack([]int{10, 20, 5, 15})
		node, err := s.FindMaxNode()

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if node.Nbr != 20 {
			t.Errorf("expected max value to be 20, got %d", node.Nbr)
		}
	})
}

func TestCurrentIndex(t *testing.T) {
	t.Run("CurrentIndex assigns indices and Above_median correctly", func(t *testing.T) {
		s := NewStack()
		s.InitializeStack([]int{10, 20, 30, 40, 50})
		s.CurrentIndex()

		current := s.Head
		expectedIndices := []struct {
			Index        int
			Above_median bool
		}{
			{0, true}, {1, true}, {2, true}, {3, false}, {4, false},
		}

		for i, expected := range expectedIndices {
			if current == nil {
				t.Fatalf("stack ended prematurely at index %d", i)
			}
			if current.Index != expected.Index {
				t.Errorf("expected index %d, got %d", expected.Index, current.Index)
			}
			if current.Above_median != expected.Above_median {
				t.Errorf("expected Above_median %v, got %v", expected.Above_median, current.Above_median)
			}
			current = current.Next
		}
	})
}

func TestGetCheapest(t *testing.T) {
	t.Run("GetCheapest returns nil if no node is marked as Cheapest", func(t *testing.T) {
		s := NewStack()
		s.InitializeStack([]int{10, 20, 30, 40, 50})
		node := s.GetCheapest()

		if node != nil {
			t.Errorf("expected nil, got %v", node)
		}
	})

	t.Run("GetCheapest returns the first node marked as Cheapest", func(t *testing.T) {
		s := NewStack()
		s.InitializeStack([]int{10, 20, 30, 40, 50})
		// Mark the third node as Cheapest
		current := s.Head.Next.Next
		current.Cheapest = true

		node := s.GetCheapest()
		if node == nil {
			t.Fatalf("expected a node, got nil")
		}
		if node.Nbr != 30 {
			t.Errorf("expected Cheapest node value to be 30, got %d", node.Nbr)
		}
	})
}