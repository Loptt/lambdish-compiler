package sem

import (
	"github.com/Loptt/lambdish-compiler/ast"
	"github.com/Loptt/lambdish-compiler/dir"
)

// idExistsInFuncStack checks if the given id exists at any parameter declaration in the stack of 
// FuncEntry. If it exist, it means that the id has been declared and is in scope
func idExistsInFuncStack(id *ast.Id, fes *dir.FuncEntryStack) bool {
	for !fes.Empty() {
		fe := fes.Top()
		if fe.VarDir().Exists(id.String()) {
			return true
		}

		fes.Pop()
	}

	return false
}

