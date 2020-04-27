package dir

import (
	"fmt"
	"github.com/Loptt/lambdish-compiler/types"
	"strings"
)

type FuncEntry struct {
	id         string
	returnval  *types.LambdishType
	params     []*types.LambdishType
	vardir     *VarDirectory
	lambdas    []*FuncEntry
}

// Key returns the key of the FuncEntry used for the FuncDirectory
func (fe *FuncEntry) Key() string {
	var b strings.Builder

	for _, p := range fe.params {
		b.WriteString(p.String())
	}

	return fmt.Sprintf("%s@%s", fe.id, b.String())
}

// Lambdas returns the array of FuncEntry which represents the lambdas inside this FuncEntry
func (fe *FuncEntry) Lambdas() []*FuncEntry {
	return fe.lambdas
}

// Id returns the name of the funcentry
func (fe *FuncEntry) Id() string {
	return fe.id
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
	return &FuncEntry{id, returnval, params, vardir, make([]*FuncEntry,0)}
}

func FuncEntryKey(id string, params []*types.LambdishType) string {
	var b strings.Builder

	for _, p := range params {
		b.WriteString(p.String())
	}

	return fmt.Sprintf("%s@%s", id, b.String())
}

// FuncDirectory represents a table of FuncEntry structs used to store all the function declarations
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

func (fd *FuncDirectory) FuncIdExists(id string) bool {
	for _, fe := range fd.table {
		if fe.Id() == id {
			return true
		}
	}

	return false
}


func NewFuncDirectory() *FuncDirectory {
	return &FuncDirectory{make(map[string]*FuncEntry)}
}


