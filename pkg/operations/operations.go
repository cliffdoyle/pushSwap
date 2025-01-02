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
