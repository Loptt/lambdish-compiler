package ast

import (
	"fmt"
	"github.com/Loptt/lambdish-compiler/dir"
	"github.com/Loptt/lambdish-compiler/types"
	"strings"
)

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

func (p *Program) Functions() []*Function {
	return p.functions
}
func (p *Program) Call() *FunctionCall {
	return p.call
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
	key       string
	params    []*dir.VarEntry
	t         *types.LambdishType
	statement Statement
}

func (f *Function) Id() string {
	return f.id
}

func (f *Function) Params() []*dir.VarEntry {
	return f.params
}

func (f *Function) Type() *types.LambdishType {
	return f.t
}

func (f *Function) Statement() Statement {
	return f.statement
}

func (f *Function) Key() string {
	return f.key
}

func (f *Function) CreateKey() {
	var b strings.Builder

	for _, p := range f.params {
		b.WriteString(p.Type().String())
	}

	f.key = fmt.Sprintf("%s@%s@%s", f.id, f.t, b.String())
}

// Statement interface represents the body of the function
type Statement interface {
	IsId() bool
	IsConstant() bool
	IsLambda() bool
	IsLambdaCall() bool
	IsFunctionCall() bool
}

// Id is a wrapper for a string to represent an id for a variable as a statement
type Id string

// IsId conforms to the Statement interface to determine if object is Id
func (i *Id) IsId() bool {
	return true
}

// IsConstant conforms to the Statement interface to determine if object is Constant
func (i *Id) IsConstant() bool {
	return false
}

// IsLambda conforms to the Statement interface to determine if object is Lambda
func (i *Id) IsLambda() bool {
	return false
}

// IsLambdaCall conforms to the Statement interface to determine if object is LambdaCall
func (i *Id) IsLambdaCall() bool {
	return false
}

// IsFunctionCall conforms to the Statement interface to determine if object is FunctionCall
func (i *Id) IsFunctionCall() bool {
	return false
}

// String returns the string casting of Id
func (i *Id) String() string {
	return string(*i)
}

// FunctionCall represents a call to a function either in the body of a function or as
// the main entry point of the program
//
// -id: name of the function
// -args: list of arguments to the function
//
type FunctionCall struct {
	id   string
	args []Statement
}

func (fc *FunctionCall) Args() []Statement {
	return fc.args
}

func (fc *FunctionCall) Id() string {
	return fc.id
}

// IsId conforms to the Statement interface to determine if object is Id
func (fc *FunctionCall) IsId() bool {
	return false
}

// IsConstant conforms to the Statement interface to determine if object is Constant
func (fc *FunctionCall) IsConstant() bool {
	return false
}

// IsLambda conforms to the Statement interface to determine if object is Lambda
func (fc *FunctionCall) IsLambda() bool {
	return false
}

// IsLambdaCall conforms to the Statement interface to determine if object is LambdaCall
func (fc *FunctionCall) IsLambdaCall() bool {
	return false
}

// IsFunctionCall conforms to the Statement interface to determine if object is FunctionCall
func (fc *FunctionCall) IsFunctionCall() bool {
	return true
}

// Lambda represents the definition of a lambda function without its corresponding call. This it
// should be treated as a variable
type Lambda struct {
	params    []*dir.VarEntry
	statement Statement
	retval    *types.LambdishType
}

// IsId conforms to the Statement interface to determine if object is Id
func (l *Lambda) IsId() bool {
	return false
}

// IsConstant conforms to the Statement interface to determine if object is Constant
func (l *Lambda) IsConstant() bool {
	return false
}

// IsLambda conforms to the Statement interface to determine if object is Lambda
func (l *Lambda) IsLambda() bool {
	return true
}

// IsLambdaCall conforms to the Statement interface to determine if object is LambdaCall
func (l *Lambda) IsLambdaCall() bool {
	return false
}

// IsFunctionCall conforms to the Statement interface to determine if object is FunctionCall
func (l *Lambda) IsFunctionCall() bool {
	return false
}

func (l *Lambda) Retval() *types.LambdishType {
	return l.retval
}

func (l *Lambda) Params() []*types.LambdishType {
	params := make([]*types.LambdishType, 0)

	for _, p := range l.params {
		params = append(params, p.Type())
	}
	return params
}

func (l *Lambda) Statement() Statement {
	return l.statement
}

func (l *Lambda) CreateVarDir() (*dir.VarDirectory, bool) {
	vd := dir.NewVarDirectory()

	for _, p := range l.params {
		ok := vd.Add(p)
		if !ok {
			return nil, false
		}
	}
	return vd, true
}

// Lambda call represents the definition of a lambda and subsequently calling the lamdbda function
// with the provided arguments in args
type LambdaCall struct {
	params    []*dir.VarEntry
	args      []Statement
	statement Statement
	retval    *types.LambdishType
}

func (lc *LambdaCall) Params() []*types.LambdishType {

	params := make([]*types.LambdishType, 0)

	for _, p := range lc.params {
		params = append(params, p.Type())
	}
	return params
}

func (lc *LambdaCall) Args() []Statement {
	return lc.args
}

func (lc *LambdaCall) Statement() Statement {
	return lc.statement
}

func (lc *LambdaCall) Retval() *types.LambdishType {
	return lc.retval
}

func (lc *LambdaCall) CreateVarDir() (*dir.VarDirectory, bool) {
	vd := dir.NewVarDirectory()

	for _, p := range lc.params {
		ok := vd.Add(p)
		if !ok {
			return nil, false
		}
	}
	return vd, true
}

// IsId conforms to the Statement interface to determine if object is Id
func (lc *LambdaCall) IsId() bool {
	return false
}

// IsConstant conforms to the Statement interface to determine if object is Constant
func (lc *LambdaCall) IsConstant() bool {
	return false
}

// IsLambda conforms to the Statement interface to determine if object is Lambda
func (lc *LambdaCall) IsLambda() bool {
	return false
}

// IsLambdaCall conforms to the Statement interface to determine if object is LambdaCall
func (lc *LambdaCall) IsLambdaCall() bool {
	return true
}

// IsFunctionCall conforms to the Statement interface to determine if object is FunctionCall
func (lc *LambdaCall) IsFunctionCall() bool {
	return false
}

// Constant represents a constant value which can be either a num, bool, char or a list of these
// as defined by the LambdishType struct
type Constant interface {
	Statement
	IsValue() bool
	IsList() bool
}

// ConstantValue implements the Constant interface and defines a type with a single basic value
type ConstantValue struct {
	t     *types.LambdishType
	value string
}

// IsList conforms to the constant interface to determine is the object is a list
func (cv *ConstantValue) IsValue() bool {
	return true
}

// IsList conforms to the constant interface to determine is the object is a list
func (cv *ConstantValue) IsList() bool {
	return false
}

// IsId conforms to the Statement interface to determine if object is Id
func (c *ConstantValue) IsId() bool {
	return false
}

// IsConstant conforms to the Statement interface to determine if object is Constant
func (c *ConstantValue) IsConstant() bool {
	return true
}

// IsLambda conforms to the Statement interface to determine if object is Lambda
func (c *ConstantValue) IsLambda() bool {
	return false
}

// IsLambdaCall conforms to the Statement interface to determine if object is LambdaCall
func (c *ConstantValue) IsLambdaCall() bool {
	return false
}

// IsFunctionCall conforms to the Statement interface to determine if object is FunctionCall
func (c *ConstantValue) IsFunctionCall() bool {
	return false
}

// ConstantList implements the Constant interface and defines a list which is a collection of
// statements
type ConstantList struct {
	contents []Statement
}

// IsList conforms to the constant interface to determine is the object is a list
func (cl *ConstantList) IsList() bool {
	return true
}

// IsList conforms to the constant interface to determine is the object is a list
func (cl *ConstantList) IsValue() bool {
	return false
}

// IsId conforms to the Statement interface to determine if object is Id
func (cl *ConstantList) IsId() bool {
	return false
}

// IsConstant conforms to the Statement interface to determine if object is Constant
func (cl *ConstantList) IsConstant() bool {
	return true
}

// IsLambda conforms to the Statement interface to determine if object is Lambda
func (cl *ConstantList) IsLambda() bool {
	return false
}

// IsLambdaCall conforms to the Statement interface to determine if object is LambdaCall
func (cl *ConstantList) IsLambdaCall() bool {
	return false
}

// IsFunctionCall conforms to the Statement interface to determine if object is FunctionCall
func (cl *ConstantList) IsFunctionCall() bool {
	return false
}

func (cl *ConstantList) Contents() []Statement {
	return cl.contents
}
