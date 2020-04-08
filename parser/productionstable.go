// Code generated by gocc; DO NOT EDIT.

package parser



type (
	//TODO: change type and variable names to be consistent with other tables
	ProdTab      [numProductions]ProdTabEntry
	ProdTabEntry struct {
		String     string
		Id         string
		NTType     int
		Index      int
		NumSymbols int
		ReduceFunc func([]Attrib) (Attrib, error)
	}
	Attrib interface {
	}
)

var productionsTable = ProdTab{
	ProdTabEntry{
		String: `S' : Program	<<  >>`,
		Id:         "S'",
		NTType:     0,
		Index:      0,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Program : List	<<  >>`,
		Id:         "Program",
		NTType:     1,
		Index:      1,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `List : FunctionOrCall List	<<  >>`,
		Id:         "List",
		NTType:     2,
		Index:      2,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `List : empty	<<  >>`,
		Id:         "List",
		NTType:     2,
		Index:      3,
		NumSymbols: 0,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return nil, nil
		},
	},
	ProdTabEntry{
		String: `FunctionOrCall : Function	<<  >>`,
		Id:         "FunctionOrCall",
		NTType:     3,
		Index:      4,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `FunctionOrCall : FunctionCall	<<  >>`,
		Id:         "FunctionOrCall",
		NTType:     3,
		Index:      5,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Function : id "::" Params "=>" "(" Statement ")"	<<  >>`,
		Id:         "Function",
		NTType:     4,
		Index:      6,
		NumSymbols: 7,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Params : Type id "," Params	<<  >>`,
		Id:         "Params",
		NTType:     5,
		Index:      7,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Params : Type id	<<  >>`,
		Id:         "Params",
		NTType:     5,
		Index:      8,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Params : empty	<<  >>`,
		Id:         "Params",
		NTType:     5,
		Index:      9,
		NumSymbols: 0,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return nil, nil
		},
	},
	ProdTabEntry{
		String: `Type : "num"	<<  >>`,
		Id:         "Type",
		NTType:     6,
		Index:      10,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Type : "bool"	<<  >>`,
		Id:         "Type",
		NTType:     6,
		Index:      11,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Type : "char"	<<  >>`,
		Id:         "Type",
		NTType:     6,
		Index:      12,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Type : "[" Type "]"	<<  >>`,
		Id:         "Type",
		NTType:     6,
		Index:      13,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Statement : FunctionCall	<<  >>`,
		Id:         "Statement",
		NTType:     7,
		Index:      14,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Statement : Lambda	<<  >>`,
		Id:         "Statement",
		NTType:     7,
		Index:      15,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Statement : Expression	<<  >>`,
		Id:         "Statement",
		NTType:     7,
		Index:      16,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Statement : Constant	<<  >>`,
		Id:         "Statement",
		NTType:     7,
		Index:      17,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Statement : id	<<  >>`,
		Id:         "Statement",
		NTType:     7,
		Index:      18,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Expression : Expression plusminus Term	<<  >>`,
		Id:         "Expression",
		NTType:     8,
		Index:      19,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Expression : Term	<<  >>`,
		Id:         "Expression",
		NTType:     8,
		Index:      20,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Term : Term multdiv Factor	<<  >>`,
		Id:         "Term",
		NTType:     9,
		Index:      21,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Term : Factor	<<  >>`,
		Id:         "Term",
		NTType:     9,
		Index:      22,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Factor : "(" Expression ")"	<<  >>`,
		Id:         "Factor",
		NTType:     10,
		Index:      23,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Factor : Statement	<<  >>`,
		Id:         "Factor",
		NTType:     10,
		Index:      24,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `FunctionCall : id "(" Args ")"	<<  >>`,
		Id:         "FunctionCall",
		NTType:     11,
		Index:      25,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Lambda : "(" "\\" Params "=>" Statement ")"	<<  >>`,
		Id:         "Lambda",
		NTType:     12,
		Index:      26,
		NumSymbols: 6,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Args : Value "," Args	<<  >>`,
		Id:         "Args",
		NTType:     13,
		Index:      27,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Args : Value	<<  >>`,
		Id:         "Args",
		NTType:     13,
		Index:      28,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Args : empty	<<  >>`,
		Id:         "Args",
		NTType:     13,
		Index:      29,
		NumSymbols: 0,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return nil, nil
		},
	},
	ProdTabEntry{
		String: `Value : Constant	<<  >>`,
		Id:         "Value",
		NTType:     14,
		Index:      30,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Value : Statement	<<  >>`,
		Id:         "Value",
		NTType:     14,
		Index:      31,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Constant : boolean	<<  >>`,
		Id:         "Constant",
		NTType:     15,
		Index:      32,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Constant : number	<<  >>`,
		Id:         "Constant",
		NTType:     15,
		Index:      33,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Constant : charac	<<  >>`,
		Id:         "Constant",
		NTType:     15,
		Index:      34,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Constant : "[" Args "]"	<<  >>`,
		Id:         "Constant",
		NTType:     15,
		Index:      35,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
}
