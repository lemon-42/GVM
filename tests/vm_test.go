package tests

import (
	"gvm/vm"
	"testing"
)

// Push/Pop instruction test function
func TestPushPop(t *testing.T) {
	v := vm.NewVM(1024, 1)

	v.Push(101)
	if v.Pop() != 101 {
		t.Errorf("Expected 101")
	}

	v.Push(42)
	v.Push(69)
	if v.Pop() != 69 {
		t.Errorf("Expected 69")
	}
	if v.Pop() != 42 {
		t.Errorf("Expected 42")
	}
}
