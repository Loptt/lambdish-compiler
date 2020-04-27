package sem

import (
	"github.com/Loptt/lambdish-compiler/dir"
	"github.com/Loptt/lambdish-compiler/ast"
	"github.com/mewkiz/pkg/errutil"
)

func scopeCheckProgram(program *ast.Program, funcdir *dir.FuncDirectory) error {
	for _, f := range program.Functions() {
		if err := scopeCheckFunction(f, funcdir); err != nil {
			return err
		}
	}

	return nil
}

func scopeCheckFunction(function *ast.Function, funcdir *dir.FuncDirectory) error {
	fe := funcdir.Get(function.Key())
	if fe == nil {
		return errutil.Newf("Function entry %+v not found in FuncDirectory", fe)
	}

	fes := dir.NewFuncEntryStack()
	fes.Push(fe)

	if err := scopeCheckStatement(function.Statement(), fes, funcdir); err != nil {
		return err
	}
	
	return nil
}

func scopeCheckStatement(statement ast.Statement, fes *dir.FuncEntryStack, funcdir *dir.FuncDirectory) error {
	if id, ok := statement.(*ast.Id); ok {
		if idExistsInFuncStack(id, fes) {
			return nil
		}
		// TODO: check if id exists in funcdir
		return errutil.Newf("Id %s not declared in this scope", id)
	} else if fcall, ok := statement.(*ast.FunctionCall); ok {
		if err := scopeCheckStatement(fcall.Statement(), fes, funcdir); err != nil {
			return err
		}

		for _, arg := range fcall.Args() {
			if err := scopeCheckStatement(arg, fes, funcdir); err != nil {
				return err
			}
		}
		return nil
	} 

	return nil
}