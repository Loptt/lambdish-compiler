package sem

import (
	"github.com/Loptt/lambdish-compiler/ast"
	"github.com/Loptt/lambdish-compiler/dir"
	"github.com/mewkiz/pkg/errutil"
)

//typeCheckProgram: of the program
func typeCheckProgram(program *ast.Program, funcdir *dir.FuncDirectory, semcube *SemanticCube) error {
	for _, f := range program.Functions() {
		if err := typeCheckFunction(f, funcdir, semcube); err != nil {
			return err
		}
	}

	fes := dir.NewFuncEntryStack()

	if err := typeCheckFunctionCall(program.Call(), fes, funcdir, semcube); err != nil {
		return err
	}

	return nil
}

//typeCheckFunction: Verifies
func typeCheckFunction(function *ast.Function, funcdir *dir.FuncDirectory, semcube *SemanticCube) error {

	fe := funcdir.Get(function.Key())
	if fe == nil {
		return errutil.Newf("%+v: Cannot get function %s from Func Directory", function.Token(), function.Id())
	}

	rv := fe.ReturnVal()
	fes := dir.NewFuncEntryStack()
	fes.Push(fe)

	statementType, err := getTypeStatement(function.Statement(), fes, funcdir, semcube)
	if err != nil {
		return err
	}

	if !rv.Equal(statementType) {
		return errutil.Newf("%+v: Statement type does not match return type in function %s", function.Token(), function.Id())
	}

	if err := typeCheckStatement(function.Statement(), fes, funcdir, semcube); err != nil {
		return err
	}

	fes.Pop()
	return nil
}

//typeCheckStatement
func typeCheckStatement(statement ast.Statement, fes *dir.FuncEntryStack, funcdir *dir.FuncDirectory, semcube *SemanticCube) error {
	if _, ok := statement.(*ast.Id); ok {
		return nil
	} else if fcall, ok := statement.(*ast.FunctionCall); ok {
		return typeCheckFunctionCall(fcall, fes, funcdir, semcube)
	} else if lambda, ok := statement.(*ast.Lambda); ok {
		return typeCheckLambda(lambda, fes, funcdir, semcube)
	} else if cl, ok := statement.(*ast.ConstantList); ok {
		return typeCheckConstantList(cl, fes, funcdir, semcube)
	} else if _, ok := statement.(*ast.ConstantValue); ok {
		return nil
	}

	return errutil.Newf("Statement cannot be casted to any valid form")
}

//typeCheckFunctionCall
func typeCheckFunctionCall(fcall *ast.FunctionCall, fes *dir.FuncEntryStack, funcdir *dir.FuncDirectory, semcube *SemanticCube) error {
	if _, err := getTypeFunctionCall(fcall, fes, funcdir, semcube); err != nil {
		return err
	}

	for _, s := range fcall.Args() {
		if err := typeCheckStatement(s, fes, funcdir, semcube); err != nil {
			return err
		}
	}

	return nil
}

//typeCheckLambda
func typeCheckLambda(lambda *ast.Lambda, fes *dir.FuncEntryStack, funcdir *dir.FuncDirectory, semcube *SemanticCube) error {
	if lambdaEntry := fes.Top().GetLambdaEntryById(lambda.Id()); lambdaEntry != nil {
		fes.Push(lambdaEntry)

		t, err := getTypeStatement(lambda.Statement(), fes, funcdir, semcube)
		if err != nil {
			return err
		}

		if !(t.Equal(lambdaEntry.ReturnVal())) {
			return errutil.Newf("%+v: Return type of lambda statement does not match lambda definition", lambda.Token())
		}

		if err := typeCheckStatement(lambda.Statement(), fes, funcdir, semcube); err != nil {
			return err
		}

		fes.Pop()

		return nil
	}
	return errutil.Newf("%+v: Lambda could not be found in the func entry stack.", lambda.Token())
}

//typeCheckConstantList
func typeCheckConstantList(cl *ast.ConstantList, fes *dir.FuncEntryStack, funcdir *dir.FuncDirectory, semcube *SemanticCube) error {
	if _, err := getTypeConstantList(cl, fes, funcdir, semcube); err != nil {
		return err
	}

	for _, s := range cl.Contents() {
		if err := typeCheckStatement(s, fes, funcdir, semcube); err != nil {
			return err
		}
	}
	return nil
}
