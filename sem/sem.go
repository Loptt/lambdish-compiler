package sem

import (
	"strconv"
	"strings"
	"github.com/Loptt/lambdish-compiler/types"
	"github.com/Loptt/lambdish-compiler/dir"
) 

type Param struct {
	id string,
	t  types.LambdishType
}

func AppendFunction(e *FuncEntry, fd *FuncDirectory) (*FuncDirectory,error){
	fd.Add(e)
	return fd,nil
}

// NewFunctionEntry Function that recieves the id of the function, an array of LambdishTypes and the  return type
func NewFunctionEntry(id string, params []*dir.VarEntry, returntype types.LambdishType) (*FuncEntry, error) {
	var paramtypes []types.LambdishType

	for _, p := params {
		paramtypes = append(paramtypes, p.t)
	}

	return dir.NewFuncEntry(id, returntype, len(params), paramtypes), nil
}

func AppendParams(t types.LambdishType, id string, params []*dir.VarEntry) ([]*dir.VarEntry, error){
	newParam := dir.NewVarEntry(id,t)

	return append(params,newParam), nil
}

func NewParams(t types.LambdishType, id string) ([]*dir.VarEntry, error){

	newParam := dir.NewVarEntry(id, t)

	return append(make([]*dir.VarEntry),newParam), nil
}

func NewParams() ([]*dir.VarEntry, error){
	
	return make([]*dir.VarEntry), nil
}




/* // NewLambdaEntry Function that recieves an array of LambdishTypes and registers it in the functionDirectory
func NewLambdaEntry(params []types.LambdishType) (*FuncEntry,error){

}

func NewLambdishType(t string) (types.LambdishType, nil) {

} */