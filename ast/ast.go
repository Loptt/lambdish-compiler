package ast

/* import (
	"strconv"
	"strings"
	"github.com/Loptt/lambdish-compiler/gocc/token"
) */

type Attribute interface{}

type Val interface {
	Attrib
	Eval() bool
	String() string
}

func AppendConstant(data Attribute) {

}
