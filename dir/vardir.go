package dir

import (
	"github.com/Loptt/lambdish-compiler/types"
	"fmt"
)

type varentry struct {
	id string
	t types.LambdishType
	value string
}

type VarDirectory struct {
	table map[string]varentry
}

func (e *varentry) String() string {
	return fmt.Sprintf("%s", e.id)
}

func (vd *VarDirectory) Add(e varentry) bool {

	_, ok := vd.table[e.String()]
	if !ok {
		vd.table[e.String()] = e
	}
	return !ok
}

func (vd *VarDirectory) Get(key string) *varentry {

	if result, ok := vd.table[key]; ok {
		return &result
	}

	return nil
}

func (vd *VarDirectory) Exists(key string) bool {

	_, ok := vd.table[key]
	return ok
}