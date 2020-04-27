package dir

import (
	"fmt"
	"github.com/Loptt/lambdish-compiler/types"
)

type VarEntry struct {
	id string
	t  *types.LambdishType
}
func(ve *VarEntry) Id() string {
	return ve.id
}

func(ve *VarEntry) Type() *types.LambdishType {
	return ve.t
}

type VarDirectory struct {
	table map[string]*VarEntry
}

func (e *VarEntry) String() string {
	return fmt.Sprintf("%s", e.id)
}

func NewVarEntry(id string, t *types.LambdishType) *VarEntry {
	return &VarEntry{id, t}
}

func (vd *VarDirectory) Add(e *VarEntry) bool {

	_, ok := vd.table[e.String()]
	if !ok {
		vd.table[e.String()] = e
	}
	return !ok
}

func (vd *VarDirectory) Get(key string) *VarEntry {

	if result, ok := vd.table[key]; ok {
		return result
	}

	return nil
}

func (vd *VarDirectory) Exists(key string) bool {

	_, ok := vd.table[key]
	return ok
}

func NewVarDirectory() *VarDirectory {
	return &VarDirectory{make(map[string]*VarEntry)}
}