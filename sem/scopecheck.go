package sem

import (
	"github.com/Loptt/lambdish-compiler/ast"
	"github.com/Loptt/lambdish-compiler/dir"
	"github.com/mewkiz/pkg/errutil"
)

//scopeCheckProgram starts the scope checking for the whole program
func scopeCheckProgram(program *ast.Program, funcdir *dir.FuncDirectory, semcube *SemanticCube) error {
	for _, f := range program.Functions() {
		if err := scopeCheckFunction(f, funcdir, semcube); err != nil {
			return err
		}
	}

	fes := dir.NewFuncEntryStack()
	fes.Push(funcdir.Get("main"))
	if err := scopeCheckFunctionCall(program.Call(), fes, funcdir, semcube); err != nil {
		return err
	}

	return nil
}

//scopeCheckFunction verifies the function is added to the func directory and the checks its statement
func scopeCheckFunction(function *ast.Function, funcdir *dir.FuncDirectory, semcube *SemanticCube) error {
	fe := funcdir.Get(function.Key())
	if fe == nil {
		return errutil.NewNoPosf("%+v: Function entry %+v not found in FuncDirectory", function.Token(), fe)
	}

	fes := dir.NewFuncEntryStack()
	fes.Push(fe)

	if err := scopeCheckStatement(function.Statement(), fes, funcdir, semcube); err != nil {
		return err
	}

	fes.Pop()

	return nil
}

//scopeCheckStatement calls the corresponding function to check the statement depending on the type
func scopeCheckStatement(statement ast.Statement, fes *dir.FuncEntryStack, funcdir *dir.FuncDirectory, semcube *SemanticCube) error {
	if id, ok := statement.(*ast.Id); ok {
		return scopeCheckID(id, fes, funcdir)
	} else if fcall, ok := statement.(*ast.FunctionCall); ok {
		return scopeCheckFunctionCall(fcall, fes, funcdir, semcube)
	} else if lambda, ok := statement.(*ast.Lambda); ok {
		return scopeCheckLambda(lambda, fes, funcdir, semcube)
	} else if cl, ok := statement.(*ast.ConstantList); ok {
		return scopeCheckConstantList(cl, fes, funcdir, semcube)
	} else if _, ok := statement.(*ast.ConstantValue); ok {
		return nil
	}

	return errutil.NewNoPosf("Statement cannot be casted to any valid form")
}

//scopeCheckID
func scopeCheckID(id *ast.Id, fes *dir.FuncEntryStack, funcdir *dir.FuncDirectory) error {
	if idExistsInFuncStack(id, fes) {
		return nil
	}
	if idExistsInFuncDir(id, funcdir) {
		return nil
	}
	if IsReservedFunction(id.String()) {
		return nil
	}
	return errutil.NewNoPosf("%+v: Id %s not declared in this scope", id.Token(), id.String())
}

//scopeCheckFunctionCall
func scopeCheckFunctionCall(fcall *ast.FunctionCall, fes *dir.FuncEntryStack, funcdir *dir.FuncDirectory, semcube *SemanticCube) error {
	if err := scopeCheckStatement(fcall.Statement(), fes, funcdir, semcube); err != nil {
		return err
	}

	for _, arg := range fcall.Args() {
		if err := scopeCheckStatement(arg, fes, funcdir, semcube); err != nil {
			return err
		}
	}
	return nil
}

//scopeCheckLambda
func scopeCheckLambda(lambda *ast.Lambda, fes *dir.FuncEntryStack, funcdir *dir.FuncDirectory, semcube *SemanticCube) error {
	// Get the FuncEntry for the lambda in the top of the FuncEntry stack
	if lambdaEntry := fes.Top().GetLambdaEntryById(lambda.Id()); lambdaEntry != nil {
		// Add the new func entry to the top of the stack and then call scopeCheckStatement
		// with the updated stack
		fes.Push(lambdaEntry)
		if err := scopeCheckStatement(lambda.Statement(), fes, funcdir, semcube); err != nil {
			return err
		}
		// After checking pop the FuncEntry
		fes.Pop()

		return nil
	}
	return errutil.NewNoPosf("%+v: Lambda not in lambda list of top of FuncEntry stack", lambda.Token())
}

//scopeCheckConstantList
func scopeCheckConstantList(cl *ast.ConstantList, fes *dir.FuncEntryStack, funcdir *dir.FuncDirectory, semcube *SemanticCube) error {
	for _, s := range cl.Contents() {
		if err := scopeCheckStatement(s, fes, funcdir, semcube); err != nil {
			return err
		}
	}
	return nil
}
