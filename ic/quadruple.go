package ic

type Operation int

const (
	Add Operation = iota
	Sub
	Mult
)

type Quadruple struct {
	op Operation
	a1 int
	a2 int
	r  int
}

func NewQuadruple(op Operation, a1, a2, r int) *Quadruple {
	return &Quadruple{op, a1, a2, r}
}