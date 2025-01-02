package stack

import "fmt"

type StackNode struct {
	Nbr          int
	Index        int
	Push_cost    int
	Above_median int
	Target_node  *StackNode
	Next         *StackNode
	Prev         *StackNode
}

type Stack struct {
	Head *StackNode // Pointer to the top node of the stack
}

// NewStack creates and returns a new empty stack.
func NewStack() *Stack {
	return &Stack{Head: nil} // Initialize an empty stack with Head as nil
}

// Push adds an element to the top of the stack.
func (s *Stack) Push(value int) {
	// Initialize a new node with value and link to the current head
	newNode := &StackNode{Nbr: value, Next: s.Head}
	if s.Head != nil {
		s.Head.Prev = newNode // Update the previous pointer of the current head
	}
	s.Head = newNode // Update the head of the stack to the new node
}

// Pop removes and returns the top element of the stack.
func (s *Stack) Pop() (int, bool) {
	if s.Head == nil {
		return 0, false // Empty stack
	}
	value := s.Head.Nbr
	s.Head = s.Head.Next
	if s.Head != nil {
		s.Head.Prev = nil // Unlink the popped node
	}
	return value, true
}

// AppendNode adds a new node with value n to the end of the stack
func (s *Stack) AppendNode(n int) {
	// Create a new node
	node := &StackNode{Nbr: n}

	if s.Head == nil {
		// If stack is empty, set its head to the new node
		s.Head = node
	} else {
		// If stack is not empty, find the last node
		lastNode := s.Head
		for lastNode.Next != nil {
			lastNode = lastNode.Next // Traverse to the end of the list
		}
		lastNode.Next = node // Link last node's next to new node
		node.Prev = lastNode // Link new node's prev back to last node
	}
}

// InitializeStack initializes stack A with an array of integers
func (s *Stack) InitializeStack(data []int) {
	for _, value := range data {
		s.AppendNode(value) // Append each number to the stack
	}
}

// IsSorted checks if the stack is sorted in ascending order
func (s *Stack) IsSorted() bool {
	if s.Head == nil { // An empty stack is considered sorted
		return true
	}

	current := s.Head
	for current.Next != nil {
		if current.Nbr > current.Next.Nbr { // Compare current with next node
			return false // If current is greater than next, it's not sorted
		}
		current = current.Next // Move to next node
	}
	return true // If all comparisons are fine, it is sorted
}

// Size returns the number of nodes in the stack.
func (s *Stack) Size() int {
	count := 0        // Initialize a counter to zero
	current := s.Head // Start from the head of the stack

	// Traverse through the stack until we reach the end
	for current != nil {
		count++                // Increment count for each node
		current = current.Next // Move to the next node
	}

	return count // Return the total count of nodes
}

// PrintStack prints the values in the stack from top to bottom (head to tail)
func (s *Stack) PrintStack() {
	current := s.Head
	for current != nil {
		fmt.Print(current.Nbr, " ")
		current = current.Next
	}
	fmt.Println()
}

// Swap swaps the first two elements on top of a stack.
func (s *Stack) Swap() {
	if s.Head == nil || s.Head.Next == nil {
		// Do nothing if there are fewer than 2 elements
		return
	}

	first := s.Head
	second := first.Next

	// Adjust the third node's Prev pointer, if it exists
	if second.Next != nil {
		second.Next.Prev = first
	}

	// Swap the nodes
	first.Next = second.Next // First now points to the node after second
	second.Prev = nil        // Second becomes the new head, so its Prev is nil
	second.Next = first      // Second points to first
	first.Prev = second      // First's Prev is now second

	// Update head to point to the new top element (second)
	s.Head = second
}

// Rotate shifts up all elements of the specified stack by 1.
// The first element becomes the last one.
func (s *Stack) Rotate() {
	if s.Head == nil || s.Head.Next == nil {
		// Do nothing if there are fewer than 2 elements
		return
	}

	// Save the first element
	first := s.Head

	// Update head to the second element
	s.Head = first.Next
	s.Head.Prev = nil   // The new head's previous pointer should be nil

	// Traverse to the last node
	lastNode := s.Head
	for lastNode.Next != nil {
		lastNode = lastNode.Next
	}

	// Link the last node to the old head
	lastNode.Next = first
	first.Prev = lastNode
	first.Next = nil // The old head becomes the last node
}

// ReverseRotate shifts down all elements of the specified stack by 1.
// The last element becomes the first one.
func (s *Stack) ReverseRotate() {
	if s.Head == nil || s.Head.Next == nil {
		// Do nothing if there are fewer than 2 elements
		return
	}

	// Traverse to find the last and penultimate nodes
    lastNode := s.Head
    var penultimate *StackNode

    for lastNode.Next != nil {
        penultimate = lastNode
        lastNode = lastNode.Next
    }

	// Update pointers to shift the last element to the front
    if penultimate != nil {
        penultimate.Next = nil // Detach last node from
    }
    lastNode.Next = s.Head      // Link the last node to the current head
    s.Head.Prev = lastNode      // Update the current head's previous pointer
    lastNode.Prev = nil         // The new head's previous pointer should be nil
    s.Head = lastNode           // Update head to the last node
}

// FindMaxNode returns the node with the maximum value in the stack.
// If the stack is empty, it returns an error.
func (s *Stack) FindMaxNode() (*StackNode, error) {
	// Check if the stack is empty
	if s.Head == nil {
		return nil, fmt.Errorf("stack is empty")
	}

	// Initialize maxNode with the head of the stack
	maxNode := s.Head

	// Start checking from the second node
	current := s.Head.Next

	// Traverse the stack to find the maximum value
	for current != nil {
		if current.Nbr > maxNode.Nbr { // Compare current value with maxNode value
			maxNode = current // Update maxNode if a larger value is found
		}
		current = current.Next // Move to the next node
	}
	return maxNode, nil // Return the node with the maximum value
}
