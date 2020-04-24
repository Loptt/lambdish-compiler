package dir

import (
	"fmt"
	"github.com/Loptt/lambdish-compiler/types"
	"strings"
)

type FuncEntry struct {
	id         string
	returnval  *types.LambdishType
	paramcount int
	params     []*types.LambdishType
	vardir     *VarDirectory
}

func (e *FuncEntry) Key() string {
	var b strings.Builder

	for _, p := range e.params {
		b.WriteString(p.String())
	}

	return fmt.Sprintf("%s@%s@%s", e.id, e.returnval, b.String())
}

func NewFuncEntry(id string, returnval *types.LambdishType, paramcount int, params []*types.LambdishType, vardir *VarDirectory) *FuncEntry {
	return &FuncEntry{id, returnval, paramcount, params, vardir}
}

type FuncDirectory struct {
	table map[string]*FuncEntry
}

// Add function adds a new funcentry to the table. If the addition is successful
// the function returns true and false otherwise.
func (fd *FuncDirectory) Add(e *FuncEntry) bool {

	_, ok := fd.table[e.Key()]
	if !ok {
		fd.table[e.Key()] = e
	}
	return !ok
}

func (fd *FuncDirectory) Get(key string) *FuncEntry {

	if result, ok := fd.table[key]; ok {
		return result
	}

	return nil
}

func (fd *FuncDirectory) Exists(key string) bool {

	_, ok := fd.table[key]
	return ok
}

func NewFuncDirectory() *FuncDirectory {
	return &FuncDirectory{make(map[string]*FuncEntry)}
}
