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
