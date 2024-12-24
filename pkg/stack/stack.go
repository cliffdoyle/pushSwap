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
