package ic

import (
	"github.com/Loptt/lambdish-compiler/ast"
	"github.com/Loptt/lambdish-compiler/gocc/lexer"
	"github.com/Loptt/lambdish-compiler/gocc/parser"
	"github.com/Loptt/lambdish-compiler/mem"
	"github.com/Loptt/lambdish-compiler/sem"
	//"github.com/davecgh/go-spew/spew"
	"os"
	"testing"
)

func readFile(path string) ([]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	fileinfo, err := file.Stat()
	if err != nil {
		return nil, err
	}

	filesize := fileinfo.Size()
	buffer := make([]byte, filesize)

	_, err = file.Read(buffer)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

func TestGenerateCodeProgram(t *testing.T) {
	p := parser.NewParser()
	tests := []string{
		"tests/test5.lsh",
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

		funcdir, err := sem.SemanticCheck(program)
		if err != nil {
			t.Errorf("Error from semantic: %v", err)
		}

		semcube := sem.NewSemanticCube()
		vm := mem.NewVirtualMemory()

		err = generateCodeProgram(program, funcdir, semcube, NewGenerator(), vm)
		if err != nil {
			t.Errorf("Error from generate code: %v", err)
		}

	}
}
