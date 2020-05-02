package ic

import (
	"github.com/Loptt/lambdish-compiler/mem"
	"fmt"
)

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
		return "equal"
	case And:
		return "and"
	case Or:
		return "or"
	case Not:
		return "!"
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

func NewQuadruple(op Operation, a1, a2, r mem.Address) *Quadruple {
	return &Quadruple{op, a1, a2, r}
}

func (q Quadruple) String() string {
	return fmt.Sprintf("%s %d %d %d", q.op, q.a1, q.a2, q.r)
}
