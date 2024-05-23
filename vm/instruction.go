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
	case "ADD":
		vm.Add()
	case "SUB":
		vm.Sub()
	case "MULT":
		vm.Mult()
	case "DIV":
		vm.Div()
	case "MOD":
		vm.Mod()
	case "AND":
		vm.And()
	case "OR":
		vm.Or()
	case "NOT":
		vm.Not()
	case "LOAD":
		vm.Load()
	case "STORE":
		vm.Store()
	default:
		panic("unknown instruction")
	}
}
