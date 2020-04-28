package types

import (
	"strings"
)

// BasicType indicates the three elemental types in the Lambdish language
//
// Num: any integer or floating point number, either positive or negative
// Char: any character representable as an ascii value
// Bool: a boolean value that can only be true or false
// Null: It is used in the LambdishType struct to indicate that that type is a function
// type and thus the BasicType should not be used
type BasicType int

const (
	Num BasicType = iota
	Char
	Bool
	Null
)

func (t BasicType) convert() rune {
	if t == Num {
		return '1'
	}

	if t == Char {
		return '2'
	}

	if t == Bool {
		return '3'
	}

	return 'n'
}

// LambdishType represents any type on the lambdish language.
// A type in the language can be either a basic type (num, bool, char), or a function type.
//
// - In the case of the function type, the following should be set
//		-function: true
//		-params: non null (might be empty anyways)
//		-basic: NULL
//
// - In the case of a basic type, the following should be set
// 		-function: false
//		-params: null
//		-basic: non null, the corresponing type
//
// Additional to this, the type might represent a list. If that is the case, all rules set above will
// remain true, and list will be set to a non-zero value, indicating the levels of nesting of the type 
// in the list.
//
// For example, a value of list = 1 for a basic type num will consists of a list as following
// 		- [num]
//
// And a value of list = 3 for a function type could then look like this
// 		- [[[(num, num => bool)]]]
// 
// Nontheless external users of this package should only construct Lambdishtype structs using the
// predefined constructors provided below. When a function type is needed, the NewFuncLambdishType
// function should be called, and NewDataLambdishType with the basic type accordingly.
// This will ensure that the values are initialized correctly according to the rules set above.
type LambdishType struct {
	basic    BasicType
	retval   *LambdishType
	params   []*LambdishType
	function bool
	list     int
}

// String converts the type to its string representation which is used only in the dirfunc package
// to build the composite key of an entry
func (l LambdishType) String() string {
	var builder strings.Builder

	for i := 0; i < l.list; i++ {
		builder.WriteRune('[')
	}

	if l.function {
		builder.WriteRune('(')

		for _, t := range l.params {
			builder.WriteString(t.String())
		}

		builder.WriteString("=>")
		builder.WriteString(l.retval.String())
		builder.WriteRune(')')
	} else {
		builder.WriteRune(l.basic.convert())
	}

	for i := 0; i < l.list; i++ {
		builder.WriteRune(']')
	}

	return builder.String()
}

// List
func (lt *LambdishType) List() int {
	return lt.list
}

// Type
func (lt *LambdishType) Basic() BasicType {
	return lt.basic
}

// Params
func (lt *LambdishType) Params() []*LambdishType {
	return lt.params
}

// Function
func (lt *LambdishType) Function() bool {
	return lt.function
}

// Retval
func (lt *LambdishType) Retval() *LambdishType {
	return lt.retval
}

// List
func (lt *LambdishType) DecreaseList() {
	lt.list = lt.list-1
}

// List
func (lt *LambdishType) IncreaseList() {
	lt.list = lt.list+1
}

//Equal
func(lt *LambdishType) Equal(lt2 *LambdishType) bool {
	return lt.String() == lt2.String()
}

// NewDataLambdishType
func NewDataLambdishType(b BasicType, list int) *LambdishType {
	return &LambdishType{b, nil, nil, false, list}
}

// NewDataLambdishType
func NewFuncLambdishType(retval *LambdishType, params []*LambdishType, list int) *LambdishType {
	return &LambdishType{Null, retval, params, true, list}
}
