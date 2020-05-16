package ast

import (
	"github.com/Loptt/lambdish-compiler/dir"
	"github.com/Loptt/lambdish-compiler/gocc/token"
	"github.com/Loptt/lambdish-compiler/types"
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
	tok       *token.Token
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
	f.key = f.id
}

func (f *Function) Token() *token.Token {
	return f.tok
}

// Statement interface represents the body of the function
type Statement interface {
	IsId() bool
	IsConstant() bool
	IsLambda() bool
	IsFunctionCall() bool
	Token() *token.Token
}

// TODO: Remove IsLambdaCall from interface statement

// Id is a wrapper for a string to represent an id for a variable as a statement
type Id struct {
	id  string
	tok *token.Token
}

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

// IsFunctionCall conforms to the Statement interface to determine if object is FunctionCall
func (i *Id) IsFunctionCall() bool {
	return false
}

// String returns the string casting of Id
func (i *Id) String() string {
	return i.id
}

func (i *Id) Token() *token.Token {
	return i.tok
}

// FunctionCall represents a call to a function either in the body of a function or as
// the main entry point of the program
//
// -id: name of the function
// -args: list of arguments to the function
//
type FunctionCall struct {
	s    Statement
	args []Statement
}

func (fc *FunctionCall) Args() []Statement {
	return fc.args
}

func (fc *FunctionCall) Statement() Statement {
	return fc.s
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

func (fc *FunctionCall) Token() *token.Token {
	return fc.s.Token()
}

// Lambda represents the definition of a lambda function without its corresponding call. This it
// should be treated as a variable
type Lambda struct {
	params    []*dir.VarEntry
	statement Statement
	retval    *types.LambdishType
	id        string
	tok       *token.Token
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

func (l *Lambda) Id() string {
	return l.id
}

func (l *Lambda) Token() *token.Token {
	return l.tok
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

func (l *Lambda) SetId(id string) {
	l.id = id
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
	tok   *token.Token
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

func (c *ConstantValue) Type() *types.LambdishType {
	return c.t
}

func (c *ConstantValue) Token() *token.Token {
	return c.tok
}

func (c *ConstantValue) Value() string {
	return c.value
}

// ConstantList implements the Constant interface and defines a list which is a collection of
// statements
type ConstantList struct {
	contents []Statement
	tok      *token.Token
	t        *types.LambdishType
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

func (cl *ConstantList) Token() *token.Token {
	return cl.tok
}

func (cl *ConstantList) Type() *types.LambdishType {
	return cl.t
}
