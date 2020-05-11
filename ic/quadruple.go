package ic

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
	return fmt.Sprintf("%s %s %s %s", q.op, q.a1, q.a2, q.r)
}
