package vm

import (
	"fmt"
	"strings"
)

type VM struct {
	Memory []int
	Stack  []int
	PC     int // program counter
}

func NewVM(memSize int, numProc int) *VM {
	return &VM{
		Memory: make([]int, memSize),
		Stack:  []int{},
		PC:     0,
	}
}

// Print the current state of the stack
func (vm *VM) PrintStack() {
	if len(vm.Stack) == 0 {
		fmt.Println("Stack is empty.")
		return
	}

	var stackRepresentation strings.Builder
	stackRepresentation.WriteString("Stack [")
	for i, val := range vm.Stack {
		if i > 0 {
			stackRepresentation.WriteString(", ")
		}
		stackRepresentation.WriteString(fmt.Sprintf("%d", val))
	}
	stackRepresentation.WriteString("]")

	fmt.Println(stackRepresentation.String())
}

// Push a value onto the stack
func (vm *VM) Push(value int) {
	vm.Stack = append(vm.Stack, value)
}

// Pop a value onto the stack
func (vm *VM) Pop() int {
	if len(vm.Stack) == 0 {
		panic("stack underflow")
	}

	value := vm.Stack[len(vm.Stack)-1]
	vm.Stack = vm.Stack[:len(vm.Stack)-1]
	return value
}

// Add two value on the stack
func (vm *VM) Add() {
	if len(vm.Stack) < 2 {
		panic("Cannot add, not enough value on the stack.")
	}

	firstOperand := vm.Pop()
	secondOperand := vm.Pop()

	result := firstOperand + secondOperand

	vm.Push(result)
}

// Substract two value on the stack
func (vm *VM) Sub() {
	if len(vm.Stack) < 2 {
		panic("Cannot substract, not enough value on the stack.")
	}

	firstOperand := vm.Pop()
	secondOperand := vm.Pop()

	result := secondOperand - firstOperand

	vm.Push(result)
}

// Multiplies two value on the stack
func (vm *VM) Mult() {
	if len(vm.Stack) < 2 {
		panic("Cannot multiplies, not enough value on the stack.")
	}

	firstOperand := vm.Pop()
	secondOperand := vm.Pop()

	if firstOperand == 0 || secondOperand == 0 {
		vm.Push(0)
		return
	}

	result := secondOperand * firstOperand

	vm.Push(result)
}

// Divide two value on the stack
func (vm *VM) Div() {
	if len(vm.Stack) < 2 {
		panic("Cannot multiplies, not enough value on the stack.")
	}

	firstOperand := vm.Pop()
	secondOperand := vm.Pop()

	if firstOperand == 0 {
		panic("Cannot divide by zero.")
	}

	result := secondOperand / firstOperand

	vm.Push(result)
}
