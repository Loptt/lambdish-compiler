package dir

import (
	"fmt"
	"github.com/Loptt/lambdish-compiler/types"
	"strings"
)

type funcentry struct {
	id         string
	returnval  types.LambdishType
	paramcount int
	params     []types.LambdishType
}

func (e *funcentry) String() string {
	var b strings.Builder

	for _, p := range e.params {
		b.WriteString(p.String())
	}

	return fmt.Sprintf("%s@%s@%s", e.id, e.returnval, b.String())
}

type FuncDirectory struct {
	table map[string]funcentry
}

// Add function adds a new funcentry to the table. If the addition is successful
// the function returns true and false otherwise.
func (fd *FuncDirectory) Add(e funcentry) bool {

	_, ok := fd.table[e.String()]
	if !ok {
		fd.table[e.String()] = e
	}
	return !ok
}

func (fd *FuncDirectory) Get(key string) *funcentry {

	if result, ok := fd.table[key]; ok {
		return &result
	}

	return nil
}

func (fd *FuncDirectory) Exists(key string) bool {

	_, ok := fd.table[key]
	return ok
}
