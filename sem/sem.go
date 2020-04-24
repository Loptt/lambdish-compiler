package sem

import (
	"github.com/Loptt/lambdish-compiler/dir"
	"github.com/Loptt/lambdish-compiler/ast"
) 

// SemanticCheck
func SemanticCheck(program *ast.Program) (*dir.FuncDirectory,error) {
	funcdir := dir.NewFuncDirectory()

	err := buildFuncDirProgram(program, funcdir);
	if err != nil {
		return nil,err
	}

	return funcdir, nil
}


// Construir directorio de funciones en func.go

// Checar que las funciones y variables que se usen existan u

// Checar la cohesion de tipos
