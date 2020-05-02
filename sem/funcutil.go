package sem

import (
	"github.com/Loptt/lambdish-compiler/dir"
)

func idIsReserved(id string) bool {
	return id == "equal" || id == "and" || id == "or" || id == "if" || id == "append" || id == "head" || id == "empty" || id == "tail" || id == "insert"
}

func checkVarDirReserved(vardir *dir.VarDirectory) (string, bool) {
	for _, f := range reservedFunctions {
		if vardir.Exists(f) {
			return f, false
		}
	}

	return "", true
}
