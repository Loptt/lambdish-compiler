package ic

import (
	"github.com/Loptt/lambdish-compiler/mem"
	"github.com/Loptt/lambdish-compiler/dir"
	"github.com/Loptt/lambdish-compiler/ast"
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