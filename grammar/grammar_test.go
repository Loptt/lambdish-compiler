package main

import (
	"os"
	"testing"

	"github.com/Loptt/lambdish-compiler/gocc/lexer"
	"github.com/Loptt/lambdish-compiler/gocc/parser"
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

//TestGrammar Function that tests the grammar from the lexer and parser
func TestGrammar(t *testing.T) {
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
		program, errtest := p.Parse(s)

		spew.Dump(program)

		if errtest != nil {
			t.Errorf("%s: %v", test, errtest)
		}
	}
}
