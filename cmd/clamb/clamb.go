package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/Loptt/lambdish-compiler/ast"
	"github.com/Loptt/lambdish-compiler/gocc/lexer"
	"github.com/Loptt/lambdish-compiler/gocc/parser"
	"github.com/Loptt/lambdish-compiler/ic"
	"github.com/Loptt/lambdish-compiler/sem"
	"github.com/mewkiz/pkg/errutil"
)

func usage() {
	fmt.Printf("Usage: clamb <lambdish source file>\n")
}

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

func compile(file string) error {
	p := parser.NewParser()
	input, err := readFile(file)

	if err != nil {
		return err
	}

	s := lexer.NewLexer(input)
	pro, err := p.Parse(s)

	if err != nil {
		return err
	}

	program, ok := pro.(*ast.Program)
	if !ok {
		return errutil.NewNoPos("Cannot cast program")
	}

	funcdir, err := sem.SemanticCheck(program)
	if err != nil {
		return err
	}

	gen, vm, err := ic.GenerateIntermediateCode(program, funcdir)
	if err != nil {
		return err
	}

	fileroot := strings.Split(file, ".")

	return gen.CreateFile(fileroot[0], vm)
}

func main() {
	if len(os.Args) < 2 {
		usage()
		return
	}

	file := os.Args[1]

	err := compile(file)
	if err != nil {
		fmt.Printf("Compilation %v\n", err)
		return
	}

	fileroot := strings.Split(file, ".")[0]

	fmt.Printf("Compilation successful: file %s.obj generated\n", fileroot)
}
