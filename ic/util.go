package ic

import (
	"strconv"

	"github.com/Loptt/lambdish-compiler/ast"
	"github.com/Loptt/lambdish-compiler/dir"
	"github.com/Loptt/lambdish-compiler/mem"
	"github.com/Loptt/lambdish-compiler/types"
	"github.com/mewkiz/pkg/errutil"
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

func getFunctionTypeFromFuncStack(id *ast.Id, fes *dir.FuncEntryStack) (*types.LambdishType, error) {
	fescpy := *fes
	for !fescpy.Empty() {
		fe := fescpy.Top()
		if fe.VarDir().Exists(id.String()) {
			return fe.VarDir().Get(id.String()).Type(), nil
		}
		fescpy.Pop()
	}

	return nil, errutil.NewNoPosf("%+v: Cannot find function in Function Entry stack", id.Token())
}

func getIntOfType(t *types.LambdishType) (int, error) {
	if t.Function() {
		return 4, nil
	}

	if t.List() > 0 {
		return 5, nil
	}

	return strconv.Atoi(t.String())
}
