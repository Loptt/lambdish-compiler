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
	fescpy := *fes
	for !fescpy.Empty() {
		fe := fescpy.Top()
		if fe.VarDir().Exists(id.String()) {
			return fe.VarDir().Get(id.String()).Type(), nil
		}

		fescpy.Pop()
	}

	return nil, errutil.Newf("Id %s not declared in this scope", id.String())
}

//IsReservedFunction
func IsReservedFunction(s string) bool {
	for _, f := range reservedFunctions {
		if s == f {
			return true
		}
	}
	return false
}

func getReservedFunctionType(id string, args []*types.LambdishType) (*types.LambdishType, error) {
	switch id {
	case "if":
		return checkAndGetIfType(id, args)
	case "append":
		return checkAndGetAppendType(id, args)
	case "empty":
		return checkAndGetEmptyType(id, args)
	case "head":
		return checkAndGetHeadType(id, args)
	case "tail":
		return checkAndGetTailType(id, args)
	case "insert":
		return checkAndGetInsertType(id, args)
	}

	return nil, errutil.Newf("Cannot find reserved function")
}

//  getTypeFunctionCall
func getTypeFunctionCall(fcall *ast.FunctionCall, fes *dir.FuncEntryStack, funcdir *dir.FuncDirectory, semcube *SemanticCube) (*types.LambdishType, error) {

	// First we check if the statement of the call is an ID
	if id, ok := fcall.Statement().(*ast.Id); ok {
		// If it is an ID, we check if it is declared in the function stack
		if idExistsInFuncStack(id, fes) {
			// If it is declared in the function stack, we get its type from the stack and check for
			// any errors
			t, err := getIdTypeFromFuncStack(id, fes)
			if err != nil {
				return nil, err
			}

			// If the type returned is not a function, then we cannot call the function and we
			// return an error
			if !t.Function() {
				return nil, errutil.Newf("%+v: Cannot call %s as a function in this scope", fcall.Token(), id)
			}

			// Otherwise we return the return type of this function type
			return t.Retval(), nil
			// If the id is not in the function stack, we must check the global function directory for its definition
			// To do this we must fist get the types of all of its arguments in order to construct the
			// key so that we can query the Func Directory
		} else {
			argTypes := make([]*types.LambdishType, 0)

			for _, arg := range fcall.Args() {
				t, err := getTypeStatement(arg, fes, funcdir, semcube)
				if err != nil {
					return nil, err
				}
				argTypes = append(argTypes, t)
			}

			// Once we got the info to query, we get the function entry and we return its return type
			if fe := funcdir.Get(id.String()); fe != nil {
				return fe.ReturnVal(), nil
				// If it is not in the func directory, we must check if the function is an operation
			} else if isOperationFromSemanticCube(id.String()) {
				key := getSemanticCubeKey(id.String(), argTypes)
				if basic, ok := semcube.Get(key); ok {
					return types.NewDataLambdishType(basic, 0), nil
				} else {
					return nil, errutil.Newf("%+v: Cannot perform operation %s on arguments %+v", fcall.Token(), id.String(), argTypes)
				}
				// If it is not an operation, we must check if it is a reserverd function
			} else if IsReservedFunction(id.String()) {
				return getReservedFunctionType(id.String(), argTypes)
			} else {
				return nil, errutil.Newf("%+v: Function %s not declared on local or global scope", fcall.Token(), id)
			}
		}
	} else {
		t, err := getTypeStatement(fcall.Statement(), fes, funcdir, semcube)
		if err != nil {
			return nil, err
		}
		if !t.Function() {
			return nil, errutil.Newf("%+v: Cannot call as a function in this scope", fcall.Token())
		}

		return t.Retval(), nil
	}

	return nil, nil
}

//getTypeConstantList
func getTypeConstantList(cl *ast.ConstantList, fes *dir.FuncEntryStack, funcdir *dir.FuncDirectory, semcube *SemanticCube) (*types.LambdishType, error) {
	ts := make([]*types.LambdishType, 0)

	if len(cl.Contents()) == 0 {
		return nil, errutil.Newf("Empty list delcaration currently not supported")
	}

	for _, s := range cl.Contents() {
		if typ, err := getTypeStatement(s, fes, funcdir, semcube); err == nil {
			ts = append(ts, typ)
		} else {
			return nil, err
		}
	}

	listType := *ts[0]

	for _, t := range ts {
		if !t.Equal(&listType) {
			return nil, errutil.Newf("%+v: Cannot create list of multiple types", cl.Token())
		}
	}

	listType.IncreaseList()

	return &listType, nil
}

//getTypeStatement
func getTypeStatement(statement ast.Statement, fes *dir.FuncEntryStack, funcdir *dir.FuncDirectory, semcube *SemanticCube) (*types.LambdishType, error) {
	if id, ok := statement.(*ast.Id); ok {
		return getIdTypeFromFuncStack(id, fes)
	} else if fcall, ok := statement.(*ast.FunctionCall); ok {
		return getTypeFunctionCall(fcall, fes, funcdir, semcube)
	} else if cv, ok := statement.(*ast.ConstantValue); ok {
		return cv.Type(), nil
	} else if cl, ok := statement.(*ast.ConstantList); ok {
		return getTypeConstantList(cl, fes, funcdir, semcube)
	} else if l, ok := statement.(*ast.Lambda); ok {
		return types.NewFuncLambdishType(l.Retval(), l.Params(), 0), nil
	}
	return nil, errutil.Newf("Statement cannot be casted to any valid form")
}
