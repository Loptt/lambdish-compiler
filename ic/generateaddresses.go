package ic

import (
	"github.com/Loptt/lambdish-compiler/ast"
	"github.com/Loptt/lambdish-compiler/dir"
)

func generateAddressesProgram(program *ast.Program, ctx *GenerationContext) error {
	for _, fe := range ctx.funcdir.Table() {
		if err := generateAddressesFuncEntry(fe, ctx); err != nil {
			return err
		}
	}

	return nil
}

func generateAddressesFuncEntry(fe *dir.FuncEntry, ctx *GenerationContext) error {
	ctx.vm.ResetLocal()
	for _, ve := range fe.VarDir().Table() {
		ve.SetAddress(ctx.vm.GetNextLocal())
	}

	for _, l := range fe.Lambdas() {
		if err := generateAddressesFuncEntry(l, ctx); err != nil {
			return err
		}
	}

	return nil
}
 
func generateAddressesFunction(function *ast.Function)