package sem

import (
	"github.com/Loptt/lambdish-compiler/ast"
	"github.com/Loptt/lambdish-compiler/dir"
)

// idExistsInFuncStack checks if the given id exists at any parameter declaration in the stack of
// FuncEntry. If it exist, it means that the id has been declared and is in scope
func idExistsInFuncStack(id *ast.Id, fes *dir.FuncEntryStack) bool {
	fescpy := *fes
	for !fescpy.Empty() {
		fe := fescpy.Top()
		if fe.VarDir().Exists(id.String()) {
			return true
		}

		fescpy.Pop()
	}

	return false
}

// idExistsInFuncDir checks if the given id is the name of a function declared in the
// function directory.
func idExistsInFuncDir(id *ast.Id, funcdir *dir.FuncDirectory) bool {
	return funcdir.Exists(id.String())
}
