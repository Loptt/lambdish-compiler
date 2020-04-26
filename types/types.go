package types

import (
	"strings"
)

// BasicType indicates the three elemental types in the Lambdish language
//
// Num: any integer or floating point number, either positive or negative
// Char: any character representable as an ascii value
// Bool: a boolean value that can only be true or false
type BasicType int

const (
	Num BasicType = iota
	Char
	Bool
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

// LambdishType represents any type on the lambdish language. It consists of a basic type
// and in the case it is an array, then list will be greater than 0 indicating the amount of nested
// arrays. Otherwise 0 in list indicates that the type is just a basic type.
type LambdishType struct {
	t    BasicType
	list int
	params []*LambdishType
	function bool
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
		builder.WriteRune(l.t.convert())
		builder.WriteRune(')')
	} else {
		builder.WriteRune(l.t.convert())
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
func (lt *LambdishType) Type() BasicType {
	return lt.t
}

// Function
func (lt *LambdishType) Params() []*LambdishType {
	return lt.params
}

// Function
func (lt *LambdishType) Function() bool {
	return lt.function
}

// NewDataLambdishType
func NewDataLambdishType(t BasicType, list int) *LambdishType {
	return &LambdishType{t, list, nil, false}
}

// NewDataLambdishType
func NewFuncLambdishType(t BasicType, list int, params []*LambdishType) *LambdishType {
	return &LambdishType{t, list, params, true}
}

// Equal
func (l LambdishType) Equal(l2 LambdishType) bool {
	return l.t == l2.t && l.list == l2.list
}
