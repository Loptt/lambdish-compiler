package main

import (
	"fmt"
	"os"

	"github.com/Loptt/lambdish-compiler/vm"
)

//usage The correct use of the rlamb command
func usage() {
	fmt.Printf("Usage: rlamb <lambdish object file>\n")
}

//run Function that initializes the new Virtual Machine and loads the program file and prints the output of the program
func run(file string) error {
	machine := vm.NewVirtualMachine()

	err := machine.LoadProgram(file)
	if err != nil {
		return err
	}

	err = machine.Run()
	if err != nil {
		return err
	}

	return nil
}

func main() {
	if len(os.Args) < 2 {
		usage()
		return
	}

	file := os.Args[1]

	err := run(file)
	if err != nil {
		fmt.Printf("Runtime %v\n", err)
		return
	}
}
