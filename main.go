package main

import (
	"fmt"
	"gvm/vm"
)

func main() {
	virtualMachine := vm.NewVM(15, 1)

	instructions := []vm.Instruction{
		{OpCode: "PUSH", Operand: 5},
		{OpCode: "PUSH", Operand: 20},
		{OpCode: "STORE", Operand: 0},
	}

	for _, instruction := range instructions {
		virtualMachine.Execute(instruction)
	}

	instructions = []vm.Instruction{
		{OpCode: "PUSH", Operand: 5},
		{OpCode: "LOAD", Operand: 0},
	}

	for _, instruction := range instructions {
		virtualMachine.Execute(instruction)
	}

	fmt.Println("Value at memory address 5 : ", virtualMachine.Pop())
}
