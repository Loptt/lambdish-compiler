package ast

import (
	"github.com/Loptt/lambdish-compiler/dir"
	"github.com/Loptt/lambdish-compiler/types"
)

// Node represents a single node in the ast. Currently node has no specific methods to
// implement thus, this can be any type
type Node interface {
}

// Program defines the root of the tree
// This node consists of
//
// -List of Function
// -FunctionCall
//
type Program struct {
	functions []*Function
	call      *FunctionCall
}

// Function represents a function definition and consists of
//
// -id: name of the function
// -params: list of parameters
// -t: return value of the funcion
// -statement: body of the function
//
type Function struct {
	id        string
	params    []*dir.VarEntry
	t         *types.LambdishType
	statement *Statement
}

// Statement interface represents the body of the function
type Statement interface {
}

// Args represents a list of arguments to call a function. 
type Args struct {
	statements []*Statement
}

// FunctionCall represents a call to a function either in the body of a function or as
// the main entry point of the program
//
// -id: name of the function
// -args: list of arguments to the function
//
type FunctionCall struct {
	id   string
	args *Args
}

// LambdaExpr represents either a lambda function call or a lambda function definition without
// the call. If args is nil. Then the expression is just a function definition without its
// corresponding call. If args is present. It should be treated as a call to a function
//
// -params: list of parameters
// - args: list of arguments to call the function
//
type LambdaExpr struct {
	params []*dir.VarEntry
	args   *Args
}

// Constant represents a constant value which can be either a num, bool, char or a list of these
// as defined by the LambdishType struct
type Constant struct {
	t *types.LambdishType
	value string
}
