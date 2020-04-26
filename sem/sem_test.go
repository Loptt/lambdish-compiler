package sem

import (
	"github.com/Loptt/lambdish-compiler/ast"
	"github.com/Loptt/lambdish-compiler/gocc/lexer"
	"github.com/Loptt/lambdish-compiler/gocc/parser"
	"github.com/davecgh/go-spew/spew"
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

func TestSemanticCheck(t *testing.T) {
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
		pro, errtest := p.Parse(s)
		if errtest != nil {
			t.Errorf("%s: %v", test, errtest)
		}

		program, ok := pro.(*ast.Program)
		if !ok {
			t.Fatalf("Cannot cast to Program")
		}

		funcdir, err := SemanticCheck(program)
		if err != nil {
			t.Errorf("Error from semantic: %v", err)
		}
		spew.Dump(funcdir)
	}
}
