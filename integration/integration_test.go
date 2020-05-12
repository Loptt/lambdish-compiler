package integration

import (
	"fmt"
	"os"
	"testing"

	"github.com/Loptt/lambdish-compiler/ast"
	"github.com/Loptt/lambdish-compiler/gocc/lexer"
	"github.com/Loptt/lambdish-compiler/gocc/parser"
	"github.com/Loptt/lambdish-compiler/ic"
	"github.com/Loptt/lambdish-compiler/sem"
	"github.com/davecgh/go-spew/spew"
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

func TestGenerateIntermediateCode(t *testing.T) {
	p := parser.NewParser()
	tests := []string{
		"test1.lsh",
	}

	for _, test := range tests {
		input, err := readFile(test)

		if err != nil {
			t.Fatalf("Error reading file %s", test)
		}

		s := lexer.NewLexer(input)
		pro, err := p.Parse(s)
		if err != nil {
			t.Fatalf("%s: %v", test, err)
		}

		program, ok := pro.(*ast.Program)
		if !ok {
			t.Fatalf("Cannot cast to Program")
		}

		fmt.Printf("======PROGRAM AST=======\n")
		spew.Dump(program)

		funcdir, err := sem.SemanticCheck(program)
		if err != nil {
			t.Fatalf("Error from semantic: %v", err)
		}

		fmt.Printf("\n======FUNCTION DIRECTORY=======\n")
		spew.Dump(funcdir)

		gen, err := ic.GenerateIntermediateCode(program, funcdir)
		if err != nil {
			t.Fatalf("Error from generate code: %v", err)
		}

		fmt.Printf("\n======INTERMEDIATE CODE=======\n")
		fmt.Printf("%s\n", gen)
	}
}
