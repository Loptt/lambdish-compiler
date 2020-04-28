package sem

import (
	"github.com/Loptt/lambdish-compiler/dir"
	"github.com/mewkiz/pkg/errutil"
)

func idIsReserved(id string) bool {
	return id == "equal" || id == "and" || id == "or" || id == "if" || id == "append" || id == "head" || id == "empty" || id == "tail" || id == "insert"
}

func checkVarDirReserved(vardir *dir.VarDirectory) error {
	for _, f := range reservedFunctions {
		if vardir.Exists(f) {
			return errutil.Newf("Cannot declare variable with reserved keyword %s", f)
		}
	}

	return nil
}