package ast

import (
	"fmt"

	"github.com/Loptt/lambdish-compiler/dir"
	"github.com/Loptt/lambdish-compiler/gocc/token"
	"github.com/Loptt/lambdish-compiler/types"
	"github.com/mewkiz/pkg/errutil"
)

// NewProgram creates a new Program node which acts as the root of the tree
func NewProgram(functions, call interface{}) (*Program, error) {
	fs, ok := functions.([]*Function)
	if !ok {
		return nil, errutil.NewNoPosf("Invalid type for functions. Expected []*Function")
	}

	c, ok := call.(*FunctionCall)
	if !ok {
		return nil, errutil.NewNoPosf("Invalid type for function call. Expected *FunctionCall")
	}

	return &Program{fs, c}, nil
}

// NewFunctionList
func NewFunctionList(function interface{}) ([]*Function, error) {
	f, ok := function.(*Function)
	if !ok {
		return nil, errutil.NewNoPosf("Invalid type for function. Expected *Function")
	}

	return []*Function{f}, nil
}

// AppendFunctionList
func AppendFunctionList(function, list interface{}) ([]*Function, error) {
	f, ok := function.(*Function)
	if !ok {
		return nil, errutil.NewNoPosf("Invalid type for function. Expected *Function")
	}

	flist, ok := list.([]*Function)
	if !ok {
		return nil, errutil.NewNoPosf("Invalid type for functions. Expected []*Function")
	}

	return append([]*Function{f}, flist...), nil
}

// NewFunction creates a new Function node which acts as the children of the tree
func NewFunction(id, params, typ, statement interface{}) (*Function, error) {
	i, ok := id.(*token.Token)
	if !ok {
		return nil, errutil.NewNoPosf("Invalid type for id. Expected token")
	}

	d := string(i.Lit)

	p, ok := params.([]*dir.VarEntry)
	if !ok {
		return nil, errutil.NewNoPosf("Invalid type for params. Expected []*dir.VarEntry")
	}

	t, ok := typ.(*types.LambdishType)
	if !ok {
		return nil, errutil.NewNoPosf("Invalid type for typ. Expected *types.LambdishType")
	}

	s, ok := statement.(Statement)
	if !ok {
		return nil, errutil.NewNoPosf("Invalid type for statement. Expected Statement")
	}

	f := &Function{d, "", p, t, s, i}
	f.CreateKey()

	return f, nil
}

// NewStatementList
func NewStatementList(statement interface{}) ([]Statement, error) {
	s, ok := statement.(Statement)
	if !ok {
		return nil, errutil.NewNoPosf("NewStatementList: Invalid type for statement. Expected Statement interface")
	}

	return []Statement{s}, nil
}

// AppendStatementList
func AppendStatementList(statement, list interface{}) ([]Statement, error) {
	l, ok := list.([]Statement)
	if !ok {
		return nil, errutil.NewNoPosf("Invalid type for statement list. Expected []Statement")
	}

	// Check if the value is an id and cast it fist to a token
	if s, ok := statement.(*token.Token); ok {
		id := &Id{string(s.Lit), s}
		return append([]Statement{id}, l...), nil
		// If not, cast the value to a statement interface
	} else if s, ok := statement.(Statement); ok {
		return append([]Statement{s}, l...), nil
	}

	return nil, errutil.NewNoPosf("Invalid type for statement. Expected Statement interface got %v", statement)
}

// NewStatement creates a new Statement node which acts as the children of the function, which is the body of the function
func NewStatement(value interface{}) (Statement, error) {
	// Check if the value is an id and cast it fist to a token
	if t, ok := value.(*token.Token); ok {
		return &Id{string(t.Lit), t}, nil
		// If not, cast the value to a statement interface
	} else if s, ok := value.(Statement); ok {
		return s, nil
	}

	return nil, errutil.NewNoPosf("Invalid type for statement. Expected Statement interface got %v", value)
}

// AppendParamsList
func AppendParamsList(typ, id, list interface{}) ([]*dir.VarEntry, error) {
	i, ok := id.(*token.Token)
	if !ok {
		return nil, errutil.NewNoPosf("Invalid type for id. Expected token")
	}

	d := string(i.Lit)

	t, ok := typ.(*types.LambdishType)
	if !ok {
		return nil, errutil.NewNoPosf("Invalid type for typ. Expected *types.LambdishType")
	}

	vlist, ok := list.([]*dir.VarEntry)
	if !ok {
		return nil, errutil.NewNoPosf("Invalid type for parameters. Expected []*dir.VarEntry")
	}

	v := dir.NewVarEntry(d, t, i, len(vlist))

	return append([]*dir.VarEntry{v}, vlist...), nil
}

// NewParamsList
func NewParamsList(typ, id interface{}) ([]*dir.VarEntry, error) {

	i, ok := id.(*token.Token)
	if !ok {
		return nil, errutil.NewNoPosf("Invalid type for id. Expected token")
	}

	d := string(i.Lit)

	t, ok := typ.(*types.LambdishType)
	if !ok {
		return nil, errutil.NewNoPosf("Invalid type for typ. Expected *types.LambdishType")
	}

	v := dir.NewVarEntry(d, t, i, 0)

	return []*dir.VarEntry{v}, nil
}

// NewType
func NewType(t interface{}) (*types.LambdishType, error) {
	typ, ok := t.(*token.Token)
	if !ok {
		return nil, errutil.NewNoPosf("Invalid type for id. Expected token")
	}

	tstring := string(typ.Lit)

	if !ok {
		return nil, errutil.NewNoPosf("Invalid type for type. Expected string, got %v", t)
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

	return nil, errutil.NewNoPosf("Invalid type for type. Expected BasicType enum")
}

// NewType
func NewFunctionType(params, ret interface{}) (*types.LambdishType, error) {
	p, ok := params.([]*types.LambdishType)
	if !ok {
		return nil, errutil.NewNoPosf("Invalid type for params. Expected []LambdishType")
	}

	rv, err := ret.(*types.LambdishType)

	if !err {
		return nil, errutil.NewNoPosf("Invalid type for type. Expected LambdishType.")
	}

	return types.NewFuncLambdishType(rv, p, 0), nil
}

// AppendType
func AppendType(typ interface{}) (*types.LambdishType, error) {
	t, ok := typ.(*types.LambdishType)
	if !ok {
		return nil, errutil.NewNoPosf("Invalid type for typ. Expected *types.LambdishType")
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
		return nil, errutil.NewNoPosf("Invalid type for typ. Expected *types.LambdishType")
	}

	return []*types.LambdishType{t}, nil
}

// NewFuncType
func AppendFuncTypeList(typ, list interface{}) ([]*types.LambdishType, error) {
	t, ok := typ.(*types.LambdishType)
	if !ok {
		return nil, errutil.NewNoPosf("Invalid type for typ. Expected *types.LambdishType")
	}

	l, ok := list.([]*types.LambdishType)
	if !ok {
		return nil, errutil.NewNoPosf("Invalid type for typ. Expected []*types.LambdishType")
	}

	return append([]*types.LambdishType{t}, l...), nil
}

// NewFunctionCall creates a new FunctionCall node which acts as the children of the program or function, which is the function call
func NewFunctionCall(id, args interface{}) (*FunctionCall, error) {
	i, ok := id.(Statement)
	if !ok {
		return nil, errutil.NewNoPosf("Invalid type for id. Expected statement")
	}

	a, ok := args.([]Statement)
	if !ok {
		return nil, errutil.NewNoPosf("Invalid type for args. Expected []Statement, got %v", args)
	}

	return &FunctionCall{i, a}, nil
}

func NewFunctionReservedCall(id, args interface{}) (*FunctionCall, error) {
	i, ok := id.(*token.Token)
	if !ok {
		return nil, errutil.NewNoPosf("Invalid type for id. Expected statement")
	}

	v := string(i.Lit)
	idstruct := Id{v, i}

	a, ok := args.([]Statement)
	if !ok {
		return nil, errutil.NewNoPosf("Invalid type for args. Expected []Statement, got %v", args)
	}

	return &FunctionCall{&idstruct, a}, nil
}

// NewLambda creates a new NewLambda node which acts as a child of a Function Node, which is the lambda declaration
func NewLambda(tok, params, retval, statement interface{}) (*Lambda, error) {
	tk, ok := tok.(*token.Token)
	if !ok {
		return nil, errutil.NewNoPosf("Invalid type for paratkms. Expected *token.Token")
	}

	p, ok := params.([]*dir.VarEntry)
	if !ok {
		return nil, errutil.NewNoPosf("Invalid type for params. Expected []*dir.VarEntry]")
	}

	s, ok := statement.(Statement)
	if !ok {
		return nil, errutil.NewNoPosf("Invalid type for params. Expected Statement")
	}

	t, ok := retval.(*types.LambdishType)
	if !ok {
		return nil, errutil.NewNoPosf("Invalid type for type. Expected *types.LambdishType")
	}

	return &Lambda{p, s, t, "", tk}, nil
}

// NewConstantBool
func NewConstantBool(value interface{}) (Constant, error) {
	val, ok := value.(*token.Token)
	if !ok {
		return nil, errutil.NewNoPosf("Invalid type for id. Expected token")
	}

	v := string(val.Lit)

	return &ConstantValue{types.NewDataLambdishType(types.Bool, 0), v, val}, nil
}

// NewConstantNum
func NewConstantNum(value interface{}) (Constant, error) {
	val, ok := value.(*token.Token)
	if !ok {
		return nil, errutil.NewNoPosf("Invalid type for id. Expected token")
	}

	v := string(val.Lit)

	return &ConstantValue{types.NewDataLambdishType(types.Num, 0), v, val}, nil
}

// NewConstantChar
func NewConstantChar(value interface{}) (Constant, error) {
	val, ok := value.(*token.Token)
	if !ok {
		return nil, errutil.NewNoPosf("Invalid type for id. Expected token")
	}

	v := string(val.Lit)

	return &ConstantValue{types.NewDataLambdishType(types.Char, 0), v, val}, nil
}

// AppendConstant
func AppendConstant(start, list interface{}) (*ConstantList, error) {
	val, ok := start.(*token.Token)
	if !ok {
		return nil, errutil.NewNoPosf("Invalid type for start. Expected token")
	}
	l, ok := list.([]Statement)
	if !ok {
		return nil, errutil.NewNoPosf("Invalid type for statement list. Expected []Statement")
	}

	return &ConstantList{l, val, nil}, nil
}

// AppendEmptyConstant creates an empty list with a given type
func AppendEmptyConstant(start, t interface{}) (*ConstantList, error) {
	val, ok := start.(*token.Token)
	if !ok {
		return nil, errutil.NewNoPosf("Invalid type for start. Expected token")
	}

	typ, ok := t.(*types.LambdishType)
	if !ok {
		return nil, errutil.NewNoPosf("Invalid type for typ. Expected *types.LambdishType")
	}

	// We increase the list content because the read type is always one list level lower
	// than the actual list
	typ.IncreaseList()

	return &ConstantList{make([]Statement, 0), val, typ}, nil
}

func AppendStringConstant(str interface{}) (*ConstantList, error) {
	val, ok := str.(*token.Token)
	if !ok {
		return nil, errutil.NewNoPosf("Invalid type for start. Expected token")
	}

	substr := val.Lit[1 : len(val.Lit)-1]

	chars := make([]Statement, 0)

	for _, c := range []byte(substr) {
		chars = append(chars, &ConstantValue{types.NewDataLambdishType(types.Char, 0), fmt.Sprintf("'%c'", c), val})
	}

	return &ConstantList{[]Statement(chars), val, nil}, nil
}
