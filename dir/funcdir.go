package dir

import (
	"github.com/Loptt/lambdish-compiler/types"
)

type FuncEntry struct {
	id        string
	returnval *types.LambdishType
	params    []*types.LambdishType
	vardir    *VarDirectory
	lambdas   []*FuncEntry
}

// Key returns the key of the FuncEntry used for the FuncDirectory
func (fe *FuncEntry) Key() string {
	return fe.id
}

// Lambdas returns the array of FuncEntry which represents the lambdas inside this FuncEntry
func (fe *FuncEntry) Lambdas() []*FuncEntry {
	return fe.lambdas
}

// Id returns the name of the funcentry
func (fe *FuncEntry) Id() string {
	return fe.id
}

// Params returns the name of the funcentry
func (fe *FuncEntry) Params() []*types.LambdishType {
	return fe.params
}

// ReturnVal returns the name of the funcentry
func (fe *FuncEntry) ReturnVal() *types.LambdishType {
	return fe.returnval
}

// AddLambda adds a new lambda func entry to the current func entry
func (fe *FuncEntry) AddLambda(retval *types.LambdishType, params []*types.LambdishType, vardir *VarDirectory) *FuncEntry {
	id := string(len(fe.lambdas))
	lambda := &FuncEntry{id, retval, params, vardir, make([]*FuncEntry, 0)}
	fe.lambdas = append([]*FuncEntry{lambda}, fe.lambdas...)
	return lambda
}

// VarDir returns the Var Directory of the FuncEntry
func (fe *FuncEntry) VarDir() *VarDirectory {
	return fe.vardir
}

// VarDir returns the Var Directory of the FuncEntry
func (fe *FuncEntry) GetLambdaEntryById(id string) *FuncEntry {
	for _, l := range fe.lambdas {
		if l.Id() == id {
			return l
		}
	}

	return nil
}

// NewFuncEntry creates a new FuncEntry struct
func NewFuncEntry(id string, returnval *types.LambdishType, params []*types.LambdishType, vardir *VarDirectory) *FuncEntry {
	return &FuncEntry{id, returnval, params, vardir, make([]*FuncEntry, 0)}
}

func FuncEntryKey(id string) string {
	return id
}

// FuncDirectory represents a table of FuncEntry structs used to store all the function declarations
type FuncDirectory struct {
	table map[string]*FuncEntry
}

func (fd *FuncDirectory) Table() map[string]*FuncEntry {
	return fd.table
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
