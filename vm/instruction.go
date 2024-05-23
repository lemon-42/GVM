package vm

type Instruction struct {
	OpCode  string
	Operand int
}

func (vm *VM) Execute(instruction Instruction) {
	switch instruction.OpCode {
	case "PUSH":
		vm.Push(instruction.Operand)
	case "POP":
		vm.Pop()
	default:
		panic("unknown instruction")
	}
}
