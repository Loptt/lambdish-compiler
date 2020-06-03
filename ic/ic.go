//Package ic provides the generation of intermediate code
package ic

import (
	"github.com/Loptt/lambdish-compiler/ast"
	"github.com/Loptt/lambdish-compiler/dir"
	"github.com/Loptt/lambdish-compiler/mem"
	"github.com/Loptt/lambdish-compiler/sem"
)

//GenerationContext djsknfkjsdfkj
type GenerationContext struct {
	funcdir *dir.FuncDirectory
	semcube *sem.SemanticCube
	gen     *Generator
	vm      *mem.VirtualMemory
}

// FuncDir ...
func (ctx *GenerationContext) FuncDir() *dir.FuncDirectory {
	return ctx.funcdir
}

// SemCube ...
func (ctx *GenerationContext) SemCube() *sem.SemanticCube {
	return ctx.semcube
}

// Generator ...
func (ctx *GenerationContext) Generator() *Generator {
	return ctx.gen
}

//VM ...
func (ctx *GenerationContext) VM() *mem.VirtualMemory {
	return ctx.vm
}

//GenerateIntermediateCode calls the two main code generation functions, first the function to generate addresses and then the
// function to generate the code itself
func GenerateIntermediateCode(program *ast.Program, funcdir *dir.FuncDirectory) (*Generator, *mem.VirtualMemory, error) {
	ctx := &GenerationContext{funcdir, sem.NewSemanticCube(), NewGenerator(), mem.NewVirtualMemory()}

	// GenerateAddresses intilializes all entries in every VarDirectory with an address
	// assigned by the VirtualMemory manager
	// This function must be called before the generateCode function to ensure every
	// variable and constant has a valid address
	if err := generateAddressesProgram(program, ctx); err != nil {
		return nil, nil, err
	}

	// GenerateCodeProgram
	if err := generateCodeProgram(program, ctx); err != nil {
		return nil, nil, err
	}

	return ctx.gen, ctx.vm, nil
}
