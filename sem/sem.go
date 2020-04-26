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
    //  *If a function is declared twice an error will be returned
    //  *If two parameters in the same function have the same id
    //
	err := buildFuncDirProgram(program, funcdir)
	if err != nil {
		return nil, err
    }
    
    // Check the scope of function calls and variable uses.
    // Errors to check:
    //  *If a function is called that does not exist
    //  *If a variable is used and it has not been declared in the parameters
    //
    err := scopeCheckProgram(program, funcdir)
    if err != nil {
        return err
    }

	return funcdir, nil
}

// TODO: Checar que las funciones y variables que se usen existan

// TODO: Checar la cohesion de tipos
