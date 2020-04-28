package sem

import (
	"github.com/Loptt/lambdish-compiler/ast"
	"github.com/Loptt/lambdish-compiler/dir"
	"github.com/Loptt/lambdish-compiler/types"
	"github.com/mewkiz/pkg/errutil"
)
//buildFuncDirProgram: Receives the program and the function directory
func buildFuncDirProgram(program *ast.Program, funcdir *dir.FuncDirectory) error {
	for _, f := range program.Functions() {
		if err := buildFuncDirFunction(f, funcdir); err != nil {
			return err
		}
	}
	return nil
}
//buildFuncDirFunction: 
func buildFuncDirFunction(function *ast.Function, funcdir *dir.FuncDirectory) error {
	id := function.Id()
	t := function.Type()
	vardir := dir.NewVarDirectory()
	params := make([]*types.LambdishType, 0)

	if idIsReserved(id) {
		return errutil.Newf("Cannot declare a function with reserved keyword %s", id)
	}

	for _, p := range function.Params() {
		params = append(params, p.Type())
		if err := buildVarDirFunction(p, vardir); err != nil {
			return err
		}
	}

	if err := checkVarDirReserved(vardir); err != nil {
		return err
	}
	
	fe := dir.NewFuncEntry(id, t, params, vardir)

	if err := buildFuncDirStatement(function.Statement(), fe); err != nil {
		return err
	}

	if ok := funcdir.Add(fe); !ok {
		return errutil.Newf("Invalid FuncEntry. This FuncEntry already exists.")
	}

	return nil
}
//buildVarDirFunction
func buildVarDirFunction(ve *dir.VarEntry, vardir *dir.VarDirectory) error {
	if ok := vardir.Add(ve); !ok {
		return errutil.Newf("Invalid VarEntry. This VarEntry already exists.")
	}
	return nil
}
//buildFuncDirStatement
func buildFuncDirStatement(statement ast.Statement, fe *dir.FuncEntry) error {
	if lambda, ok := statement.(*ast.Lambda); ok {
		vardir, ok := lambda.CreateVarDir()
		if !ok {
			return errutil.Newf("Multiple parameter declaration in lambda")
		}

		if err := checkVarDirReserved(vardir); err != nil {
			return err
		}

		lambdaEntry := fe.AddLambda(lambda.Retval(), lambda.Params(), vardir)

		// Add ID to lambda ast so that its func entry can be retreived later
		lambda.SetId(lambdaEntry.Id())

		if err := buildFuncDirStatement(lambda.Statement(), lambdaEntry); err != nil {
			return err
		}
		return nil
	} else if fcall, ok := statement.(*ast.FunctionCall); ok {
		if err := buildFuncDirStatement(fcall.Statement(), fe); err != nil {
			return err
		}
		for _, s := range fcall.Args() {
			if err := buildFuncDirStatement(s, fe); err != nil {
				return err
			}
		}
		return nil
	} else if cl, ok := statement.(*ast.ConstantList); ok {
		for _, c := range cl.Contents() {
			if err := buildFuncDirStatement(c, fe); err != nil {
				return err
			}
		}
		return nil
	} else if _, ok := statement.(*ast.ConstantValue); ok {
		return nil
	} else if _, ok := statement.(*ast.Id); ok {
		return nil
	}
	
	return errutil.Newf("Statement cannot be casted to any valid form.")
}