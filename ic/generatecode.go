package ic

import (
	"github.com/Loptt/lambdish-compiler/ast"
	"github.com/Loptt/lambdish-compiler/dir"
	"github.com/Loptt/lambdish-compiler/mem"
	"github.com/Loptt/lambdish-compiler/sem"
	"github.com/mewkiz/pkg/errutil"
)

func generateCodeProgram(program *ast.Program, ctx, *GenerationContext) error {
	for _, function := range program.Functions() {
		if err := generateCodeFunction(function, funcdir, semcube, gen, vm); err != nil {
			return err
		}
	}

	return nil
}

func generateCodeFunction(function *ast.Function, ctx, *GenerationContext) error {
	fe := ctx.funcdir.Get(function.Key())

	fes := dir.NewFuncEntryStack()
	fes.Push(fe)

	if err := generateCodeStatement(fe.Statement(), fes, ctx); err != nil {
		return err
	}

	return nil
}

func generateCodeStatement(statement ast.Statement, fes *FuncEntryStack, ctx *GenerationContext) error {
	if _, ok := statement.(*ast.Id); ok {
		// TOOD: Implement code generation for Id
		return nil
	} else if fcall, ok := statement.(*ast.FunctionCall); ok {
		return generateCodeFunctionCall(fcall, fes, funcdir, semcube)
	} else if lambda, ok := statement.(*ast.Lambda); ok {
		// TOOD: Implement code generation for Id
		return nil
	} else if cl, ok := statement.(*ast.ConstantList); ok {
		// TOOD: Implement code generation for Id
		return nil
	} else if _, ok := statement.(*ast.ConstantValue); ok {
		// TOOD: Implement code generation for Id
		return nil
	}

	return errutil.Newf("Statement cannot be casted to any valid form")
}

func generateCodeFunctionCall(fcall *ast.FunctionCall, fes *FuncEntryStack, ctx *GenerationContext) error {
	if id, ok := fcall.Statement().(*ast.Id); ok {
		if sem.isReservedFunction(id.String()) {
			if err := generateCodeReservedFunctionCall(id, fcall, funcdir, semcube, gen, fes); err != nil {
				return err
			}
		} else {
			// TODO: Generate code with user defined function calls
		}
	} else {
		// TODO: Generate code when function call is not Id
		return nil
	}
}

func generateCodeReservedFunctionCall(id *ast.Id, fcall *ast.FunctionCall, fes *FuncEntryStack, ctx *GenerationContext) error {
	switch id.String() {
	case "+":
		if err := generateAdd(fcall, fes, ctx); err != nil {
			return err
		}
	}
}

func generateAdd(fcall *ast.FunctionCall, fes *FuncEntryStack, ctx *GenerationContext) error {
	args := fcall.Args()

	// First we get the left and right operands from the args array
	// TODO: Check if they are in the right order to ensure left associativity
	lop := args[0]
	rop := args[1]

	// We check if both operands are ids
	idlop, ok := lop.(*ast.Id)
	if !ok {
		return errutil.Newf("Only Ids supported")
	}

	idrop, ok := rop.(*ast.Id)
	if !ok {
		return errutil.Newf("Only Ids supported")
	}

	gen.Generate(Add, )
}
