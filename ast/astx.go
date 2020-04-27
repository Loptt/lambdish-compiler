package ast

import (
	"github.com/Loptt/lambdish-compiler/dir"
	"github.com/Loptt/lambdish-compiler/gocc/token"
	"github.com/Loptt/lambdish-compiler/types"
	"github.com/mewkiz/pkg/errutil"
)

// NewProgram
func NewProgram(functions, call interface{}) (*Program, error) {
	fs, ok := functions.([]*Function)
	if !ok {
		return nil, errutil.Newf("Invalid type for functions. Expected []*Function")
	}

	c, ok := call.(*FunctionCall)
	if !ok {
		return nil, errutil.Newf("Invalid type for function call. Expected *FunctionCall")
	}

	return &Program{fs, c}, nil
}

// NewFunctionList
func NewFunctionList(function interface{}) ([]*Function, error) {
	f, ok := function.(*Function)
	if !ok {
		return nil, errutil.Newf("Invalid type for function. Expected *Function")
	}

	return []*Function{f}, nil
}

// AppendFunctionList
func AppendFunctionList(function, list interface{}) ([]*Function, error) {
	f, ok := function.(*Function)
	if !ok {
		return nil, errutil.Newf("Invalid type for function. Expected *Function")
	}

	flist, ok := list.([]*Function)
	if !ok {
		return nil, errutil.Newf("Invalid type for functions. Expected []*Function")
	}

	return append([]*Function{f}, flist...), nil
}

// NewFunction
func NewFunction(id, params, typ, statement interface{}) (*Function, error) {
	i, ok := id.(*token.Token)
	if !ok {
		return nil, errutil.Newf("Invalid type for id. Expected token")
	}

	d := string(i.Lit)

	p, ok := params.([]*dir.VarEntry)
	if !ok {
		return nil, errutil.Newf("Invalid type for params. Expected []*dir.VarEntry")
	}

	t, ok := typ.(*types.LambdishType)
	if !ok {
		return nil, errutil.Newf("Invalid type for typ. Expected *types.LambdishType")
	}

	s, ok := statement.(Statement)
	if !ok {
		return nil, errutil.Newf("Invalid type for statement. Expected Statement")
	}

	f := &Function{d, "", p, t, s}
	f.CreateKey()

	return f, nil
}

// NewStatementList
func NewStatementList(statement interface{}) ([]Statement, error) {
	s, ok := statement.(Statement)
	if !ok {
		return nil, errutil.Newf("NewStatementList: Invalid type for statement. Expected Statement interface")
	}

	return []Statement{s}, nil
}

// AppendStatementList
func AppendStatementList(statement, list interface{}) ([]Statement, error) {
	l, ok := list.([]Statement)
	if !ok {
		return nil, errutil.Newf("Invalid type for statement list. Expected []Statement")
	}

	// Check if the value is an id and cast it fist to a token
	if s, ok := statement.(*token.Token); ok {
		id :=  Id(s.Lit)
		return append([]Statement{&id}, l...), nil
	// If not, cast the value to a statement interface
	} else if s, ok := statement.(Statement); ok {
		return append([]Statement{s}, l...), nil
	}

	return nil, errutil.Newf("Invalid type for statement. Expected Statement interface got %v", statement)
}

// NewStatement
func NewStatement(value interface{}) (Statement, error) {
	// Check if the value is an id and cast it fist to a token
	if t, ok := value.(*token.Token); ok {
		id :=  Id(t.Lit)
		return &id, nil
	// If not, cast the value to a statement interface
	} else if s, ok := value.(Statement); ok {
		return s, nil
	}

	return nil, errutil.Newf("Invalid type for statement. Expected Statement interface got %v", value)
}

// AppendParamsList
func AppendParamsList(typ, id, list interface{}) ([]*dir.VarEntry, error) {
	i, ok := id.(*token.Token)
	if !ok {
		return nil, errutil.Newf("Invalid type for id. Expected token")
	}

	d := string(i.Lit)

	t, ok := typ.(*types.LambdishType)
	if !ok {
		return nil, errutil.Newf("Invalid type for typ. Expected *types.LambdishType")
	}

	v := dir.NewVarEntry(d, t)

	vlist, ok := list.([]*dir.VarEntry)
	if !ok {
		return nil, errutil.Newf("Invalid type for parameters. Expected []*dir.VarEntry")
	}

	return append([]*dir.VarEntry{v}, vlist...), nil
}

// NewParamsList
func NewParamsList(typ, id interface{}) ([]*dir.VarEntry, error) {

	i, ok := id.(*token.Token)
	if !ok {
		return nil, errutil.Newf("Invalid type for id. Expected token")
	}

	d := string(i.Lit)

	t, ok := typ.(*types.LambdishType)
	if !ok {
		return nil, errutil.Newf("Invalid type for typ. Expected *types.LambdishType")
	}

	v := dir.NewVarEntry(d, t)

	return []*dir.VarEntry{v}, nil
}

// NewType
func NewType(t interface{}) (*types.LambdishType, error) {
	typ, ok := t.(*token.Token)
	if !ok {
		return nil, errutil.Newf("Invalid type for id. Expected token")
	}

	tstring := string(typ.Lit)

	if !ok {
		return nil, errutil.Newf("Invalid type for type. Expected string, got %v", t)
	}

	if tstring == "num" {
		return types.NewDataLambdishType(types.Num, 0), nil
	}
	if tstring == "bool" {
		return types.NewDataLambdishType(types.Bool, 0), nil
	}
	if tstring == "char" {
		return types.NewDataLambdishType(types.Char, 0), nil
	}

	return nil, errutil.Newf("Invalid type for type. Expected BasicType enum")
}

// NewType
func NewFunctionType(params, ret interface{}) (*types.LambdishType, error) {
	p, ok := params.([]*types.LambdishType)
	if !ok {
		return nil, errutil.Newf("Invalid type for params. Expected []LambdishType")
	}

	rv, err := ret.(*types.LambdishType)

	if !err {
		return nil, errutil.Newf("Invalid type for type. Expected LambdishType.")
	}

	return types.NewFuncLambdishType(rv, p, 0), nil
}

// AppendType
func AppendType(typ interface{}) (*types.LambdishType, error) {
	t, ok := typ.(*types.LambdishType)
	if !ok {
		return nil, errutil.Newf("Invalid type for typ. Expected *types.LambdishType")
	}

	if t.Function() {
		return types.NewFuncLambdishType(t.Retval(), t.Params(), t.List()+1), nil
	} else {
		return types.NewDataLambdishType(t.Basic(), t.List()+1), nil
	}
}

// NewFuncType
func NewFuncTypeList(typ interface{}) ([]*types.LambdishType, error) {
	t, ok := typ.(*types.LambdishType)
	if !ok {
		return nil, errutil.Newf("Invalid type for typ. Expected *types.LambdishType")
	}

	return []*types.LambdishType{t}, nil
}

// NewFuncType
func AppendFuncTypeList(typ, list interface{}) ([]*types.LambdishType, error) {
	t, ok := typ.(*types.LambdishType)
	if !ok {
		return nil, errutil.Newf("Invalid type for typ. Expected *types.LambdishType")
	}

	l, ok := list.([]*types.LambdishType)
	if !ok {
		return nil, errutil.Newf("Invalid type for typ. Expected []*types.LambdishType")
	}

	return append([]*types.LambdishType{t}, l...), nil
}

// NewFunctionCall
func NewFunctionCall(id, args interface{}) (*FunctionCall, error) {
	i, ok := id.(Statement)
	if !ok {
		return nil, errutil.Newf("Invalid type for id. Expected statement")
	}
	
	a, ok := args.([]Statement)
	if !ok {
		return nil, errutil.Newf("Invalid type for args. Expected []Statement, got %v", args)
	}

	return &FunctionCall{i, a}, nil
}

// NewLambda
func NewLambda(params, retval, statement interface{}) (*Lambda, error) {
	p, ok := params.([]*dir.VarEntry)
	if !ok {
		return nil, errutil.Newf("Invalid type for params. Expected []*dir.VarEntry]")
	}

	s, ok := statement.(Statement)
	if !ok {
		return nil, errutil.Newf("Invalid type for params. Expected Statement")
	}

	t,ok := retval.(*types.LambdishType)
	if !ok {
		return nil, errutil.Newf("Invalid type for type. Expected *types.LambdishType")
	}


	return &Lambda{p, s, t}, nil
}

// NewConstantBool
func NewConstantBool(value interface{}) (Constant, error) {
	val, ok := value.(*token.Token)
	if !ok {
		return nil, errutil.Newf("Invalid type for id. Expected token")
	}

	v := string(val.Lit)

	return &ConstantValue{types.NewDataLambdishType(types.Bool, 0), v}, nil
}

// NewConstantNum
func NewConstantNum(value interface{}) (Constant, error) {
	val, ok := value.(*token.Token)
	if !ok {
		return nil, errutil.Newf("Invalid type for id. Expected token")
	}

	v := string(val.Lit)

	return &ConstantValue{types.NewDataLambdishType(types.Num, 0), v}, nil
}

// NewConstantChar
func NewConstantChar(value interface{}) (Constant, error) {
	val, ok := value.(*token.Token)
	if !ok {
		return nil, errutil.Newf("Invalid type for id. Expected token")
	}

	v := string(val.Lit)

	return &ConstantValue{types.NewDataLambdishType(types.Char, 0), v}, nil
}

// AppendConstant
func AppendConstant(list interface{}) (*ConstantList, error) {

	l, ok := list.([]Statement)
	if !ok {
		return nil, errutil.Newf("Invalid type for statement list. Expected []Statement")
	}

	return &ConstantList{l}, nil
}
