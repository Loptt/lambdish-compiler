package ic

import (
	"sort"

	"github.com/Loptt/lambdish-compiler/ast"
	"github.com/Loptt/lambdish-compiler/dir"
	"github.com/mewkiz/pkg/errutil"
)

// TODO: Implement check to figure out if the variable space has been filled and return error

func generateAddressesProgram(program *ast.Program, ctx *GenerationContext) error {
	for _, fe := range ctx.funcdir.Table() {
		if err := generateAddressesFuncEntry(fe, ctx); err != nil {
			return err
		}
	}

	for _, f := range program.Functions() {
		if err := generateAddressesFunction(f, ctx); err != nil {
			return err
		}
	}

	if err := generateAddressesFunctionCall(program.Call(), ctx); err != nil {
		return err
	}

	return nil
}

func generateAddressesFuncEntry(fe *dir.FuncEntry, ctx *GenerationContext) error {
	ctx.vm.ResetLocal()

	// Before we assign an address to each parameter of the function, we need to sort the
	// entries by their position in the definition. This will keep the order in runtime
	ves := make([]*dir.VarEntry, 0)

	// We extract every single var entry
	for _, ve := range fe.VarDir().Table() {
		ves = append(ves, ve)
	}

	// We sort them by their position in the function declaration
	sort.SliceStable(ves, func(i, j int) bool {
		return ves[i].Pos() > ves[j].Pos()
	})

	// Now we can request address assigning using the sorted array which
	// will guarantee that the addresses are given in order
	for _, ve := range ves {
		addr, err := ctx.vm.GetNextLocal(ve.Type())
		if err != nil {
			return err
		}
		ve.SetAddress(addr)
		fe.SetEra(fe.Era() + 1)
	}

	for _, l := range fe.Lambdas() {
		if err := generateAddressesFuncEntry(l, ctx); err != nil {
			return err
		}
	}

	return nil
}

func generateAddressesFunction(function *ast.Function, ctx *GenerationContext) error {
	if err := generateAddressesStatement(function.Statement(), ctx); err != nil {
		return err
	}

	return nil
}

func generateAddressesStatement(statement ast.Statement, ctx *GenerationContext) error {
	if cv, ok := statement.(*ast.ConstantValue); ok {
		if err := generateAddressesConstantValue(cv, ctx); err != nil {
			return err
		}
		return nil
	} else if cl, ok := statement.(*ast.ConstantList); ok {
		if err := generateAddressesConstantList(cl, ctx); err != nil {
			return err
		}
		return nil
	} else if fcall, ok := statement.(*ast.FunctionCall); ok {
		if err := generateAddressesFunctionCall(fcall, ctx); err != nil {
			return err
		}
		return nil
	} else if l, ok := statement.(*ast.Lambda); ok {
		if err := generateAddressesStatement(l.Statement(), ctx); err != nil {
			return err
		}
		return nil
	} else if _, ok := statement.(*ast.Id); ok {
		return nil
	}

	return errutil.Newf("Cannot cast statement to valid form")
}

func generateAddressesConstantValue(cv *ast.ConstantValue, ctx *GenerationContext) error {
	if !ctx.vm.ConstantExists(cv.Value()) {
		ctx.vm.AddConstant(cv.Value(), cv.Type())
	}

	return nil
}

func generateAddressesFunctionCall(fcall *ast.FunctionCall, ctx *GenerationContext) error {
	if err := generateAddressesStatement(fcall.Statement(), ctx); err != nil {
		return err
	}

	for _, args := range fcall.Args() {
		if err := generateAddressesStatement(args, ctx); err != nil {
			return err
		}
	}

	return nil
}

func generateAddressesConstantList(cl *ast.ConstantList, ctx *GenerationContext) error {
	for _, arg := range cl.Contents() {
		if err := generateAddressesStatement(arg, ctx); err != nil {
			return err
		}
	}
	return nil
}
