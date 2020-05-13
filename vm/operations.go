package vm

import (
	"github.com/Loptt/lambdish-compiler/mem"
	"github.com/mewkiz/pkg/errutil"
)

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
