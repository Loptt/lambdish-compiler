package ic

import (
	"github.com/Loptt/lambdish-compiler/ast"
	"github.com/Loptt/lambdish-compiler/dir"
	"github.com/Loptt/lambdish-compiler/mem"
)

func generateAddressesProgram(program *ast.Program, funcdir *dir.FuncDirectory, vm *mem.VirtualMemory) error {
	for _, fe := range funcdir.Table() {
		if err := generateAddressesFuncEntry(fe, funcdir, vm); err != nil {
			return err
		}
	}

	return nil
}

func generateAddressesFuncEntry(fe *dir.FuncEntry, funcdir *dir.FuncDirectory, vm *mem.VirtualMemory) error {
	for _, ve := range fe.VarDir().Table() {
		ve.SetAddress(vm.GetNextLocal())
	}

	for _, l := range fe.Lambdas() {
		if err := generateAddressesFuncEntry(l, funcdir, vm); err != nil {
			return err
		}
	}

	return nil
}
