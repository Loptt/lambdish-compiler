package sem

import (
	"github.com/Loptt/lambdish-compiler/types"
	"github.com/Loptt/lambdish-compiler/dir"
	"github.com/Loptt/lambdish-compiler/ast"
) 

// SemanticCheck
func SemanticCheck(program *ast.Program) error {
	funcdir := dir.NewFuncDirectory()

	return walkProgram(program, funcdir)
}

func walkProgram(program *ast.Program, funcdir *dir.FuncDirectory) error {
	for _, f := range program.functions {
		err := walkFunction(f, funcdir)
		if err != nil {
			return err
		}
	}

	err := walkFunctionCall(program.call, funcdir)
}

func walkFunction(function *Function, funcdir *dir.FuncDirectory) error {

}

func walkFunctionCall(call *FunctionCall, funcdir *dir.FuncDirectory) error {

}