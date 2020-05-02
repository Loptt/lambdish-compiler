package ic

type Generator struct {
	jumpStack *JumpStack
	icounter  int
	quads     []*Quadruple
}

func NewGenerator() *Generator {
	return &Generator{NewJumpStack(), 0, make([]*Quadruple, 0)}
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

func (g *Generator) Increase() {
	g.counter++
}

func (g *Generator) Generate(op Operation, a1, a2, r mem.Address) {
	g.quads = append(g.quads, NewQuadruple(op, a1, a2, r))
	icounter++
}
