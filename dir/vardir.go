package dir

import (
	"fmt"
	"github.com/Loptt/lambdish-compiler/gocc/token"
	"github.com/Loptt/lambdish-compiler/mem"
	"github.com/Loptt/lambdish-compiler/types"
)

type VarEntry struct {
	id   string
	t    *types.LambdishType
	tok  *token.Token
	addr mem.Address
}

func (ve *VarEntry) Id() string {
	return ve.id
}

func (ve *VarEntry) Type() *types.LambdishType {
	return ve.t
}

func (ve *VarEntry) Token() *token.Token {
	return ve.tok
}

func (ve *VarEntry) Address() mem.Address {
	return ve.addr
}

func (ve *VarEntry) SetAddress(addr mem.Address) {
	ve.addr = addr
}

type VarDirectory struct {
	table map[string]*VarEntry
}

func (e *VarEntry) String() string {
	return fmt.Sprintf("%s", e.id)
}

func NewVarEntry(id string, t *types.LambdishType, tok *token.Token) *VarEntry {
	return &VarEntry{id, t, tok, 0}
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

func (vd *VarDirectory) Table() map[string]*VarEntry {
	return vd.table
}

func NewVarDirectory() *VarDirectory {
	return &VarDirectory{make(map[string]*VarEntry)}
}
