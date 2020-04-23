package main

import (
	"github.com/Loptt/lambdish-compiler/gocc/lexer"
	"github.com/Loptt/lambdish-compiler/gocc/parser"
	"testing"
	"os"
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

func TestGrammar(t *testing.T) {
	p := parser.NewParser()
	tests := []string {
		"tests/test1.lsh",
		"tests/test2.lsh",
		"tests/test3.lsh",
		"tests/test4.lsh",
		"tests/test5.lsh",
	}


	for _, test := range tests {
		input, err := readFile(test)

		if err != nil {
			t.Fatalf("Error reading file %s", test);
		}

		s := lexer.NewLexer(input);
		_, errtest := p.Parse(s);

		if errtest != nil {
			t.Errorf("%s: %v", test, errtest);
		}
	}
}