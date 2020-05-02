package ic

import (
	"github.com/Loptt/lambdish-compiler/mem"
	"fmt"
	"strings"
)

type Generator struct {
	jumpStack *JumpStack
	addrStack *AddressStack
	icounter  int
	quads     []*Quadruple
}

func NewGenerator() *Generator {
	return &Generator{NewJumpStack(), NewAddressStack(), 0, make([]*Quadruple, 0)}
}

func (g *Generator) JumpStack() *JumpStack {
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

func (g *Generator) PushToJumpStack(a int) {
	g.addrStack.Push(a)
}

func (g *Generator) GetFromJumpStack() int {
	val := g.JumpStack.Top()
	g.JumpStack.Pop()
	return val
}

func (g *Generator) FillQuadruple(location int, jump int)

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