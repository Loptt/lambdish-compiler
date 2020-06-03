package sem

import (
	"github.com/Loptt/lambdish-compiler/ast"
	"github.com/Loptt/lambdish-compiler/dir"
)

// SemanticCheck calls the 3 main functions that perform the semantic analysis and
// reports any errors
func SemanticCheck(program *ast.Program) (*dir.FuncDirectory, error) {
	funcdir := dir.NewFuncDirectory()
	semcube := NewSemanticCube()

	// Build the function directory and their corresponding Var directiories
	// Errors to check:
	//  * If a function is declared twice
	//  * If two parameters in the same function have the same id
	//
	if err := buildFuncDirProgram(program, funcdir); err != nil {
		return nil, err
	}

	// Check the scope of function calls and variable uses.
	// Errors to check:
	//  * If a function is called that does not exist
	//  * If a variable is used and it has not been declared in the parameters
	//
	if err := scopeCheckProgram(program, funcdir, semcube); err != nil {
		return funcdir, err
	}

	// Check type cohesion
	// Errors to check:
	//  * If a function is called and no function in the funcdir match the argument types
	//  * If the value in the statement of a function does not match its return value
	//  * Illegal use of use of built-in functions
	//      - Arithmetic operators: +, -, *, /,%
	//      - Relational operators: < , >, equal
	//      - Logical operatios: (only to be used with bools) and, or, !
	//  * If statements:
	//      - Check that first argument is of type bool
	//      - Check that second and third arguments are of the same type
	//      - The type of the second and third argument will define the type of the if statement
	//  * To check whether a combination of params for an operator is valid, the semantic cube must be consulted
	//
	if err := typeCheckProgram(program, funcdir, semcube); err != nil {
		return funcdir, err
	}

	return funcdir, nil
}
