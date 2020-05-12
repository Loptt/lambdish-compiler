package mem

import (
	"fmt"

	"github.com/Loptt/lambdish-compiler/types"
	"github.com/mewkiz/pkg/errutil"
)

type Address int

const globalstart = 0
const localstart = 5000
const tempstart = 10000
const constantstart = 15000
const scopestart = 20000

const numOffset = 0
const charOffset = 1000
const boolOffset = 2000
const functionOffset = 3000
const listOffset = 4000

const segmentsize = 1000

func (a Address) String() string {
	if a < 0 {
		return "-1"
	}

	return fmt.Sprintf("%d", a)
}

/*
========Global = 0
0-999   Num
1000-1999 Char
2000-2999 Bool
3000-3999 Functions
4000-4999 Lists
========Local = 5000
5000-5999   Num
6000-6999 Char
7000-7999 Bool
8000-8999 Functions
9000-9999 Lists
======= Temp = 100000
10000-10999   Num
11000-11999 Char
12000-12999 Bool
1300-13999 Functions
14000-14999 Lists
========Constant = 15000
15000-15999  Num
16000-16999 Char
17000-17999 Bool
18000-18999 Functions
19000-19999 Lists
========OutScope = 20000
20000-20999  Num
21000-21999 Char
22000-22999 Bool
23000-23999 Functions
24000-24999 Lists
*/

type VirtualMemory struct {
	globalnumcount      int
	globalcharcount     int
	globalboolcount     int
	globalfunctioncount int
	globallistcount     int

	localnumcount      int
	localcharcount     int
	localboolcount     int
	localfunctioncount int
	locallistcount     int

	tempnumcount      int
	tempcharcount     int
	tempboolcount     int
	tempfunctioncount int
	templistcount     int

	constantnumcount      int
	constantcharcount     int
	constantboolcount     int
	constantfunctioncount int
	constantlistcount     int

	scopenumcount      int
	scopecharcount     int
	scopeboolcount     int
	scopefunctioncount int
	scopelistcount     int

	constantmap map[string]int
}

// NewVirtualMemory creates a new virtual memory
func NewVirtualMemory() *VirtualMemory {

	return &VirtualMemory{
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		make(map[string]int),
	}
}

func (vm *VirtualMemory) GetNextLocal(t *types.LambdishType) (Address, error) {
	switch t.String() {
	// Num
	case "1":
		if vm.localnumcount >= segmentsize {
			return Address(-1), errutil.Newf("Error: local variables for numbers exceeded.")
		}
		result := vm.localnumcount + numOffset + localstart
		vm.localnumcount++
		return Address(result), nil
	// Char
	case "2":
		if vm.localcharcount >= segmentsize {
			return Address(-1), errutil.Newf("Error: local variables for chars exceeded.")
		}
		result := vm.localcharcount + charOffset + localstart
		vm.localcharcount++
		return Address(result), nil
	// Bool
	case "3":
		if vm.localboolcount >= segmentsize {
			return Address(-1), errutil.Newf("Error: local variables for bools exceeded.")
		}
		result := vm.localboolcount + boolOffset + localstart
		vm.localboolcount++
		return Address(result), nil
	}

	if t.List() > 0 {
		if vm.locallistcount >= segmentsize {
			return Address(-1), errutil.Newf("Error: local variables for lists exceeded.")
		}
		result := vm.locallistcount + listOffset + localstart
		vm.locallistcount++
		return Address(result), nil
	}

	if t.Function() {
		if vm.localfunctioncount >= segmentsize {
			return Address(-1), errutil.Newf("Error: local variables for function exceeded.")
		}
		result := vm.localfunctioncount + functionOffset + localstart
		vm.localfunctioncount++
		return Address(result), nil
	}

	return Address(-1), errutil.Newf("Error: variable type not identified.")
}

func (vm *VirtualMemory) GetNextTemp(t *types.LambdishType) (Address, error) {
	switch t.String() {
	// Num
	case "1":
		if vm.tempnumcount >= segmentsize {
			return Address(-1), errutil.Newf("Error: temp variables for numbers exceeded.")
		}
		result := vm.tempnumcount + numOffset + tempstart
		vm.tempnumcount++
		return Address(result), nil
	// Char
	case "2":
		if vm.tempcharcount >= segmentsize {
			return Address(-1), errutil.Newf("Error: temp variables for chars exceeded.")
		}
		result := vm.tempcharcount + charOffset + tempstart
		vm.tempcharcount++
		return Address(result), nil
	// Bool
	case "3":
		if vm.tempboolcount >= segmentsize {
			return Address(-1), errutil.Newf("Error: temp variables for bools exceeded.")
		}
		result := vm.tempboolcount + boolOffset + tempstart
		vm.tempboolcount++
		return Address(result), nil
	}

	if t.List() > 0 {
		if vm.templistcount >= segmentsize {
			return Address(-1), errutil.Newf("Error: temp variables for lists exceeded.")
		}
		result := vm.templistcount + listOffset + tempstart
		vm.templistcount++
		return Address(result), nil
	}

	if t.Function() {
		if vm.tempfunctioncount >= segmentsize {
			return Address(-1), errutil.Newf("Error: temp variables for functions exceeded.")
		}
		result := vm.tempfunctioncount + functionOffset + tempstart
		vm.tempfunctioncount++
		return Address(result), nil
	}

	return Address(-1), errutil.Newf("Error: variable type not identified.")
}

func (vm *VirtualMemory) getNextConstant(t *types.LambdishType) (Address, error) {
	switch t.String() {
	// Num
	case "1":
		if vm.constantnumcount >= segmentsize {
			return Address(-1), errutil.Newf("Error: constant variables for numbers exceeded.")
		}
		result := vm.constantnumcount + numOffset + constantstart
		vm.constantnumcount++
		return Address(result), nil
	// Char
	case "2":
		if vm.constantcharcount >= segmentsize {
			return Address(-1), errutil.Newf("Error: constant variables for chars exceeded.")
		}
		result := vm.constantcharcount + charOffset + constantstart
		vm.constantcharcount++
		return Address(result), nil
	// Bool
	case "3":
		if vm.constantboolcount >= segmentsize {
			return Address(-1), errutil.Newf("Error: constant variables for bools exceeded.")
		}
		result := vm.constantboolcount + boolOffset + constantstart
		vm.constantboolcount++
		return Address(result), nil
	}

	if t.List() > 0 {
		if vm.constantlistcount >= segmentsize {
			return Address(-1), errutil.Newf("Error: constant variables for lists exceeded.")
		}
		result := vm.constantlistcount + listOffset + constantstart
		vm.constantlistcount++
		return Address(result), nil
	}

	if t.Function() {
		if vm.constantfunctioncount >= segmentsize {
			return Address(-1), errutil.Newf("Error: constant variables for functions exceeded.")
		}
		result := vm.constantfunctioncount + functionOffset + constantstart
		vm.constantfunctioncount++
		return Address(result), nil
	}

	return Address(-1), errutil.Newf("Error: variable type not identified.")
}

func (vm *VirtualMemory) GetNextScope(t *types.LambdishType) (Address, error) {
	switch t.String() {
	// Num
	case "1":
		if vm.scopenumcount >= segmentsize {
			return Address(-1), errutil.Newf("Error: scope variables for numbers exceeded.")
		}
		result := vm.scopenumcount + numOffset + scopestart
		vm.scopenumcount++
		return Address(result), nil
	// Char
	case "2":
		if vm.scopecharcount >= segmentsize {
			return Address(-1), errutil.Newf("Error: scope variables for chars exceeded.")
		}
		result := vm.scopecharcount + charOffset + scopestart
		vm.scopecharcount++
		return Address(result), nil
	// Bool
	case "3":
		if vm.scopeboolcount >= segmentsize {
			return Address(-1), errutil.Newf("Error: scope variables for bools exceeded.")
		}
		result := vm.scopeboolcount + boolOffset + scopestart
		vm.scopeboolcount++
		return Address(result), nil
	}

	if t.List() > 0 {
		if vm.scopelistcount >= segmentsize {
			return Address(-1), errutil.Newf("Error: scope variables for lists exceeded.")
		}
		result := vm.scopelistcount + listOffset + scopestart
		vm.scopelistcount++
		return Address(result), nil
	}

	if t.Function() {
		if vm.scopefunctioncount >= segmentsize {
			return Address(-1), errutil.Newf("Error: scope variables for numbers exceeded.")
		}
		result := vm.scopefunctioncount + functionOffset + scopestart
		vm.scopefunctioncount++
		return Address(result), nil
	}

	return Address(-1), errutil.Newf("Error: variable type not identified.")
}

func (vm *VirtualMemory) ResetLocal() {
	vm.localnumcount = 0
	vm.localboolcount = 0
	vm.localcharcount = 0
	vm.locallistcount = 0
}

func (vm *VirtualMemory) ResetTemp() {
	vm.tempnumcount = 0
	vm.tempboolcount = 0
	vm.tempcharcount = 0
	vm.templistcount = 0
}

func (vm *VirtualMemory) ConstantExists(c string) bool {
	_, ok := vm.constantmap[c]
	return ok
}

func (vm *VirtualMemory) AddConstant(c string, t *types.LambdishType) (Address, error) {
	if vm.ConstantExists(c) {
		addr := Address(vm.constantmap[c])
		return addr, nil
	}

	// TODO: Determine the type of the constant and address accordingly
	nextAddr, err := vm.getNextConstant(t)
	if err != nil {
		return Address(-1), err
	}
	vm.constantmap[c] = int(nextAddr)

	return Address(nextAddr), nil
}

func (vm *VirtualMemory) GetConstantAddress(c string) Address {
	a, ok := vm.constantmap[c]
	if !ok {
		return Address(-1)
	}

	return Address(a)
}
