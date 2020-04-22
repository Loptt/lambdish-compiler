package ast

import (
	"github.com/Loptt/lambdish-compiler/dir"
	"github.com/Loptt/lambdish-compiler/types"
	"github.com/mewkiz/pkg/errutil"
)


func NewProgram(function, call interface{}) (*Program, error) {
	f, ok := function.([]*Function)
	if !ok ¨{
		return nil, errutil.Newf("Invalid type for functions. Expected []*Function")
	}

	c, ok := call.(*FunctionCall)
	if !ok {
		return nil, errutil.Newf("Invalid type for function call. Expected *FunctionCall")
	}

	return &Program{f, c}, nil
} 
