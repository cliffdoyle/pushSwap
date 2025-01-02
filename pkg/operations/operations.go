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
	pa=func(b,a *stack.Stack) error {
		return Push(b,a)
	}
	pb=func(a,b *stack.Stack) error {
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
