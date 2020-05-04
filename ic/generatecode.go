package ic

import (
	"github.com/Loptt/lambdish-compiler/ast"
	"github.com/Loptt/lambdish-compiler/dir"
	"github.com/Loptt/lambdish-compiler/mem"
	"github.com/Loptt/lambdish-compiler/sem"
	"github.com/mewkiz/pkg/errutil"
)

func generateCodeProgram(program *ast.Program, ctx *GenerationContext) error {
	for _, function := range program.Functions() {
		if err := generateCodeFunction(function, ctx); err != nil {
			return err
		}
	}

	return nil
}

func generateCodeFunction(function *ast.Function, ctx *GenerationContext) error {
	fe := ctx.funcdir.Get(function.Key())

	fes := dir.NewFuncEntryStack()
	fes.Push(fe)

	if err := generateCodeStatement(function.Statement(), fes, ctx); err != nil {
		return err
	}

	return nil
}

func generateCodeStatement(statement ast.Statement, fes *dir.FuncEntryStack, ctx *GenerationContext) error {
	if _, ok := statement.(*ast.Id); ok {
		// TOOD: Implement code generation for Id
		return nil
	} else if fcall, ok := statement.(*ast.FunctionCall); ok {
		return generateCodeFunctionCall(fcall, fes, ctx)
	} else if _, ok := statement.(*ast.Lambda); ok {
		// TOOD: Implement code generation for Id
		return nil
	} else if _, ok := statement.(*ast.ConstantList); ok {
		// TOOD: Implement code generation for Id
		return nil
	} else if _, ok := statement.(*ast.ConstantValue); ok {
		// TOOD: Implement code generation for Id
		return nil
	}

	return errutil.Newf("Statement cannot be casted to any valid form")
}

func generateCodeFunctionCall(fcall *ast.FunctionCall, fes *dir.FuncEntryStack, ctx *GenerationContext) error {
	if id, ok := fcall.Statement().(*ast.Id); ok {
		if sem.IsReservedFunction(id.String()) {
			if err := generateCodeReservedFunctionCall(id, fcall, fes, ctx); err != nil {
				return err
			}
		} else {
			// TODO: Generate code with user defined function calls
			return nil
		}
	} else {
		// TODO: Generate code when function call is not Id
		return nil
	}

	return nil
}

func generateCodeReservedFunctionCall(id *ast.Id, fcall *ast.FunctionCall, fes *dir.FuncEntryStack, ctx *GenerationContext) error {
	switch id.String() {
	case "+", "-", "/", "*", "%":
		if err := generateArithmeticalOperators(id.String(), fcall, fes, ctx); err != nil {
			return err
		}
	case "<", ">", "equal":
		if err := generateRelationalOperators(id.String(), fcall, fes, ctx); err != nil {
			return err
		}
	case "and", "or", "!":
		if err := generateLogicalOperators(id.String(), fcall, fes, ctx); err != nil {
			return err
		}
	case "if":
		if err := generateIf(fcall, fes, ctx); err != nil {
			return err
		}
	}
	return nil
}

func generateArithmeticalOperators(id string, fcall *ast.FunctionCall, fes *dir.FuncEntryStack, ctx *GenerationContext) error {
	args := fcall.Args()

	// First we get the left and right operands from the args array
	// TODO: Check if they are in the right order to ensure left associativity
	lop := args[0]
	rop := args[1]

	//Receive
	op := GetOperation(id)
	if op == Invalid {
		return errutil.Newf("%+v: Cannot generate for arithmetical operator %s", fcall.Token(), id)
	}

	laddr, err := getArgumentAddress(lop, fes, ctx)
	if err != nil {
		return err
	}

	raddr, err := getArgumentAddress(rop, fes, ctx)
	if err != nil {
		return err
	}

	// Get the address of the next available temp to store the result of the operation
	nextTemp := ctx.vm.GetNextTemp()

	// Generate the quadruple
	ctx.gen.Generate(op, laddr, raddr, nextTemp)

	// Push the result of the quadruple to the address stack
	ctx.gen.PushToAddrStack(nextTemp)
	return nil
}

func generateRelationalOperators(id string, fcall *ast.FunctionCall, fes *dir.FuncEntryStack, ctx *GenerationContext) error {
	args := fcall.Args()

	// First we get the left and right operands from the args array
	// TODO: Check if they are in the right order to ensure left associativity
	lop := args[0]
	rop := args[1]

	//Receive
	op := GetOperation(id)
	if op == Invalid {
		return errutil.Newf("%+v: Cannot generate for arithmetical operator %s", fcall.Token(), id)
	}

	laddr, err := getArgumentAddress(lop, fes, ctx)
	if err != nil {
		return err
	}

	raddr, err := getArgumentAddress(rop, fes, ctx)
	if err != nil {
		return err
	}

	// Get the address of the next available temp to store the result of the operation
	nextTemp := ctx.vm.GetNextTemp()

	// Generate the quadruple
	ctx.gen.Generate(op, laddr, raddr, nextTemp)

	// Push the result of the quadruple to the address stack
	ctx.gen.PushToAddrStack(nextTemp)
	return nil
}

func generateLogicalOperators(id string, fcall *ast.FunctionCall, fes *dir.FuncEntryStack, ctx *GenerationContext) error {
	args := fcall.Args()

	if len(args) == 1 {
		lop := args[0]

		//Receive
		op := GetOperation(id)
		if op == Invalid {
			return errutil.Newf("%+v: Cannot generate for arithmetical operator %s", fcall.Token(), id)
		}

		laddr, err := getArgumentAddress(lop, fes, ctx)
		if err != nil {
			return err
		}
		nextTemp := ctx.vm.GetNextTemp()

		// Generate the quadruple
		ctx.gen.Generate(op, laddr, mem.Address(-1), nextTemp)

		// Push the result of the quadruple to the address stack
		ctx.gen.PushToAddrStack(nextTemp)
		return nil
	}
	// First we get the left and right operands from the args array
	// TODO: Check if they are in the right order to ensure left associativity
	lop := args[0]
	rop := args[1]

	//Receive
	op := GetOperation(id)
	if op == Invalid {
		return errutil.Newf("%+v: Cannot generate for arithmetical operator %s", fcall.Token(), id)
	}

	laddr, err := getArgumentAddress(lop, fes, ctx)
	if err != nil {
		return err
	}

	raddr, err := getArgumentAddress(rop, fes, ctx)
	if err != nil {
		return err
	}

	// Get the address of the next available temp to store the result of the operation
	nextTemp := ctx.vm.GetNextTemp()

	// Generate the quadruple
	ctx.gen.Generate(op, laddr, raddr, nextTemp)

	// Push the result of the quadruple to the address stack
	ctx.gen.PushToAddrStack(nextTemp)
	return nil
}

func generateIf(fcall *ast.FunctionCall, fes *dir.FuncEntryStack, ctx *GenerationContext) error {
	args := fcall.Args()

	caddr, err := getArgumentAddress(args[0], fes, ctx)
	if err != nil {
		return err
	}

	ctx.gen.PushToJumpStack(mem.Address(ctx.gen.Counter()))
	ctx.gen.Generate(GotoF, caddr, mem.Address(-1), mem.Address(-1))

	laddr, err := getArgumentAddress(args[1], fes, ctx)
	if err != nil {
		return err
	}

	fjump := ctx.gen.GetFromJumpStack()
	ctx.gen.Generate(Ret, laddr, mem.Address(-1), mem.Address(-1))
	ctx.gen.FillJumpQuadruple(fjump, mem.Address(ctx.gen.Counter()))

	raddr, err := getArgumentAddress(args[2], fes, ctx)
	if err != nil {
		return err
	}

	ctx.gen.Generate(Ret, raddr, mem.Address(-1), mem.Address(-1))

	return nil
}

func getArgumentAddress(s ast.Statement, fes *dir.FuncEntryStack, ctx *GenerationContext) (mem.Address, error) {
	if id, ok := s.(*ast.Id); ok {
		if addr, ok := getAddressFromFuncStack(id, fes); ok {
			return addr, nil
		}
		// TODO: Check if it refers to a function in the global scope
		return mem.Address(-1), errutil.Newf("%+v: id %s not found in this scope", s.Token(), id.String())

	} else if fcall, ok := s.(*ast.FunctionCall); ok {
		if err := generateCodeFunctionCall(fcall, fes, ctx); err != nil {
			return mem.Address(-1), nil
		}
		return ctx.gen.GetFromAddrStack(), nil
	} else if cv, ok := s.(*ast.ConstantValue); ok {
		return ctx.vm.GetConstantAddress(cv.Value()), nil
	}
	//TODO: Generate code for regular function call

	return mem.Address(-1), nil
}
