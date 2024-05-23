package main

import (
	"fmt"
	"gvm/vm"
)

func main() {
	virtualMachine := vm.NewVM(1024, 1)

	virtualMachine.Push(0)
	virtualMachine.Push(5)

	virtualMachine.PrintStack()

	//fmt.Println(virtualMachine.Pop())
	//fmt.Println(virtualMachine.Pop())

	//virtualMachine.Add()
	//virtualMachine.Sub()
	virtualMachine.Mult()
	//virtualMachine.Div()

	virtualMachine.PrintStack()

	fmt.Println(virtualMachine.Pop())
}
