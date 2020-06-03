package main

import (
	"fmt"
	"os"

	"github.com/Loptt/lambdish-compiler/vm"
)

func usage() {
	fmt.Printf("Usage: rlamb <lambdish object file>\n")
}

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
