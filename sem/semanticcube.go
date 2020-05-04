package sem

import (
	"fmt"
	"strings"

	"github.com/Loptt/lambdish-compiler/types"
	"github.com/mewkiz/pkg/errutil"
)

var reservedFunctions = []string{"if", "append", "empty", "head", "tail", "insert", "and", "or", "equal", "+", "-", "*", "/", "%", "<", ">", "!", "main"}

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
)

func (o Operation) convert() string {
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

// GetOperation ...
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

	return Add
}

// SemanticCube ...
type SemanticCube struct {
	operations map[string]types.BasicType
}

// NewSemanticCube ...
func NewSemanticCube() *SemanticCube {

	return &SemanticCube{
		map[string]types.BasicType{
			//Arithmetical Operators
			"+@11": types.Num,
			"-@11": types.Num,
			"/@11": types.Num,
			"*@11": types.Num,
			"%@11": types.Num,
			//Relational Operators
			"<@11":     types.Bool,
			">@11":     types.Bool,
			">@22":     types.Bool,
			"<@22":     types.Bool,
			"equal@11": types.Bool,
			"equal@22": types.Bool,
			"equal@33": types.Bool,
			//Logical Operators
			"and@33": types.Bool,
			"or@33":  types.Bool,
			"!@3":    types.Bool,
		},
	}
}

func (c *SemanticCube) Get(key string) (types.BasicType, bool) {
	typ, ok := c.operations[key]
	if !ok {
		return types.Null, false
	}
	return typ, true
}

//isOperationFromSemanticCube
func isOperationFromSemanticCube(s string) bool {
	switch s {
	case "+":
		return true
	case "-":
		return true
	case "*":
		return true
	case "/":
		return true
	case "%":
		return true
	case "<":
		return true
	case ">":
		return true
	case "equal":
		return true
	case "and":
		return true
	case "or":
		return true
	case "!":
		return true
	}

	return false
}

func checkAndGetIfType(id string, args []*types.LambdishType) (*types.LambdishType, error) {
	if len(args) != 3 {
		return nil, errutil.Newf("Arguments for if must be exactly 3")
	}
	if args[0].Basic() != types.Bool {
		return nil, errutil.Newf("The first argument for if must be of type bool, got %s", args[0])
	}
	if !args[1].Equal(args[2]) {
		return nil, errutil.Newf("The second and third arguments for if must be of the same type. Got %s and %s", args[1], args[2])
	}
	return args[1], nil
}

func checkAndGetAppendType(id string, args []*types.LambdishType) (*types.LambdishType, error) {
	if len(args) != 2 {
		return nil, errutil.Newf("Arguments for append must be exactly 2")
	}
	if args[0].List() < 1 {
		return nil, errutil.Newf("Arguments for append must be a list")
	}
	if !args[0].Equal(args[1]) {
		return nil, errutil.Newf("Arguments for append must be lists of the same type")
	}

	return args[0], nil
}

func checkAndGetEmptyType(id string, args []*types.LambdishType) (*types.LambdishType, error) {
	if len(args) != 1 {
		return nil, errutil.Newf("Arguments for empty must be exactly 1")
	}
	if args[0].List() < 1 {
		return nil, errutil.Newf("Arguments for empty must be a list")
	}

	return types.NewDataLambdishType(types.Bool, 0), nil
}

func checkAndGetHeadType(id string, args []*types.LambdishType) (*types.LambdishType, error) {
	if len(args) != 1 {
		return nil, errutil.Newf("Arguments for head must be exactly 1")
	}
	if args[0].List() < 1 {
		return nil, errutil.Newf("Arguments for head must be a list")
	}

	t := *args[0]
	t.DecreaseList()
	return &t, nil
}

func checkAndGetTailType(id string, args []*types.LambdishType) (*types.LambdishType, error) {
	if len(args) != 1 {
		return nil, errutil.Newf("Arguments for tail must be exactly 1")
	}
	if args[0].List() < 1 {
		return nil, errutil.Newf("Arguments for tail must be a list")
	}

	return args[0], nil
}

func checkAndGetInsertType(id string, args []*types.LambdishType) (*types.LambdishType, error) {
	if len(args) != 2 {
		return nil, errutil.Newf("Arguments for insert must be exactly 2")
	}
	if args[1].List() < 1 {
		return nil, errutil.Newf("Second argument for insert must be a list")
	}
	t1 := *args[0]
	t2 := &t1
	t2.IncreaseList()

	if !t2.Equal(args[1]) {
		return nil, errutil.Newf("Second argument for insert must be a list of the first argument %s %s", t2, args[1])
	}

	return args[1], nil
}

func getSemanticCubeKey(id string, params []*types.LambdishType) string {
	var b strings.Builder

	for _, p := range params {
		b.WriteString(p.String())
	}

	return fmt.Sprintf("%s@%s", id, b.String())
}
