package funcdir

import (
	"github.com/Loptt/lambdish-compiler/types"
	"fmt"
	"strings"
)

type entry struct {
	id string
	returnval types.LambdishType
	paramcount int
	params []types.LambdishType
}

func (e *entry) String() string {
	var b strings.Builder

	for _, p := range e.params {
		b.WriteString(p.String())
	}
	
	return fmt.Sprintf("%s@%s@%s", e.id, e.returnval, b.String())
}

type Funcdirectory struct {
	table map[string]entry
}

func (fd *Funcdirectory) Add(e entry) bool {

	_, ok := fd.table[e.String()]
	if ok {
		fd.table[e.String()] = e
	}
	return ok
}

func (fd *Funcdirectory) Get(key string) *entry {

	if result, ok := fd.table[key]; ok {
		return &result
	}

	return nil
}

func (fd *Funcdirectory) Exists(key string) bool {

	_, ok := fd.table[key]
	return ok
}