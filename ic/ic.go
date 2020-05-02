package ic

import (
	"github.com/Loptt/lambdish-compiler/ast"
	"github.com/Loptt/lambdish-compiler/dir"
	"github.com/Loptt/lambdish-compiler/mem"
	"github.com/Loptt/lambdish-compiler/sem"
)

type GenerationContext struct {
	funcdir *dir.FuncDirectory
	semcube *sem.SemanticCube
	gen *Generator
	vm *mem.VirtualMemory
}

func(ctx *GenerationContext) FuncDir() *dir.FuncDirectory{
	return ctx.funcdir
}
func(ctx *GenerationContext) SemCube() *sem.SemanticCube{
	return ctx.semcube
}
func(ctx *GenerationContext) Generator() *Generator{
	return ctx.gen
}
func(ctx *GenerationContext) VM() *mem.VirtualMemory{
	return ctx.vm
}

func GenerateIntermediateCode(program *ast.Program, funcdir *dir.FuncDirectory) (*Generator, error) {
	ctx := &GenerationContext{funcdir, sem.NewSemanticCube(), NewGenerator(), mem.NewVirtualMemory()}

	// GenerateAddresses intilializes all entries in every VarDirectory with an address
	// assigned by the VirtualMemory manager
	// This function must be called before the generateCode function to ensure every
	// variable and constant has a valid address
	if err := generateAddressesProgram(program, ctx); err != nil {
		return nil, err
	}

	// GenerateCodeProgram
	if err := generateCodeProgram(program, ctx); err != nil {
		return nil, err
	}

	return ctx.gen, nil
}
