package operations

import (
	"errors"
	"push-swap/pkg/stack"
)

//Push moves the top element from the source stack to the destination stack
func Push(source, destination *stack.Stack) error {
	topElement, ok := source.Pop()
	if !ok {
		return errors.New("source stack is empty, pop operation failed")
	}
	destination.Push(topElement)
	return nil
}
//Specific aliases
var (
	Pa=func(b,a *stack.Stack) error {
		return Push(b,a)
	}
	Pb=func(a,b *stack.Stack) error {
		return Push(a,b)
	}
)

//Sa swaps the top two elements of the stack
func Sa(a *stack.Stack) {
	a.Swap()
}

//Sb swaps the top two elements of the stack
func Sb(b *stack.Stack) {
	b.Swap()
}

//Ss swaps the top two elements of both stacks
func Ss(a,b *stack.Stack) {
	a.Swap()
	b.Swap()
}
//Ra rotates the stack by moving the top element to the bottom
func Ra(a *stack.Stack) {
	a.Rotate()
}

//Rb rotates the stack by moving the top element to the bottom
func Rb(b *stack.Stack) {
	b.Rotate()
}
//Rr rotates both stacks a and stack b by shifting all elements up by 1
func Rr(a,b *stack.Stack) {
	a.Rotate()
	b.Rotate()
}
//Rra rotates the stack a by moving the bottom element to the top
func Rra(a *stack.Stack) {
	a.ReverseRotate()
}
//Rrb rotates the stack b by moving the bottom element to the top
func Rrb(b *stack.Stack) {
	b.ReverseRotate()
}

//Rrr rotates both stacks a and stack b by shifting all elements down by 1
func Rrr(a,b *stack.Stack) {
	a.ReverseRotate()
	b.ReverseRotate()
}

// RotateBoth rotates both stacks up together until either target or cheapest node is at top
func RotateBoth(a, b *stack.Stack, cheapestNode *stack.StackNode) []string {
	operationsList := []string{}
	for b.Head != cheapestNode.Target_node && a.Head != cheapestNode {
		Rr(a, b) // Rotate both stacks up one position
		operationsList = append(operationsList, "rr")
		a.CurrentIndex() // Update indices after rotation
		b.CurrentIndex()
	}
	return operationsList
}

// RevRotateBoth rotates both stacks down together until either target or cheapest node is at top
func RevRotateBoth(a, b *stack.Stack, cheapestNode *stack.StackNode) []string {
	operationsList := []string{}
	for b.Head != cheapestNode.Target_node && a.Head != cheapestNode {
		Rrr(a, b) // Rotate both stacks down one position
		operationsList = append(operationsList, "rrr")
		a.CurrentIndex() // Update indices after rotation
		b.CurrentIndex()
	}
	return operationsList
}

