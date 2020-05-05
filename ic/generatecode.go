package ic

import (
	"github.com/Loptt/lambdish-compiler/ast"
	"github.com/Loptt/lambdish-compiler/dir"
	"github.com/Loptt/lambdish-compiler/mem"
	"github.com/Loptt/lambdish-compiler/sem"
	"github.com/mewkiz/pkg/errutil"
)

func generateCodeProgram(program *ast.Program, ctx *GenerationContext) error {

	ctx.gen.AddPendingFuncAddr(ctx.gen.ICounter(), "main")
	ctx.gen.Generate(Goto, mem.Address(-1), mem.Address(-1), mem.Address(-1))

	for _, function := range program.Functions() {
		if err := generateCodeFunction(function, ctx); err != nil {
			return err
		}
	}

	fes := dir.NewFuncEntryStack()
	fes.Push(ctx.funcdir.Get("main"))

	ctx.funcdir.Get("main").SetLocation(ctx.gen.ICounter())

	if err := generateCodeFunctionCall(program.Call(), fes, ctx); err != nil {
		return err
	}

	ctx.gen.FillPendingFuncAddr(ctx.funcdir)

	return nil
}

func generateCodeFunction(function *ast.Function, ctx *GenerationContext) error {

	ctx.vm.ResetTemp()

	fe := ctx.funcdir.Get(function.Key())

	fes := dir.NewFuncEntryStack()
	fes.Push(fe)

	fe.SetLocation(ctx.gen.ICounter())

	if err := generateCodeStatement(function.Statement(), fes, ctx); err != nil {
		return err
	}

	addr := ctx.gen.GetFromAddrStack()

	ctx.gen.Generate(Ret, addr, mem.Address(-1), mem.Address(-1))

	return nil
}

func generateCodeStatement(statement ast.Statement, fes *dir.FuncEntryStack, ctx *GenerationContext) error {
	if id, ok := statement.(*ast.Id); ok {
		return generateCodeId(id, fes, ctx)
	} else if fcall, ok := statement.(*ast.FunctionCall); ok {
		return generateCodeFunctionCall(fcall, fes, ctx)
	} else if lambda, ok := statement.(*ast.Lambda); ok {
		return generateCodeLambda(lambda, fes, ctx)
	} else if _, ok := statement.(*ast.ConstantList); ok {
		// TODO: Implement code generation for Id
		return nil
	} else if _, ok := statement.(*ast.ConstantValue); ok {
		// TODO: Implement code generation for Id
		return nil
	}

	return errutil.Newf("Statement cannot be casted to any valid form")
}

func generateCodeId(id *ast.Id, fes *dir.FuncEntryStack, ctx *GenerationContext) error {
	if addr, ok := getAddressFromFuncStack(id, fes); ok {
		ctx.gen.PushToAddrStack(addr)
		return nil
	} else if fe := ctx.funcdir.Get(id.String()); fe != nil {
		ctx.gen.AddPendingFuncAddr(ctx.gen.ICounter(), id.String())
		ctx.gen.PushToAddrStack(mem.Address(-1))
		return nil
	}

	return errutil.Newf("%+v: Cannot find id %s in local or global scope", id.Token(), id.String())
}

func generateCodeFunctionCall(fcall *ast.FunctionCall, fes *dir.FuncEntryStack, ctx *GenerationContext) error {
	if id, ok := fcall.Statement().(*ast.Id); ok {
		if sem.IsReservedFunction(id.String()) {
			if err := generateCodeReservedFunctionCall(id, fcall, fes, ctx); err != nil {
				return err
			}
		} else {
			// First we generate the ERA operation
			// TODO: Change so that arg or ERA is the size of the call
			ctx.gen.Generate(Era, mem.Address(-1), mem.Address(-1), mem.Address(-1))

			pcounter := 0

			// For each argument we get its address, which will automatically generate
			// the necesary code to resolve each argument
			for _, arg := range fcall.Args() {
				addr, err := getArgumentAddress(arg, fes, ctx)
				if err != nil {
					return err
				}
				ctx.gen.Generate(Param, addr, mem.Address(-1), mem.Address(pcounter))
				pcounter++
			}

			if err := generateCodeId(id, fes, ctx); err != nil {
				return err
			}

			tmp := ctx.vm.GetNextTemp()
			calladdr := ctx.gen.GetFromAddrStack()

			ctx.gen.Generate(Call, calladdr, mem.Address(-1), tmp)
			ctx.gen.PushToAddrStack(tmp)

			return nil
		}
	} else if l, ok := fcall.Statement().(*ast.Lambda); ok {
		if err := generateCodeLambda(l, fes, ctx); err != nil {
			return err
		}

		lambdaaddr := ctx.gen.GetFromAddrStack()

		ctx.gen.Generate(Era, mem.Address(-1), mem.Address(-1), mem.Address(-1))
		pcounter := 0

		// For each argument we get its address, which will automatically generate
		// the necesary code to resolve each argument
		for _, arg := range fcall.Args() {
			addr, err := getArgumentAddress(arg, fes, ctx)
			if err != nil {
				return err
			}
			ctx.gen.Generate(Param, addr, mem.Address(-1), mem.Address(pcounter))
			pcounter++
		}

		tmp := ctx.vm.GetNextTemp()
		ctx.gen.Generate(Call, lambdaaddr, mem.Address(-1), tmp)
		ctx.gen.PushToAddrStack(tmp)
	}
	return nil
}

func generateCodeLambda(lambda *ast.Lambda, fes *dir.FuncEntryStack, ctx *GenerationContext) error {

	fe := fes.Top().GetLambdaEntryById(lambda.Id())
	fes.Push(fe)

	// If we define a lambda, we need to add a goto to prevent the flow from executing
	// the lambda code without being explicitely called
	ctx.gen.PushToJumpStack(mem.Address(ctx.gen.ICounter()))
	ctx.gen.Generate(Goto, mem.Address(-1), mem.Address(-1), mem.Address(-1))

	// After we add the goto, we set the start of the lambda to the current icounter
	fe.SetLocation(ctx.gen.ICounter())

	if err := generateCodeStatement(lambda.Statement(), fes, ctx); err != nil {
		return err
	}
	addr := ctx.gen.GetFromAddrStack()

	ctx.gen.Generate(Ret, addr, mem.Address(-1), mem.Address(-1))

	// Once the lambda has been generated, we fill the pending goto with the
	// current icounter
	jump := ctx.gen.GetFromJumpStack()
	ctx.gen.FillJumpQuadruple(jump, mem.Address(ctx.gen.ICounter()))

	ctx.gen.PushToAddrStack(fe.Loc())

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
	case "empty", "head", "tail":
		if err := generateBuiltInOneArg(id.String(), fcall, fes, ctx); err != nil {
			return err
		}
	case "append", "insert":
		if err := generateBuiltInTwoArgs(id.String(), fcall, fes, ctx); err != nil {
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

	ctx.gen.PushToJumpStack(mem.Address(ctx.gen.ICounter()))
	ctx.gen.Generate(GotoF, caddr, mem.Address(-1), mem.Address(-1))

	laddr, err := getArgumentAddress(args[1], fes, ctx)
	if err != nil {
		return err
	}

	fjump := ctx.gen.GetFromJumpStack()
	ctx.gen.Generate(Ret, laddr, mem.Address(-1), mem.Address(-1))
	ctx.gen.FillJumpQuadruple(fjump, mem.Address(ctx.gen.ICounter()))

	raddr, err := getArgumentAddress(args[2], fes, ctx)
	if err != nil {
		return err
	}

	ctx.gen.PushToAddrStack(raddr)

	return nil
}

func generateBuiltInOneArg(id string, fcall *ast.FunctionCall, fes *dir.FuncEntryStack, ctx *GenerationContext) error {
	arg := fcall.Args()[0]

	addr, err := getArgumentAddress(arg, fes, ctx)
	if err != nil {
		return err
	}

	op := GetOperation(id)
	tmp := ctx.vm.GetNextTemp()

	ctx.gen.Generate(op, addr, mem.Address(-1), tmp)

	ctx.gen.PushToAddrStack(tmp)

	return nil
}

func generateBuiltInTwoArgs(id string, fcall *ast.FunctionCall, fes *dir.FuncEntryStack, ctx *GenerationContext) error {
	args := fcall.Args()

	laddr, err := getArgumentAddress(args[0], fes, ctx)
	if err != nil {
		return err
	}

	raddr, err := getArgumentAddress(args[1], fes, ctx)
	if err != nil {
		return err
	}

	op := GetOperation(id)
	tmp := ctx.vm.GetNextTemp()

	ctx.gen.Generate(op, laddr, raddr, tmp)

	ctx.gen.PushToAddrStack(tmp)

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
	} else if l, ok := s.(*ast.Lambda); ok {
		if err := generateCodeLambda(l, fes, ctx); err != nil {
			return mem.Address(-1), err
		}
		return ctx.gen.GetFromAddrStack(), nil
	}
	//TODO: Generate code for regular function call

	return mem.Address(-1), nil
}
