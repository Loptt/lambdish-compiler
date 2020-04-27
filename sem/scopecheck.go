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

	fes.Pop()
	
	return nil
}

func scopeCheckStatement(statement ast.Statement, fes *dir.FuncEntryStack, funcdir *dir.FuncDirectory) error {
	if id, ok := statement.(*ast.Id); ok {
		return scopeCheckId(id, fes, funcdir)
	} else if fcall, ok := statement.(*ast.FunctionCall); ok {
		return scopeCheckFunctionCall(fcall, fes, funcdir)
	} else if lambda, ok := statement.(*ast.Lambda); ok {
		return scopeCheckLambda(lambda, fes, funcdir)
	} else if cl, ok := statement.(*ast.ConstantList); ok {
		return scopeCheckConstantList(cl, fes, funcdir)
	} else if _, ok := statement.(*ast.ConstantValue); ok {
		return nil
	}

	return errutil.Newf("Statemnt cannot be casted to any valid form")
}

func scopeCheckId(id *ast.Id, fes *dir.FuncEntryStack, funcdir *dir.FuncDirectory) error {
	if idExistsInFuncStack(id, fes) {
		return nil
	}
	if idExistsInFuncDir(id, funcdir) {
		return nil
	}
	return errutil.Newf("Id %s not declared in this scope", id)
}

func scopeCheckFunctionCall(fcall *ast.FunctionCall, fes *dir.FuncEntryStack, funcdir *dir.FuncDirectory) error {
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

func scopeCheckLambda(lambda *ast.Lambda, fes *dir.FuncEntryStack, funcdir *dir.FuncDirectory) error {
	// Get the FuncEntry for the lambda in the top of the FuncEntry stack
	if lambdaEntry := fes.Top().GetLambdaEntryById(lambda.Id()); lambdaEntry != nil {
		// Add the new func entry to the top of the stack and then call scopeCheckStatement
		// with the updated stack
		fes.Push(lambdaEntry)
		if err := scopeCheckStatement(lambda.Statement(), fes, funcdir); err != nil {
			return err
		}
		// After checking pop the FuncEntry 
		fes.Pop()

		return nil
	}
	return errutil.Newf("Lambda not in lambda list of top of FuncEntry stack")
}

func scopeCheckConstantList(cl *ast.ConstantList, fes *dir.FuncEntryStack, funcdir *dir.FuncDirectory) error {
	for _, s := range cl.Contents() {
		if err := scopeCheckStatement(s, fes, funcdir); err != nil {
			return err
		} 
	}
	return nil
}