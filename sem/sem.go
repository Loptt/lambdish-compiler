package sem

import (
	"github.com/Loptt/lambdish-compiler/ast"
	"github.com/Loptt/lambdish-compiler/dir"
)

// SemanticCheck: Construcción
func SemanticCheck(program *ast.Program) (*dir.FuncDirectory, error) {
	funcdir := dir.NewFuncDirectory()

	err := buildFuncDirProgram(program, funcdir)
	if err != nil {
		return nil, err
	}

	return funcdir, nil
}

// Checar que las funciones y variables que se usen existan

// Checar la cohesion de tipos

/*
(*ast.Program)(0xc0000a0580)({
 functions: ([]*ast.Function) (len=1 cap=1) {
  (*ast.Function)(0xc00009f140)({
   id: (string) (len=3) "map",
   params: ([]*dir.VarEntry) (len=2 cap=2) {
    (*dir.VarEntry)(0xc0000a01e0)(l),
    (*dir.VarEntry)(0xc0000a01a0)(f)
   },
   t: (*types.LambdishType)(0xc000088660)([1]),
   statement: (*ast.FunctionCall)(0xc000088750)({
    id: (string) (len=6) "append",
    args: ([]ast.Statement) (len=2 cap=2) {
     (*ast.FunctionCall)(0xc0000886c0)({
      id: (string) (len=1) "f",
      args: ([]ast.Statement) (len=1 cap=1) {
       (*ast.FunctionCall)(0xc000088690)({
        id: (string) (len=4) "head",
        args: ([]ast.Statement) (len=1 cap=1) {
         (*ast.Id)(0xc000084540)((len=1) "l")
        }
       })
      }
     }),
     (*ast.FunctionCall)(0xc000088720)({
      id: (string) (len=3) "map",
      args: ([]ast.Statement) (len=2 cap=2) {
       (*ast.FunctionCall)(0xc0000886f0)({
        id: (string) (len=4) "tail",
        args: ([]ast.Statement) (len=1 cap=1) {
         (*ast.Id)(0xc000084570)((len=1) "l")
        }
       }),
       (*ast.Id)(0xc000084590)((len=1) "f")
      }
     })
    }
   })
  })
 },
 call: (*ast.FunctionCall)(0xc000088930)({
  id: (string) (len=3) "map",
  args: ([]ast.Statement) (len=2 cap=2) {
   (*ast.ConstantList)(0xc0000a0440)({
    contents: ([]ast.Statement) (len=3 cap=3) {
     (*ast.ConstantValue)(0xc0000a0360)({
      t: (*types.LambdishType)(0xc000088780)(1),
      value: (string) (len=1) "1"
     }),
     (*ast.ConstantValue)(0xc0000a0380)({
      t: (*types.LambdishType)(0xc0000887b0)(1),
      value: (string) (len=1) "2"
     }),
     (*ast.ConstantValue)(0xc0000a03a0)({
      t: (*types.LambdishType)(0xc0000887e0)(1),
      value: (string) (len=1) "3"
     })
    }
   }),
   (*ast.Lambda)(0xc000088900)({
    params: ([]*dir.VarEntry) (len=1 cap=1) {
     (*dir.VarEntry)(0xc0000a0460)(x)
    },
    statement: (*ast.FunctionCall)(0xc0000888d0)({
     id: (string) (len=1) "+",
     args: ([]ast.Statement) (len=2 cap=2) {
      (*ast.Id)(0xc000084610)((len=1) "x"),
      (*ast.ConstantValue)(0xc0000a04a0)({
       t: (*types.LambdishType)(0xc0000888a0)(1),
       value: (string) (len=1) "1"
      })
     }
    }),
    retval: (*types.LambdishType)(0xc000088870)(1)
   })
  }
 })
})
*/