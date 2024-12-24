package stack

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