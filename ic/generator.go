package ic

import (
	"github.com/Loptt/lambdish-compiler/mem"
	"fmt"
	"strings"
)

type Generator struct {
	jumpStack *AddressStack
	addrStack *AddressStack
	icounter  int
	quads     []*Quadruple
}

func NewGenerator() *Generator {
	return &Generator{NewAddressStack(), NewAddressStack(), 0, make([]*Quadruple, 0)}
}

func (g *Generator) JumpStack() *AddressStack {
	return g.jumpStack
}

func (g *Generator) Counter() int {
	return g.icounter
}

func (g *Generator) Quadruples() []*Quadruple {
	return g.quads
}

func (g *Generator) Generate(op Operation, a1, a2, r mem.Address) {
	g.quads = append(g.quads, NewQuadruple(op, a1, a2, r))
	g.icounter++
}

func (g *Generator) PushToAddrStack(a mem.Address) {
	g.addrStack.Push(a)
}

func (g *Generator) GetFromAddrStack() mem.Address {
	val := g.addrStack.Top()
	g.addrStack.Pop()
	return val
}

func (g *Generator) PushToJumpStack(a mem.Address) {
	g.addrStack.Push(a)
}

func (g *Generator) GetFromJumpStack() mem.Address {
	val := g.jumpStack.Top()
	g.jumpStack.Pop()
	return val
}

func (g *Generator) FillJumpQuadruple(location mem.Address, jump mem.Address) {
	g.quads[int(location)].r = jump
}

func (g *Generator) String() string {
	var builder strings.Builder
	builder.WriteString("Generator:\n")
	builder.WriteString(fmt.Sprintf("  JumpStack: %s\n", g.jumpStack))
	builder.WriteString(fmt.Sprintf("  Instruction Counter: %d\n", g.icounter))
	builder.WriteString("  Quads:\n")

	for _, q := range g.quads {
		builder.WriteString(fmt.Sprintf("    %s\n", q))
	}

	return builder.String()
}