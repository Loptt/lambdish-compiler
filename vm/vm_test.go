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

		err := vm.LoadProgram(test)
		if err != nil {
			t.Fatalf("Could not load program: %v", err)
		}

		//fmt.Printf("%s\n", vm)
	}
}

func TestRunProgram(t *testing.T) {
	tests := []string{
		"tests/oddmatrix.obj",
	}

	for _, test := range tests {
		vm := NewVirtualMachine()

		err := vm.LoadProgram(test)
		if err != nil {
			t.Fatalf("Could not load program: %v", err)
		}

		err = vm.Run()
		if err != nil {
			t.Fatalf("Runtime Error: %v", err)
		}

		fmt.Printf("%s\n", vm)
	}
}
