package ast

import (
	"github.com/Loptt/lambdish-compiler/dir"
	"github.com/Loptt/lambdish-compiler/types"
	"github.com/mewkiz/pkg/errutil"
)

func NewProgram(functions, call interface{}) (*Program, error) {
	fs, ok := functions.([]*Function)
	if !ok ¨{
		return nil, errutil.Newf("Invalid type for functions. Expected []*Function")
	}

	c, ok := call.(*FunctionCall)
	if !ok {
		return nil, errutil.Newf("Invalid type for function call. Expected *FunctionCall")
	}

	return &Program{fs, c}, nil
} 

func NewFunctionList(function interface{}) ([]*Function, error) {
	f, ok := function.(*Function)
	if !ok ¨{
		return nil, errutil.Newf("Invalid type for function. Expected *Function")
	}

	return []*Function{f}
}

func AppendFunctionList(function, list interface{}), ([]*Function, error) {
	f, ok := function.(*Function)
	if !ok ¨{
		return nil, errutil.Newf("Invalid type for function. Expected *Function")
	}

	flist, ok := list.([]*Function)
	if !ok ¨{
		return nil, errutil.Newf("Invalid type for functions. Expected []*Function")
	}

	return append(flist, f)
}

func NewFunction(id, params, typ, statement interface{}) (*Function, error) {
	d, ok := id.(string)
	if !ok {
		return nil, errutil.Newf("Invalid type for id. Expected string")
	}

	p, ok := params.([]*dir.VarEntry)
	if !ok {
		return nil, errutil.Newf("Invalid type for params. Expected []*dir.VarEntry")
	}

	t, ok := typ.(*types.LambdishType)
	if !ok {
		return nil, errutil.Newf("Invalid type for id. Expected *types.LambdishType")
	}

	// Statement can have one of the following forms:
	// 		-string when reffering to the id of a variable
	// 		-*Constant when refering to a constant value or list
	// 		-LambdaExpr when representing a lambda function
	// 		-FunctionCall when calling a function

	if s, ok := statement.(string); ok {
		return &Function{d, p, t, s}
	}
	else if s, ok := statement.(*Constant); ok {
		return &Function{d, p, t, s}
	}
	else if s, ok := statement.(*LambdaExpr); ok {
		return &Function{d, p, t, s}
	}
	else if s, ok := statement.(*FunctionCall); ok {
		return &Function{d, p, t, s}
	}

	return nil, errutil.Newf("Invalid type for statement. Expected string, *Constant, *LambdaExpr, or *FunctionCall")
}