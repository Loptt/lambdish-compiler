package sem

import (
	"github.com/Loptt/lambdish-compiler/ast"
	"github.com/Loptt/lambdish-compiler/dir"
	"github.com/Loptt/lambdish-compiler/types"
	"github.com/mewkiz/pkg/errutil"
)

// getIdTypeFromFuncStack returns the type of the given id by checking in the FuncEntry stack
// for its declaration, and once found, returns its type. If it is not found, the function
// returns a nil pointer.
func getIdTypeFromFuncStack(id *ast.Id, fes *dir.FuncEntryStack) (*types.LambdishType, error) {
	for !fes.Empty() {
		fe := fes.Top()
		if fe.VarDir().Exists(id.String()) {
			return fe.VarDir().Get(id.String()).Type(), nil
		}

		fes.Pop()
	}

	return nil, errutil.Newf("Id %s not declared in this scope", id.String())
}

//  getTypeFunctionCall
func getTypeFunctionCall(fcall *ast.FunctionCall, fes *dir.FuncEntryStack, funcdir *dir.FuncDirectory) (*types.LambdishType, error) {
/*
	// First check if the function call is referencing a function declared as a parameter in one of 
	// the functions of the scope. If so return the return type of that function
	if idExistsInFuncStack(fcall.Id(), fes) {
		t, err := getIdTypeFromFuncStack(fcall.Id())
		if t, err := getIdTypeFromFuncStack(fcall.Id()); err != nil {
			return nil, err
		}

		if !t.Function() {
			return nil, errutil.Newf("Cannot call %s as a function in this scope", fcall.Id())
		}

		return t.Type()
	} else {
		argTypes := make([]*types.LambdishType, 0)
	
		for _, arg := range fcall.args {
			t, err := getTypeStatement(arg)
			if err != nil {
				return false, err
			}
			argTypes = append(argTypes, t)
		}
	
		return funcdir.Exists(dir.FuncEntryKey(fcall.Id(), argTypes)), nil
	}
	*/

	return nil, nil
}

func getTypeConstantList(cl *ast.ConstantList, fes *dir.FuncEntryStack, funcdir *dir.FuncDirectory) (*types.LambdishType, error) {
	return nil, nil
}

func getTypeStatement(statement ast.Statement, fes *dir.FuncEntryStack, funcdir *dir.FuncDirectory) (*types.LambdishType, error) {
	if id, ok := statement.(*ast.Id); ok {
		return getIdTypeFromFuncStack(id, fes)
	} else if fcall, ok := statement.(*ast.FunctionCall); ok {
		return getTypeFunctionCall(fcall, fes, funcdir)
	} else if cv, ok := statement.(*ast.ConstantValue); ok {
		return cv.Type(), nil
	} else if cl, ok := statement.(*ast.ConstantList); ok {
		return getTypeConstantList(cl, fes, funcdir)
	} else if l, ok := statement.(*ast.Lambda); ok {
		return types.NewFuncLambdishType(l.Retval(), l.Params(), 0), nil
	}

	return nil, errutil.Newf("Statement cannot be casted to any valid form")
}