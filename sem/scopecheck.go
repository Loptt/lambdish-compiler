package sem

import (
	"github.com/Loptt/lambdish-compiler/dir"
)

func scopeCheckProgram(program *Program, funcdir *dir.FuncDirectory) error {
	for _, f := range program.Functions() {
		err := scopeCheckFunction(f, funcdir)
		if err != nil {
			return err
		}
	}

	return nil
}