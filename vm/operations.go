package vm

import (
	"math"

	"github.com/Loptt/lambdish-compiler/mem"
	"github.com/Loptt/lambdish-compiler/vm/ar"
	"github.com/Loptt/lambdish-compiler/vm/list"
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
		return errutil.NewNoPosf("Arithmethic exception, division by 0")
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
		return errutil.NewNoPosf("Arithmethic exception, division by 0")
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

func (vm *VirtualMachine) operationGt(lop, rop, r mem.Address) error {
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

	result := f1 > f2

	if err := vm.mm.SetValue(result, r); err != nil {
		return err
	}

	return nil
}

func (vm *VirtualMachine) operationLt(lop, rop, r mem.Address) error {
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

	result := f1 < f2

	if err := vm.mm.SetValue(result, r); err != nil {
		return err
	}

	return nil
}

func (vm *VirtualMachine) operationEqual(lop, rop, r mem.Address) error {
	lopv, err := vm.mm.GetValue(lop)
	if err != nil {
		return err
	}

	ropv, err := vm.mm.GetValue(rop)
	if err != nil {
		return err
	}

	if f1, f2, err := getNums(lopv, ropv); err == nil {
		result := f1 == f2
		if err := vm.mm.SetValue(result, r); err != nil {
			return err
		}
		return nil
	} else if c1, c2, err := getChars(lopv, ropv); err == nil {
		result := c1 == c2
		if err := vm.mm.SetValue(result, r); err != nil {
			return err
		}
		return nil
	} else if b1, b2, err := getBools(lopv, ropv); err == nil {
		result := b1 == b2
		if err := vm.mm.SetValue(result, r); err != nil {
			return err
		}
		return nil
	}

	return errutil.NewNoPosf("Cannot perform equal operation on given types")
}

func (vm *VirtualMachine) operationHead(lop, rop, r mem.Address) error {
	lopv, err := vm.mm.GetValue(lop)
	if err != nil {
		return err
	}

	l, err := getListManager(lopv)
	if err != nil {
		return err
	}

	if l.IsNum() {
		val, err := l.GetHeadNum()
		if err != nil {
			return err
		}
		if err := vm.mm.SetValue(val, r); err != nil {
			return err
		}
	} else if l.IsChar() {
		val, err := l.GetHeadChar()
		if err != nil {
			return err
		}
		if err := vm.mm.SetValue(val, r); err != nil {
			return err
		}
	} else if l.IsBool() {
		val, err := l.GetHeadBool()
		if err != nil {
			return err
		}
		if err := vm.mm.SetValue(val, r); err != nil {
			return err
		}
	} else if l.IsFunc() {
		val, err := l.GetHeadFunc()
		if err != nil {
			return err
		}
		if err := vm.mm.SetValue(val, r); err != nil {
			return err
		}
	} else if l.IsList() {
		val, err := l.GetHeadList()
		if err != nil {
			return err
		}
		if err := vm.mm.SetValue(val, r); err != nil {
			return err
		}
	} else {
		return errutil.NewNoPosf("Invalid list type")
	}

	return nil
}

func (vm *VirtualMachine) operationTail(lop, rop, r mem.Address) error {
	lopv, err := vm.mm.GetValue(lop)
	if err != nil {
		return err
	}

	l, err := getListManager(lopv)
	if err != nil {
		return err
	}

	if l.IsNum() {
		val, err := l.GetTailNum()
		if err != nil {
			return err
		}
		if err := vm.mm.SetValue(val, r); err != nil {
			return err
		}
	} else if l.IsChar() {
		val, err := l.GetTailChar()
		if err != nil {
			return err
		}
		if err := vm.mm.SetValue(val, r); err != nil {
			return err
		}
	} else if l.IsBool() {
		val, err := l.GetTailBool()
		if err != nil {
			return err
		}
		if err := vm.mm.SetValue(val, r); err != nil {
			return err
		}
	} else if l.IsFunc() {
		val, err := l.GetTailFunc()
		if err != nil {
			return err
		}
		if err := vm.mm.SetValue(val, r); err != nil {
			return err
		}
	} else if l.IsList() {
		val, err := l.GetTailList()
		if err != nil {
			return err
		}
		if err := vm.mm.SetValue(val, r); err != nil {
			return err
		}
	} else {
		return errutil.NewNoPosf("Invalid list type")
	}

	return nil
}

func (vm *VirtualMachine) operationIns(lop, rop, r mem.Address) error {
	lopv, err := vm.mm.GetValue(lop)
	if err != nil {
		return err
	}

	ropv, err := vm.mm.GetValue(rop)
	if err != nil {
		return err
	}

	l, err := getListManager(ropv)
	if err != nil {
		return err
	}

	newlist, err := l.Insert(lopv)
	if err != nil {
		return err
	}

	if err := vm.mm.SetValue(newlist, r); err != nil {
		return err
	}

	return nil
}

func (vm *VirtualMachine) operationApp(lop, rop, r mem.Address) error {
	lopv, err := vm.mm.GetValue(lop)
	if err != nil {
		return err
	}

	ropv, err := vm.mm.GetValue(rop)
	if err != nil {
		return err
	}

	l1, err := getListManager(lopv)
	if err != nil {
		return err
	}

	l2, err := getListManager(ropv)
	if err != nil {
		return err
	}

	newlist, err := l1.Append(l2)
	if err != nil {
		return err
	}

	if err := vm.mm.SetValue(newlist, r); err != nil {
		return err
	}

	return nil
}

func (vm *VirtualMachine) operationEmp(lop, rop, r mem.Address) error {
	lopv, err := vm.mm.GetValue(lop)
	if err != nil {
		return err
	}

	l, err := getListManager(lopv)
	if err != nil {
		return err
	}

	result, err := l.Empty()
	if err != nil {
		return err
	}

	if err := vm.mm.SetValue(result, r); err != nil {
		return err
	}

	return nil
}

func (vm *VirtualMachine) operationLst(lop, rop, r mem.Address) error {
	// Get the type of the list
	typ := int(lop)

	newlist := list.NewListManager(typ)

	vm.pendinglists.Push(newlist)

	return nil
}

func (vm *VirtualMachine) operationPalst(lop, rop, r mem.Address) error {
	currentlist := vm.pendinglists.Top()

	lopv, err := vm.mm.GetValue(lop)
	if err != nil {
		return err
	}

	if err := currentlist.Add(lopv); err != nil {
		return err
	}

	return nil
}

func (vm *VirtualMachine) operationGelst(lop, rop, r mem.Address) error {
	currentlist := vm.pendinglists.Top()
	vm.pendinglists.Pop()

	if err := vm.mm.SetValue(currentlist, r); err != nil {
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

func (vm *VirtualMachine) operationEra(lop, rop, r mem.Address) error {
	vm.pendingcalls.Push(ar.NewActivationRecord())
	return nil
}

func (vm *VirtualMachine) operationParam(lop, rop, r mem.Address) error {
	nextcall := vm.pendingcalls.Top()

	// We check if the address is in the global scope, if it is, it means that
	// the address is a function address so we add it directly
	if lop < mem.Localstart {
		nextcall.AddFuncParam(int(lop))
		return nil
	}

	typedaddr := getTypeAddr(lop)

	lopv, err := vm.mm.GetValue(lop)
	if err != nil {
		return err
	}

	switch {
	case typedaddr < mem.CharOffset:
		f1, err := getNum(lopv)
		if err != nil {
			return err
		}
		nextcall.AddNumParam(f1)
	case typedaddr < mem.BoolOffset:
		f1, err := getChar(lopv)
		if err != nil {
			return err
		}
		nextcall.AddCharParam(f1)
	case typedaddr < mem.FunctionOffset:
		f1, err := getBool(lopv)
		if err != nil {
			return err
		}
		nextcall.AddBoolParam(f1)
	case typedaddr < mem.ListOffset:
		f1, err := getInt(lopv)
		if err != nil {
			return err
		}
		nextcall.AddFuncParam(f1)
	default:
		f1, err := getListManager(lopv)
		if err != nil {
			return err
		}
		nextcall.AddListParam(f1)
	}

	return nil
}

func (vm *VirtualMachine) operationCall(lop, rop, r mem.Address) error {
	// First we get the current call
	currcall := vm.ar.Top()
	//fmt.Printf("CURR CALL %+v\n", currcall)
	// Now we get the call which is about to happen
	nextcall := vm.pendingcalls.Top()

	// And we remove it from the pending calls
	// This allows us to have calls within calls
	vm.pendingcalls.Pop()

	// Now we copy the temp values of the current call to its activation record
	// so that it can be restored later
	if err := vm.copyTempToAR(currcall); err != nil {
		return err
	}

	// Now we initialize the local memory with the values of the parameters of the
	// function
	if err := vm.copyParamsToLocal(nextcall); err != nil {
		return err
	}

	// Now we set the return IP of the incoming call to the current ip
	nextcall.SetRetIp(vm.ip)

	// We add the incoming call to the activation stack
	vm.ar.Push(nextcall)

	// Now we get the location for the function
	var jump int

	j, err := vm.mm.GetValue(lop)
	if err != nil {
		return err
	}

	funcloc, err := getInt(j)
	if err != nil {
		return err
	}

	jump = funcloc

	if jump < 0 {
		return errutil.NewNoPosf("Invalid instruction address")
	}

	// And finally we set the current ip to the new location
	vm.ip = jump

	return nil
}

func (vm *VirtualMachine) operationRet(lop, rop, r mem.Address) error {
	// First we get the return value from the quadruple
	retv, err := vm.mm.GetValue(lop)
	if err != nil {
		return err
	}

	// We get the return instruction pointer
	retip := vm.ar.Top().Retip()

	// We kill the current activation record
	vm.ar.Pop()

	newcurrcall := vm.ar.Top()

	// We reactivate the local variables to the main memory
	if err := vm.copyParamsToLocal(newcurrcall); err != nil {
		return err
	}

	// We copy back the temp values to the main memory
	if err := vm.copyTempToMemory(newcurrcall); err != nil {
		return err
	}

	// We get the address where we need to save the return value
	callSaveAddr := vm.quads[retip].R()

	if f, err := getNum(retv); err == nil {
		if err := vm.mm.SetValue(f, callSaveAddr); err != nil {
			return err
		}
	} else if c, err := getChar(retv); err == nil {
		if err := vm.mm.SetValue(c, callSaveAddr); err != nil {
			return err
		}
	} else if b, err := getBool(retv); err == nil {
		if err := vm.mm.SetValue(b, callSaveAddr); err != nil {
			return err
		}
	} else if i, err := getInt(retv); err == nil {
		if err := vm.mm.SetValue(i, callSaveAddr); err != nil {
			return err
		}
	} else if i, err := getListManager(retv); err == nil {
		if err := vm.mm.SetValue(i, callSaveAddr); err != nil {
			return err
		}
	} else {
		return errutil.NewNoPosf("Cannot get valid form for value")
	}

	// We update the ip with the saved value
	vm.ip = retip + 1

	return nil
}

func (vm *VirtualMachine) operationAssign(lop, rop, r mem.Address) error {
	lopv, err := vm.mm.GetValue(lop)
	if err != nil {
		return err
	}

	if err := vm.mm.SetValue(lopv, r); err != nil {
		return err
	}

	return nil
}

func (vm *VirtualMachine) operationGoto(lop, rop, r mem.Address) error {
	if r < 0 {
		return errutil.NewNoPosf("Invalid instruction address")
	}

	vm.ip = int(r)

	return nil
}

func (vm *VirtualMachine) operationGotoT(lop, rop, r mem.Address) error {
	lopv, err := vm.mm.GetValue(lop)
	if err != nil {
		return err
	}

	b1, err := getBool(lopv)
	if err != nil {
		return err
	}

	jump := int(r)

	if b1 {
		vm.ip = jump
	} else {
		vm.ip++
	}

	return nil
}

func (vm *VirtualMachine) operationGotoF(lop, rop, r mem.Address) error {
	lopv, err := vm.mm.GetValue(lop)
	if err != nil {
		return err
	}

	b1, err := getBool(lopv)
	if err != nil {
		return err
	}

	jump := int(r)

	if !b1 {
		vm.ip = jump
	} else {
		vm.ip++
	}

	return nil
}
