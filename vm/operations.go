package vm

import (
	"math"

	"github.com/Loptt/lambdish-compiler/mem"
	"github.com/mewkiz/pkg/errutil"
)

func (vm *VirtualMachine) operationAdd(lop, rop, r mem.Address) error {
	lopv, err := vm.mm.GetValue(lop)
	if err != nil {
		return err
	}

	ropv, err := vm.mm.GetValue(rop)
	if err != nil {
		return err
	}

	f1, err := getNum(lopv)
	if err != nil {
		return err
	}

	f2, err := getNum(ropv)
	if err != nil {
		return err
	}

	result := f1 + f2

	if err := vm.mm.SetValue(result, r); err != nil {
		return err
	}

	return nil
}

func (vm *VirtualMachine) operationSub(lop, rop, r mem.Address) error {
	lopv, err := vm.mm.GetValue(lop)
	if err != nil {
		return err
	}

	ropv, err := vm.mm.GetValue(rop)
	if err != nil {
		return err
	}

	f1, err := getNum(lopv)
	if err != nil {
		return err
	}

	f2, err := getNum(ropv)
	if err != nil {
		return err
	}

	result := f1 - f2

	if err := vm.mm.SetValue(result, r); err != nil {
		return err
	}

	return nil
}

func (vm *VirtualMachine) operationMult(lop, rop, r mem.Address) error {
	lopv, err := vm.mm.GetValue(lop)
	if err != nil {
		return err
	}

	ropv, err := vm.mm.GetValue(rop)
	if err != nil {
		return err
	}

	f1, err := getNum(lopv)
	if err != nil {
		return err
	}

	f2, err := getNum(ropv)
	if err != nil {
		return err
	}

	result := f1 * f2

	if err := vm.mm.SetValue(result, r); err != nil {
		return err
	}

	return nil
}

func (vm *VirtualMachine) operationDiv(lop, rop, r mem.Address) error {
	lopv, err := vm.mm.GetValue(lop)
	if err != nil {
		return err
	}

	ropv, err := vm.mm.GetValue(rop)
	if err != nil {
		return err
	}

	f1, err := getNum(lopv)
	if err != nil {
		return err
	}

	f2, err := getNum(ropv)
	if err != nil {
		return err
	}

	if f2 == 0 {
		return errutil.Newf("Arithmethic exception, division by 0")
	}

	result := f1 / f2

	if err := vm.mm.SetValue(result, r); err != nil {
		return err
	}

	return nil
}

func (vm *VirtualMachine) operationMod(lop, rop, r mem.Address) error {
	lopv, err := vm.mm.GetValue(lop)
	if err != nil {
		return err
	}

	ropv, err := vm.mm.GetValue(rop)
	if err != nil {
		return err
	}

	f1, err := getNum(lopv)
	if err != nil {
		return err
	}

	f2, err := getNum(ropv)
	if err != nil {
		return err
	}

	if f2 == 0 {
		return errutil.Newf("Arithmethic exception, division by 0")
	}

	result := math.Mod(f1, f2)

	if err := vm.mm.SetValue(result, r); err != nil {
		return err
	}

	return nil
}

func (vm *VirtualMachine) operationAnd(lop, rop, r mem.Address) error {
	lopv, err := vm.mm.GetValue(lop)
	if err != nil {
		return err
	}

	ropv, err := vm.mm.GetValue(rop)
	if err != nil {
		return err
	}

	b1, err := getBool(lopv)
	if err != nil {
		return err
	}

	b2, err := getBool(ropv)
	if err != nil {
		return err
	}

	result := b1 && b2

	if err := vm.mm.SetValue(result, r); err != nil {
		return err
	}

	return nil
}

func (vm *VirtualMachine) operationOr(lop, rop, r mem.Address) error {
	lopv, err := vm.mm.GetValue(lop)
	if err != nil {
		return err
	}

	ropv, err := vm.mm.GetValue(rop)
	if err != nil {
		return err
	}

	b1, err := getBool(lopv)
	if err != nil {
		return err
	}

	b2, err := getBool(ropv)
	if err != nil {
		return err
	}

	result := b1 || b2

	if err := vm.mm.SetValue(result, r); err != nil {
		return err
	}

	return nil
}

func (vm *VirtualMachine) operationNot(lop, rop, r mem.Address) error {
	lopv, err := vm.mm.GetValue(lop)
	if err != nil {
		return err
	}

	b1, err := getBool(lopv)
	if err != nil {
		return err
	}

	result := !b1

	if err := vm.mm.SetValue(result, r); err != nil {
		return err
	}

	return nil
}

func (vm *VirtualMachine) operationPrint(lop, rop, r mem.Address) error {
	result, err := vm.mm.GetValue(r)
	if err != nil {
		return err
	}

	vm.output = result

	return nil
}

func (vm *VirtualMachine) operationGoto(lop, rop, r mem.Address) error {
	if lop < 0 {
		return errutil.Newf("Invalid instruction address")
	}

	vm.ip = int(lop)

	return nil
}
