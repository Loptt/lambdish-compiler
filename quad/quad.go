package quad

import (
	"fmt"

	"github.com/Loptt/lambdish-compiler/mem"
)

// Operation ...
type Operation int

const (
	Add Operation = iota
	Sub
	Mult
	Div
	Mod
	Lt
	Gt
	Equal
	And
	Or
	Not
	GotoT
	GotoF
	Goto
	Ret
	Era
	Param
	Call
	Emp
	Head
	Tail
	Ins
	App
	Lst
	GeLst
	PaLst
	Print
	Invalid
)

func (o Operation) String() string {
	switch o {
	case Add:
		return "+"
	case Sub:
		return "-"
	case Mult:
		return "*"
	case Div:
		return "/"
	case Mod:
		return "%"
	case Lt:
		return "<"
	case Gt:
		return ">"
	case Equal:
		return "Equal"
	case And:
		return "And"
	case Or:
		return "Or"
	case Not:
		return "!"
	case GotoT:
		return "GotoT"
	case GotoF:
		return "GotoF"
	case Goto:
		return "Goto"
	case Ret:
		return "Ret"
	case Call:
		return "Call"
	case Era:
		return "Era"
	case Param:
		return "Param"
	case Emp:
		return "Emp"
	case Head:
		return "Head"
	case Tail:
		return "Tail"
	case Ins:
		return "Ins"
	case App:
		return "App"
	case GeLst:
		return "GeLst"
	case Lst:
		return "Lst"
	case PaLst:
		return "PaLst"
	case Print:
		return "Print"
	}

	return ""
}

func GetOperation(s string) Operation {
	switch s {
	case "+":
		return Add
	case "-":
		return Sub
	case "*":
		return Mult
	case "/":
		return Div
	case "%":
		return Mod
	case "<":
		return Lt
	case ">":
		return Gt
	case "equal":
		return Equal
	case "and":
		return And
	case "or":
		return Or
	case "!":
		return Not
	case "empty":
		return Emp
	case "head":
		return Head
	case "tail":
		return Tail
	case "insert":
		return Ins
	case "append":
		return App
	}

	return Invalid
}

func StringToOperation(s string) Operation {
	switch s {
	case "+":
		return Add
	case "-":
		return Sub
	case "*":
		return Mult
	case "/":
		return Div
	case "%":
		return Mod
	case "<":
		return Lt
	case ">":
		return Gt
	case "Equal":
		return Equal
	case "And":
		return And
	case "Or":
		return Or
	case "!":
		return Not
	case "GotoT":
		return GotoT
	case "GotoF":
		return GotoF
	case "Goto":
		return Goto
	case "Ret":
		return Ret
	case "Call":
		return Call
	case "Era":
		return Era
	case "Param":
		return Param
	case "Emp":
		return Emp
	case "Head":
		return Head
	case "Tail":
		return Tail
	case "Ins":
		return Ins
	case "App":
		return App
	case "GeLst":
		return GeLst
	case "Lst":
		return Lst
	case "PaLst":
		return PaLst
	case "Print":
		return Print
	}

	return Invalid
}

/*
Exceptions:
– Unary operators: no arg2
– Operators like param: no arg2, no result
– (Un)conditional jumps: target label is the result
*/

type Quadruple struct {
	op Operation
	a1 mem.Address
	a2 mem.Address
	r  mem.Address
}

func (q *Quadruple) SetR(addr mem.Address) {
	q.r = addr
}

func (q *Quadruple) SetLop(addr mem.Address) {
	q.a1 = addr
}

func (q *Quadruple) SetRop(addr mem.Address) {
	q.a2 = addr
}

func (q *Quadruple) Op() Operation {
	return q.op
}

func (q *Quadruple) Lop() mem.Address {
	return q.a1
}

func (q *Quadruple) Rop() mem.Address {
	return q.a2
}

func (q *Quadruple) R() mem.Address {
	return q.r
}

func (q Quadruple) String() string {
	return fmt.Sprintf("%s %s %s %s", q.op, q.a1, q.a2, q.r)
}

func NewQuadruple(op Operation, a1, a2, r mem.Address) *Quadruple {
	return &Quadruple{op, a1, a2, r}
}
