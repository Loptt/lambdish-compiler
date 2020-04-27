package sem

import (
	"github.com/Loptt/lambdish-compiler/ast"
	"github.com/Loptt/lambdish-compiler/dir"
)

// SemanticCheck: Construcción
func SemanticCheck(program *ast.Program) (*dir.FuncDirectory, error) {
	funcdir := dir.NewFuncDirectory()

    // Build the function directory and their corresponding Var directiories
    // Errors to check:
    //  *If a function is declared twice
    //  *If two parameters in the same function have the same id
    //
	if err := buildFuncDirProgram(program, funcdir); err != nil {
		return nil, err
    }
    
    // Check the scope of function calls and variable uses.
    // Errors to check:
    //  *If a function is called that does not exist
    //  *If a variable is used and it has not been declared in the parameters
    //
    if err := scopeCheckProgram(program, funcdir); err != nil {
        return funcdir, err
    }

	return funcdir, nil
}

// TODO: Checar que las funciones y variables que se usen existan

// TODO: Checar la cohesion de tipos
