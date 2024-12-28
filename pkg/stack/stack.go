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
