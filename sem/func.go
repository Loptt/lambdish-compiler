package sem

import (
	"github.com/Loptt/lambdish-compiler/ast"
    "github.com/Loptt/lambdish-compiler/dir"
    "github.com/Loptt/lambdish-compiler/types"
	"github.com/mewkiz/pkg/errutil"
)

// The functions on this file traverse the ast depth first and construct the function directory
// for the declared functions.

func buildFuncDirProgram(program *ast.Program, funcdir *dir.FuncDirectory) error {
	for _, f := range program.Functions() {
		
		err := buildFuncDirFunction(f, funcdir)
		if err != nil {
			return err
		}
    }
    return nil
}

func buildFuncDirFunction(function *ast.Function, funcdir *dir.FuncDirectory) error {

	id := function.Id()
	t := function.Type()
	vardir := dir.NewVarDirectory()
	params := make([]*types.LambdishType, 0)

	for _, p := range function.Params() {
		params = append(params, p.Type())
		errvd := buildVarDirFunction(p, vardir)
		if errvd != nil {
			return errvd
		}
	}

	fe := dir.NewFuncEntry(id, t, len(params), params, vardir)

	ok := funcdir.Add(fe)
	if !ok {
		return errutil.Newf("Invalid FuncEntry. This FuncEntry already exists.")
    }
    return nil
}

func buildVarDirFunction(ve *dir.VarEntry, vardir *dir.VarDirectory) error {

    ok := vardir.Add(ve)
    if !ok{
        return errutil.Newf("Invalid VarEntry. This VarEntry already exists.")
    }
    return nil
}

/*
EXAMPLE AST
(**ast.Program)(0xc000084520->0xc0000a0740)({
 functions: ([]*ast.Function) (len=2 cap=2) {
  (*ast.Function)(0xc00009fac0)({
   id: (string) (len=3) "max",
   params: ([]*dir.VarEntry) (len=1 cap=1) {
    (*dir.VarEntry)(0xc0000a04a0)(l)
   },
   t: (*types.LambdishType)(0xc0000982a0)(1),
   statement: (*ast.FunctionCall)(0xc0000886f0)({
    id: (string) (len=6) "maxAux",
    args: ([]ast.Statement) (len=2 cap=2) {
     (*ast.FunctionCall)(0xc0000886c0)({
      id: (string) (len=4) "head",
      args: ([]ast.Statement) (len=1 cap=1) {
       (*ast.Id)(0xc000084670)((len=1) "l")
      }
     }),
     (*ast.Id)(0xc000084660)((len=1) "l")
    }
   })
  }),
  (*ast.Function)(0xc00009f580)({
   id: (string) (len=6) "maxAux",
   params: ([]*dir.VarEntry) (len=2 cap=2) {
    (*dir.VarEntry)(0xc0000a0160)(biggest),
    (*dir.VarEntry)(0xc0000a01a0)(l)
   },
   t: (*types.LambdishType)(0xc000098220)(1),
   statement: (*ast.FunctionCall)(0xc000088690)({
    id: (string) (len=2) "if",
    args: ([]ast.Statement) (len=3 cap=4) {
     (*ast.FunctionCall)(0xc000088660)({
      id: (string) (len=2) "if",
      args: ([]ast.Statement) (len=3 cap=4) {
       (*ast.FunctionCall)(0xc000088630)({
        id: (string) (len=6) "maxAux",
        args: ([]ast.Statement) (len=2 cap=2) {
         (*ast.Id)(0xc000084620)((len=7) "biggest"),
         (*ast.FunctionCall)(0xc000088600)({
          id: (string) (len=4) "tail",
          args: ([]ast.Statement) (len=1 cap=1) {
           (*ast.Id)(0xc000084600)((len=1) "l")
          }
         })
        }
       }),
       (*ast.FunctionCall)(0xc0000885d0)({
        id: (string) (len=6) "maxAux",
        args: ([]ast.Statement) (len=2 cap=2) {
         (*ast.FunctionCall)(0xc0000885a0)({
          id: (string) (len=4) "head",
          args: ([]ast.Statement) (len=1 cap=1) {
           (*ast.Id)(0xc0000845d0)((len=1) "l")
          }
         }),
         (*ast.FunctionCall)(0xc000088570)({
          id: (string) (len=4) "tail",
          args: ([]ast.Statement) (len=1 cap=1) {
           (*ast.Id)(0xc0000845b0)((len=1) "l")
          }
         })
        }
       }),
       (*ast.FunctionCall)(0xc000088540)({
        id: (string) (len=1) ">",
        args: ([]ast.Statement) (len=2 cap=2) {
         (*ast.Id)(0xc000084590)((len=7) "biggest"),
         (*ast.FunctionCall)(0xc000088510)({
          id: (string) (len=4) "head",
          args: ([]ast.Statement) (len=1 cap=1) {
           (*ast.Id)(0xc000084570)((len=1) "l")
          }
         })
        }
       })
      }
     }),
     (*ast.Id)(0xc000084560)((len=7) "biggest"),
     (*ast.FunctionCall)(0xc0000884e0)({
      id: (string) (len=5) "empty",
      args: ([]ast.Statement) (len=1 cap=1) {
       (*ast.Id)(0xc000084540)((len=1) "l")
      }
     })
    }
   })
  })
 },
 call: (*ast.FunctionCall)(0xc000088720)({
  id: (string) (len=3) "max",
  args: ([]ast.Statement) (len=1 cap=1) {
   (*ast.ConstantList)(0xc0000a0700)({
    contents: ([]ast.Statement) (len=5 cap=8) {
     (*ast.ConstantValue)(0xc0000a0620)({
      t: (*types.LambdishType)(0xc000098300)(1),
      value: (string) (len=1) "3"
     }),
     (*ast.ConstantValue)(0xc0000a0600)({
      t: (*types.LambdishType)(0xc0000982f0)(1),
      value: (string) (len=1) "2"
     }),
     (*ast.ConstantValue)(0xc0000a05e0)({
      t: (*types.LambdishType)(0xc0000982e0)(1),
      value: (string) (len=1) "4"
     }),
     (*ast.ConstantValue)(0xc0000a05c0)({
      t: (*types.LambdishType)(0xc0000982d0)(1),
      value: (string) (len=1) "3"
     }),
     (*ast.ConstantValue)(0xc0000a05a0)({
      t: (*types.LambdishType)(0xc0000982c0)(1),
      value: (string) (len=1) "2"
     })
    }
   })
  }
 })
})
*/
