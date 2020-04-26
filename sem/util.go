package sem

import (
	"github.com/Loptt/lambdish-compiler/ast"
	"github.com/Loptt/lambdish-compiler/dir"
)

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

func functionCallExistsInFuncDirectory(fcall *ast.FunctionCall, funcdir *dir.FuncDirectory) bool {
	return true
}
/*
func getTypeStatement(statement Statement, funcdir) *types.LambdishType {

}*/