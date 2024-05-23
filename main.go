package main

import (
	"fmt"
	"gvm/vm"
)

func main() {
	virtualMachine := vm.NewVM(1024, 1)

	virtualMachine.Push(60)
	virtualMachine.Push(10)

	//fmt.Println(virtualMachine.Pop()) // show 20
	//fmt.Println(virtualMachine.Pop()) // show 10

	//virtualMachine.Add()
	virtualMachine.Sub()

	fmt.Println(virtualMachine.Pop()) // should be 50
}
