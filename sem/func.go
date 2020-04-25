package sem

import (
	"github.com/Loptt/lambdish-compiler/ast"
	"github.com/Loptt/lambdish-compiler/dir"
	"github.com/Loptt/lambdish-compiler/types"
	"github.com/mewkiz/pkg/errutil"
)

func buildFuncDirProgram(program *ast.Program, funcdir *dir.FuncDirectory) error {
	for _, f := range program.Functions() {
		err := buildFuncDirFunction(f, funcdir)
		if err != nil {
			return err
		}
	}
	return nil
}

func buildFuncDirFunction(function *ast.Function, funcdir *dir.FuncDirectory) error {
	id := function.Id()
	t := function.Type()
	vardir := dir.NewVarDirectory()
	params := make([]*types.LambdishType, 0)

	for _, p := range function.Params() {
		params = append(params, p.Type())
		err := buildVarDirFunction(p, vardir)
		if err != nil {
			return err
		}
	}
	
	fe := dir.NewFuncEntry(id, t, params, vardir)

	err := buildFuncDirStatement(function.Statement(), fe)
	if err != nil {
		return err
	}

	ok := funcdir.Add(fe)
	if !ok {
		return errutil.Newf("Invalid FuncEntry. This FuncEntry already exists.")
	}
	return nil
}

func buildVarDirFunction(ve *dir.VarEntry, vardir *dir.VarDirectory) error {
	ok := vardir.Add(ve)
	if !ok {
		return errutil.Newf("Invalid VarEntry. This VarEntry already exists.")
	}
	return nil
}

func buildFuncDirStatement(statement ast.Statement, fe *dir.FuncEntry) error {
	if lambda, ok := statement.(*ast.Lambda); ok {
		lamdbaEntry := fe.AddLambda(lambda.Retval(), lambda.Params(), lambda.VarDir())
		err := buildFuncDirStatement(lambda.Statement(), lamdbaEntry)
		if err != nil {
			return err
		} 
		return nil
	} else if lcall, ok := statement.(*ast.LambdaCall); ok {
		lamdbaEntry := fe.AddLambda(lcall.Retval(), lcall.Params(), lcall.VarDir())
		err := buildFuncDirStatement(lcall.Statement(), lamdbaEntry)
		if err != nil {
			return err
		}
		return nil
	} else if fcall, ok := statement.(*ast.FunctionCall); ok {
		for _, s := range fcall.Args() {
			err := buildFuncDirStatement(s, fe)
			if err != nil {
				return err
			}
		}
		return nil
	} else if cl, ok := statement.(*ast.ConstantList); ok {
		for _, c := range cl.Contents() {
			err := buildFuncDirStatement(c, fe)
			if err != nil {
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