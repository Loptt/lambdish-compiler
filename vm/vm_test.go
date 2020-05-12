package vm

import (
	"fmt"
	"testing"
)

func TestLoadProgram(t *testing.T) {
	tests := []string{
		"tests/test1.obj",
	}

	for _, test := range tests {
		vm := NewVirtualMachine()

		err := vm.loadProgram(test)
		if err != nil {
			t.Fatalf("Could not load program: %v", err)
		}

		fmt.Printf("%s\n", vm)
	}
}