// Code generated by gocc; DO NOT EDIT.

package parser

import (
    "github.com/Loptt/lambdish-compiler/ast" 
    "github.com/Loptt/lambdish-compiler/dir"
    "github.com/Loptt/lambdish-compiler/types")

type (
	//TODO: change type and variable names to be consistent with other tables
	ProdTab      [numProductions]ProdTabEntry
	ProdTabEntry struct {
		String     string
		Id         string
		NTType     int
		Index      int
		NumSymbols int
		ReduceFunc func([]Attrib) (Attrib, error)
	}
	Attrib interface {
	}
)

var productionsTable = ProdTab{
	ProdTabEntry{
		String: `S' : Program	<<  >>`,
		Id:         "S'",
		NTType:     0,
		Index:      0,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Program : Functions Statement	<< ast.NewProgram(X[0], X[1]) >>`,
		Id:         "Program",
		NTType:     1,
		Index:      1,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewProgram(X[0], X[1])
		},
	},
	ProdTabEntry{
		String: `Functions : Function Functions	<< ast.AppendFunctionList(X[0], X[1]) >>`,
		Id:         "Functions",
		NTType:     2,
		Index:      2,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.AppendFunctionList(X[0], X[1])
		},
	},
	ProdTabEntry{
		String: `Functions : Function	<< ast.NewFunctionList(X[0]) >>`,
		Id:         "Functions",
		NTType:     2,
		Index:      3,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewFunctionList(X[0])
		},
	},
	ProdTabEntry{
		String: `Function : "func" id "::" Params "=>" Type "(" Statement ")"	<< ast.NewFunction(X[1], X[3], X[5], X[7]) >>`,
		Id:         "Function",
		NTType:     3,
		Index:      4,
		NumSymbols: 9,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewFunction(X[1], X[3], X[5], X[7])
		},
	},
	ProdTabEntry{
		String: `Params : Type id "," Params	<< ast.AppendParamsList(X[0], X[1], X[3]) >>`,
		Id:         "Params",
		NTType:     4,
		Index:      5,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.AppendParamsList(X[0], X[1], X[3])
		},
	},
	ProdTabEntry{
		String: `Params : Type id	<< ast.NewParamsList(X[0], X[1]) >>`,
		Id:         "Params",
		NTType:     4,
		Index:      6,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewParamsList(X[0], X[1])
		},
	},
	ProdTabEntry{
		String: `Params : empty	<< make([]*dir.VarEntry, 0), nil >>`,
		Id:         "Params",
		NTType:     4,
		Index:      7,
		NumSymbols: 0,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return make([]*dir.VarEntry, 0), nil
		},
	},
	ProdTabEntry{
		String: `Type : "num"	<< ast.NewType(X[0]) >>`,
		Id:         "Type",
		NTType:     5,
		Index:      8,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewType(X[0])
		},
	},
	ProdTabEntry{
		String: `Type : "bool"	<< ast.NewType(X[0]) >>`,
		Id:         "Type",
		NTType:     5,
		Index:      9,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewType(X[0])
		},
	},
	ProdTabEntry{
		String: `Type : "char"	<< ast.NewType(X[0]) >>`,
		Id:         "Type",
		NTType:     5,
		Index:      10,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewType(X[0])
		},
	},
	ProdTabEntry{
		String: `Type : "(" FuncTypes "=>" Type ")"	<< ast.NewFunctionType(X[1], X[3]) >>`,
		Id:         "Type",
		NTType:     5,
		Index:      11,
		NumSymbols: 5,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewFunctionType(X[1], X[3])
		},
	},
	ProdTabEntry{
		String: `Type : "[" Type "]"	<< ast.AppendType(X[1]) >>`,
		Id:         "Type",
		NTType:     5,
		Index:      12,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.AppendType(X[1])
		},
	},
	ProdTabEntry{
		String: `FuncTypes : Type "," FuncTypes	<< ast.AppendFuncTypeList(X[0], X[2]) >>`,
		Id:         "FuncTypes",
		NTType:     6,
		Index:      13,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.AppendFuncTypeList(X[0], X[2])
		},
	},
	ProdTabEntry{
		String: `FuncTypes : Type	<< ast.NewFuncTypeList(X[0]) >>`,
		Id:         "FuncTypes",
		NTType:     6,
		Index:      14,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewFuncTypeList(X[0])
		},
	},
	ProdTabEntry{
		String: `FuncTypes : empty	<< make([]*types.LambdishType, 0 ), nil >>`,
		Id:         "FuncTypes",
		NTType:     6,
		Index:      15,
		NumSymbols: 0,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return make([]*types.LambdishType, 0 ), nil
		},
	},
	ProdTabEntry{
		String: `Statement : id	<< ast.NewStatement(X[0]) >>`,
		Id:         "Statement",
		NTType:     7,
		Index:      16,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewStatement(X[0])
		},
	},
	ProdTabEntry{
		String: `Statement : Constant	<< ast.NewStatement(X[0]) >>`,
		Id:         "Statement",
		NTType:     7,
		Index:      17,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewStatement(X[0])
		},
	},
	ProdTabEntry{
		String: `Statement : LambdaExpr	<< ast.NewStatement(X[0]) >>`,
		Id:         "Statement",
		NTType:     7,
		Index:      18,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewStatement(X[0])
		},
	},
	ProdTabEntry{
		String: `Statement : FunctionCall	<< ast.NewStatement(X[0]) >>`,
		Id:         "Statement",
		NTType:     7,
		Index:      19,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewStatement(X[0])
		},
	},
	ProdTabEntry{
		String: `FunctionCall : id "(" Args ")"	<< ast.NewFunctionCall(X[0], X[2]) >>`,
		Id:         "FunctionCall",
		NTType:     8,
		Index:      20,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewFunctionCall(X[0], X[2])
		},
	},
	ProdTabEntry{
		String: `FunctionCall : operations "(" Args ")"	<< ast.NewFunctionCall(X[0], X[2]) >>`,
		Id:         "FunctionCall",
		NTType:     8,
		Index:      21,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewFunctionCall(X[0], X[2])
		},
	},
	ProdTabEntry{
		String: `FunctionCall : relop "(" Args ")"	<< ast.NewFunctionCall(X[0], X[2]) >>`,
		Id:         "FunctionCall",
		NTType:     8,
		Index:      22,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewFunctionCall(X[0], X[2])
		},
	},
	ProdTabEntry{
		String: `LambdaExpr : Lambda	<< X[0], nil >>`,
		Id:         "LambdaExpr",
		NTType:     9,
		Index:      23,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `LambdaExpr : LambdaCall	<< X[0], nil >>`,
		Id:         "LambdaExpr",
		NTType:     9,
		Index:      24,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Lambda : "(" "#" Params "=>" "(" Statement ")" ")"	<< ast.NewLambda(X[2],X[5]) >>`,
		Id:         "Lambda",
		NTType:     10,
		Index:      25,
		NumSymbols: 8,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewLambda(X[2],X[5])
		},
	},
	ProdTabEntry{
		String: `LambdaCall : "(" "#" Params "=>" "(" Statement ")" ")" "(" Args ")"	<< ast.NewLambdaCall(X[2],X[5],X[9]) >>`,
		Id:         "LambdaCall",
		NTType:     11,
		Index:      26,
		NumSymbols: 11,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewLambdaCall(X[2],X[5],X[9])
		},
	},
	ProdTabEntry{
		String: `Args : Statement "," Args	<< ast.AppendStatementList(X[0],X[2]) >>`,
		Id:         "Args",
		NTType:     12,
		Index:      27,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.AppendStatementList(X[0],X[2])
		},
	},
	ProdTabEntry{
		String: `Args : Statement	<< ast.NewStatementList(X[0]) >>`,
		Id:         "Args",
		NTType:     12,
		Index:      28,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewStatementList(X[0])
		},
	},
	ProdTabEntry{
		String: `Args : empty	<< make([]ast.Statement, 0), nil >>`,
		Id:         "Args",
		NTType:     12,
		Index:      29,
		NumSymbols: 0,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return make([]ast.Statement, 0), nil
		},
	},
	ProdTabEntry{
		String: `Constant : boolean	<< ast.NewConstantBool(X[0]) >>`,
		Id:         "Constant",
		NTType:     13,
		Index:      30,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewConstantBool(X[0])
		},
	},
	ProdTabEntry{
		String: `Constant : number	<< ast.NewConstantNum(X[0]) >>`,
		Id:         "Constant",
		NTType:     13,
		Index:      31,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewConstantNum(X[0])
		},
	},
	ProdTabEntry{
		String: `Constant : charac	<< ast.NewConstantChar(X[0]) >>`,
		Id:         "Constant",
		NTType:     13,
		Index:      32,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewConstantChar(X[0])
		},
	},
	ProdTabEntry{
		String: `Constant : "[" Args "]"	<< ast.AppendConstant(X[1]) >>`,
		Id:         "Constant",
		NTType:     13,
		Index:      33,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.AppendConstant(X[1])
		},
	},
}
