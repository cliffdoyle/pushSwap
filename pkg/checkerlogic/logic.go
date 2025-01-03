package checkerlogic

import (
	"errors"
	"push-swap/pkg/operations"
	"push-swap/pkg/stack"
)

// ExecuteChecker performs all the instructions on the stacks
// and checks if the stack 'a' is sorted and stack 'b' is empty.
func ExecuteChecker(numbers []int, instructions []string) string {
	a := stack.NewStack() // Stack A
	b := stack.NewStack() // Stack B

	// Initialize stack 'a' with the parsed numbers
	a.InitializeStack(numbers)

	// Process each instruction
	for _, instruction := range instructions {
		if err := ExecuteInstruction(instruction, a, b); err != nil {
			return "Error"
		}
	}

	// Check if stack 'a' is sorted and stack 'b' is empty
	if b.Size() == 0 && a.IsSorted() {
		return "OK"
	}
	return "KO"
}

// ExecuteInstruction processes a single instruction on stacks 'a' and 'b'.
// Returns an error for invalid instructions.
func ExecuteInstruction(instruction string, a, b *stack.Stack) error {
	switch instruction {
	case "pa":
		return operations.Pa(b, a)
	case "pb":
		return operations.Pb(a, b)
	case "sa":
		operations.Sa(a)
	case "sb":
		operations.Sb(b)
	case "ss":
		operations.Ss(a, b)
	case "ra":
		operations.Ra(a)
	case "rb":
		operations.Rb(b)
	case "rr":
		operations.Rr(a, b)
	case "rra":
		operations.Rra(a)
	case "rrb":
		operations.Rrb(b)
	case "rrr":
		operations.Rrr(a, b)
	default:
		return errors.New("invalid instruction")
	}
	return nil
}
