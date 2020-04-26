package sem

import (
	"github.com/Loptt/lambdish-compiler/ast"
	"github.com/Loptt/lambdish-compiler/dir"
	"github.com/Loptt/lambdish-compiler/gocc/lexer"
	"github.com/Loptt/lambdish-compiler/gocc/parser"
	"github.com/davecgh/go-spew/spew"
	"testing"
)

func TestScopeCheckProgram(t *testing.T) {
	p := parser.NewParser()
	tests := []string{
		"tests/test6.lsh",
	}

	for _, test := range tests {
		input, err := readFile(test)

		if err != nil {
			t.Fatalf("Error reading file %s", test)
		}

		s := lexer.NewLexer(input)
		pro, err := p.Parse(s)
		if err != nil {
			t.Errorf("%s: %v", test, err)
		}

		program, ok := pro.(*ast.Program)
		if !ok {
			t.Fatalf("Cannot cast to Program")
		}

		funcdir := dir.NewFuncDirectory()

		err = buildFuncDirProgram(program, funcdir)
		if err != nil {
			t.Errorf("buildFuncDirProgram: %v", err)
		}

		err = scopeCheckProgram(program, funcdir)
		if err != nil {
			t.Errorf("scopeCheckProgram: %v", err)
		}

		spew.Dump(funcdir)
	}
}