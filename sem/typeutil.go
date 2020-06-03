package sem

import (
	"github.com/Loptt/lambdish-compiler/ast"
	"github.com/Loptt/lambdish-compiler/dir"
	"github.com/Loptt/lambdish-compiler/gocc/token"
	"github.com/Loptt/lambdish-compiler/types"
	"github.com/mewkiz/pkg/errutil"
)

// getIdTypeFromFuncStack returns the type of the given id by checking in the FuncEntry stack
// for its declaration, and once found, returns its type. If it is not found, the function
// returns a nil pointer.
func getIDTypeFromFuncStack(id *ast.Id, fes *dir.FuncEntryStack) (*types.LambdishType, error) {
	fescpy := *fes
	for !fescpy.Empty() {
		fe := fescpy.Top()
		if fe.VarDir().Exists(id.String()) {
			return fe.VarDir().Get(id.String()).Type(), nil
		}

		fescpy.Pop()
	}

	return nil, errutil.NewNoPosf("Id %s not declared in this scope", id.String())
}

// IsReservedFunction ...
func IsReservedFunction(s string) bool {
	for _, f := range reservedFunctions {
		if s == f {
			return true
		}
	}
	return false
}

func getReservedFunctionType(id string, args []*types.LambdishType, tok *token.Token) (*types.LambdishType, error) {
	switch id {
	case "if":
		return checkAndGetIfType(id, args, tok)
	case "append":
		return checkAndGetAppendType(id, args, tok)
	case "empty":
		return checkAndGetEmptyType(id, args, tok)
	case "head":
		return checkAndGetHeadType(id, args, tok)
	case "tail":
		return checkAndGetTailType(id, args, tok)
	case "insert":
		return checkAndGetInsertType(id, args, tok)
	}

	return nil, errutil.NewNoPosf("Cannot find reserved function")
}

// getTypeFunctionCall
func getTypeFunctionCall(fcall *ast.FunctionCall, fes *dir.FuncEntryStack, funcdir *dir.FuncDirectory, semcube *SemanticCube) (*types.LambdishType, error) {

	argTypes, err := GetTypesFromArgs(fcall.Args(), fes, funcdir, semcube)
	if err != nil {
		return nil, err
	}

	// First we check if the statement of the call is an ID
	if id, ok := fcall.Statement().(*ast.Id); ok {
		// If it is an ID, we check if it is declared in the function stack
		if idExistsInFuncStack(id, fes) {
			// If it is declared in the function stack, we get its type from the stack and check for
			// any errors
			t, err := getIDTypeFromFuncStack(id, fes)
			if err != nil {
				return nil, err
			}

			// If the type returned is not a function, then we cannot call the function and we
			// return an error
			if !t.Function() {
				return nil, errutil.NewNoPosf("%+v: Cannot call %s as a function in this scope", fcall.Token(), id)
			}

			// Otherwise we return the return type of this function type
			return t.Retval(), nil
			// If the id is not in the function stack, we must check the global function directory for its definition
			// To do this we must fist get the types of all of its arguments in order to construct the
			// key so that we can query the Func Directory
		}

		// Once we got the info to query, we get the function entry and we return its return type
		if fe := funcdir.Get(id.String()); fe != nil {
			if err := argumentsMatchParameters(fcall, argTypes, fe.Params(), fes, funcdir, semcube); err != nil {
				return nil, err
			}
			return fe.ReturnVal(), nil
			// If it is not in the func directory, we must check if the function is an operation
		} else if isOperationFromSemanticCube(id.String()) {
			key := GetSemanticCubeKey(id.String(), argTypes)
			if basic, ok := semcube.Get(key); ok {
				return types.NewDataLambdishType(basic, 0), nil
			}
			return nil, errutil.NewNoPosf("%+v: Cannot perform operation %s on arguments %+v", fcall.Token(), id.String(), argTypes)

			// If it is not an operation, we must check if it is a reserverd function
		} else if IsReservedFunction(id.String()) {
			return getReservedFunctionType(id.String(), argTypes, id.Token())
		} else {
			return nil, errutil.NewNoPosf("%+v: Function %s not declared on local or global scope", fcall.Token(), id)
		}

	}
	t, err := GetTypeStatement(fcall.Statement(), fes, funcdir, semcube)
	if err != nil {
		return nil, err
	}
	if !t.Function() {
		return nil, errutil.NewNoPosf("%+v: Cannot call as a function in this scope", fcall.Token())
	}

	if err := argumentsMatchParameters(fcall, argTypes, t.Params(), fes, funcdir, semcube); err != nil {
		return nil, err
	}

	return t.Retval(), nil
}

//getTypeConstantList
func GetTypeConstantList(cl *ast.ConstantList, fes *dir.FuncEntryStack, funcdir *dir.FuncDirectory, semcube *SemanticCube) (*types.LambdishType, error) {
	// Make a list of all the types in the list
	ts := make([]*types.LambdishType, 0)

	if len(cl.Contents()) == 0 {
		return cl.Type(), nil
	}

	// Get the type of each content
	for _, s := range cl.Contents() {
		if typ, err := GetTypeStatement(s, fes, funcdir, semcube); err == nil {
			ts = append(ts, typ)
		} else {
			return nil, err
		}
	}

	// Save the type of the first element of the list
	listType := *ts[0]

	// If any type is different then we send an error
	for _, t := range ts {
		if !t.Equal(&listType) {
			return nil, errutil.NewNoPosf("%+v: Cannot create list of multiple types", cl.Token())
		}
	}

	// If all types are the same we declare the type of the list as one level above the type of its elements
	listType.IncreaseList()

	return &listType, nil
}

//GetTypeStatement
func GetTypeStatement(statement ast.Statement, fes *dir.FuncEntryStack, funcdir *dir.FuncDirectory, semcube *SemanticCube) (*types.LambdishType, error) {
	if id, ok := statement.(*ast.Id); ok {
		if t, err := getIDTypeFromFuncStack(id, fes); err == nil {
			return t, nil
		} else if fe := funcdir.Get(id.String()); fe != nil {
			return convertFuncEntryToLambdishType(fe), nil
		}
		return nil, errutil.NewNoPosf("%+v: Id %s not declared in local or global scope", id.Token(), id.String())
	} else if fcall, ok := statement.(*ast.FunctionCall); ok {
		return getTypeFunctionCall(fcall, fes, funcdir, semcube)
	} else if cv, ok := statement.(*ast.ConstantValue); ok {
		return cv.Type(), nil
	} else if cl, ok := statement.(*ast.ConstantList); ok {
		return GetTypeConstantList(cl, fes, funcdir, semcube)
	} else if l, ok := statement.(*ast.Lambda); ok {
		return types.NewFuncLambdishType(l.Retval(), l.Params(), 0), nil
	}
	return nil, errutil.NewNoPosf("Statement cannot be casted to any valid form")
}

func argumentsMatchParameters(fcall *ast.FunctionCall, args []*types.LambdishType, params []*types.LambdishType, fes *dir.FuncEntryStack, funcdir *dir.FuncDirectory, semcube *SemanticCube) error {
	if len(args) != len(params) {
		return errutil.NewNoPosf("%+v: function expects %d arguments, got %d", fcall.Token(), len(params), len(args))
	}

	for i, p := range params {
		if !(p.Equal(args[i])) {
			return errutil.NewNoPosf("%+v: Function call arguments do not match its parameters", fcall.Token())
		}
	}

	return nil
}

func GetTypesFromArgs(args []ast.Statement, fes *dir.FuncEntryStack, funcdir *dir.FuncDirectory, semcube *SemanticCube) ([]*types.LambdishType, error) {
	ts := make([]*types.LambdishType, 0)

	for _, arg := range args {
		if t, err := GetTypeStatement(arg, fes, funcdir, semcube); err == nil {
			ts = append(ts, t)
		} else {
			return nil, err
		}
	}

	return ts, nil
}

func convertFuncEntryToLambdishType(fe *dir.FuncEntry) *types.LambdishType {
	return types.NewFuncLambdishType(fe.ReturnVal(), fe.Params(), 0)
}
