package sem

import (
	"github.com/Loptt/lambdish-compiler/ast"
	"github.com/Loptt/lambdish-compiler/dir"
)

// SemanticCheck: Construcción
func SemanticCheck(program *ast.Program) (*dir.FuncDirectory, error) {
	funcdir := dir.NewFuncDirectory()

	err := buildFuncDirProgram(program, funcdir)
	if err != nil {
		return nil, err
	}

	return funcdir, nil
}

// TODO: Checar que las funciones y variables que se usen existan

// TODO: Checar la cohesion de tipos
