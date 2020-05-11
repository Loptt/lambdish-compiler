package ic

import (
	"github.com/Loptt/lambdish-compiler/ast"
	"github.com/Loptt/lambdish-compiler/dir"
	"github.com/Loptt/lambdish-compiler/mem"
)

func getAddressFromFuncStack(id *ast.Id, fes *dir.FuncEntryStack) (mem.Address, bool) {
	fescpy := *fes
	for !fescpy.Empty() {
		fe := fescpy.Top()
		if fe.VarDir().Exists(id.String()) {
			return fe.VarDir().Get(id.String()).Address(), true
		}
		fescpy.Pop()
	}

	return mem.Address(-1), false
}

func isOnTopOfFuncStack(id *ast.Id, fes *dir.FuncEntryStack) bool {
	return fes.Top().VarDir().Exists(id.String())
}
