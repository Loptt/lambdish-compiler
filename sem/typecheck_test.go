package sem

import (
	"github.com/Loptt/lambdish-compiler/ast"
	"github.com/Loptt/lambdish-compiler/dir"
	"github.com/Loptt/lambdish-compiler/gocc/lexer"
	"github.com/Loptt/lambdish-compiler/gocc/parser"
	//"github.com/davecgh/go-spew/spew"
	"testing"
)

func TestTypeCheckProgram(t *testing.T) {
	p := parser.NewParser()
	tests := []string{
		"tests/test1.lsh",
		"tests/test2.lsh",
		"tests/test3.lsh",
		"tests/test4.lsh",
		"tests/test5.lsh",
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
			t.Fatalf("%s: Cannot cast to Program", test)
		}

		funcdir := dir.NewFuncDirectory()
		semcube := NewSemanticCube()

		err = buildFuncDirProgram(program, funcdir)
		if err != nil {
			t.Errorf("%s: buildFuncDirProgram: %v", test, err)
		}

		err = scopeCheckProgram(program, funcdir, semcube)
		if err != nil {
			t.Errorf("%s: scopeCheckProgram: %v", test, err)
		}

		err = typeCheckProgram(program, funcdir, semcube)
		if err != nil {
			t.Errorf("%s: typeCheckProgram: %v", test, err)
		}
		
		//spew.Dump(program)
	}
}