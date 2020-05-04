package ic

import (
	"fmt"
	"strings"

	"github.com/Loptt/lambdish-compiler/dir"
	"github.com/Loptt/lambdish-compiler/mem"
)

// Generator ...
type Generator struct {
	jumpStack       *AddressStack
	addrStack       *AddressStack
	icounter        int
	pcounter        int
	quads           []*Quadruple
	pendingFuncAddr map[int]string
}

// NewGenerator
func NewGenerator() *Generator {
	return &Generator{NewAddressStack(), NewAddressStack(), 0, 0, make([]*Quadruple, 0), make(map[int]string)}
}

// JumpStacks

func (g *Generator) JumpStack() *AddressStack {
	return g.jumpStack
}

// ICounter gets the current instruction counter
func (g *Generator) ICounter() int {
	return g.icounter
}

func (g *Generator) GetNextPCounter() int {
	val := g.pcounter
	g.pcounter += 1
	return val
}

func (g *Generator) ResetPCounter() {
	g.pcounter = 0
}

// Quadruples
func (g *Generator) Quadruples() []*Quadruple {
	return g.quads
}

// Generate
func (g *Generator) Generate(op Operation, a1, a2, r mem.Address) {
	g.quads = append(g.quads, NewQuadruple(op, a1, a2, r))
	g.icounter++
}

//PushToAddrStack
func (g *Generator) PushToAddrStack(a mem.Address) {
	g.addrStack.Push(a)
}

func (g *Generator) GetFromAddrStack() mem.Address {
	val := g.addrStack.Top()
	g.addrStack.Pop()
	return val
}

func (g *Generator) PushToJumpStack(a mem.Address) {
	g.jumpStack.Push(a)
}

func (g *Generator) GetFromJumpStack() mem.Address {
	val := g.jumpStack.Top()
	g.jumpStack.Pop()
	return val
}

func (g *Generator) FillJumpQuadruple(location mem.Address, jump mem.Address) {
	g.quads[int(location)].r = jump
}

func (g *Generator) AddPendingFuncAddr(loc int, id string) {
	g.pendingFuncAddr[loc] = id
}

func (g *Generator) FillPendingFuncAddr(funcdir *dir.FuncDirectory) {
	for loc, id := range g.pendingFuncAddr {
		fe := funcdir.Get(id)
		g.quads[loc].a1 = fe.Loc()
	}
}

func (g *Generator) String() string {
	var builder strings.Builder
	builder.WriteString("Generator:\n")
	builder.WriteString(fmt.Sprintf("  JumpStack: %s\n", g.jumpStack))
	builder.WriteString(fmt.Sprintf("  AddrStack: %s\n", g.addrStack))
	builder.WriteString(fmt.Sprintf("  Instruction Counter: %d\n", g.icounter))
	builder.WriteString("  Quads:\n")

	for i, q := range g.quads {
		builder.WriteString(fmt.Sprintf("    %d: %s\n", i, q))
	}

	return builder.String()
}
