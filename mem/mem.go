package mem

import (
	"fmt"

	"github.com/Loptt/lambdish-compiler/types"
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

func (a Address) String() string {
	if a < 0 {
		return "_"
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
		globalstart,
		globalstart,
		globalstart,
		globalstart,
		globalstart,
		localstart,
		localstart,
		localstart,
		localstart,
		localstart,
		tempstart,
		tempstart,
		tempstart,
		tempstart,
		tempstart,
		constantstart,
		constantstart,
		constantstart,
		constantstart,
		constantstart,
		scopestart,
		scopestart,
		scopestart,
		scopestart,
		scopestart,
		make(map[string]int),
	}
}

func (vm *VirtualMemory) GetNextLocal(t *types.LambdishType) Address {
	switch t.String() {
	// Num
	case "1":
		result := vm.localnumcount + numOffset
		vm.localnumcount++
		return Address(result)
	// Char
	case "2":
		result := vm.localnumcount + charOffset
		vm.localcharcount++
		return Address(result)
	// Bool
	case "3":
		result := vm.localnumcount + boolOffset
		vm.localboolcount++
		return Address(result)
	}

	if t.List() > 0 {
		result := vm.locallistcount + listOffset
		vm.locallistcount++
		return Address(result)
	}

	if t.Function() {
		result := vm.localfunctioncount + functionOffset
		vm.localfunctioncount++
		return Address(result)
	}

	return Address(-1)
}

func (vm *VirtualMemory) GetNextTemp(t *types.LambdishType) Address {
	switch t.String() {
	// Num
	case "1":
		result := vm.tempnumcount + numOffset
		vm.tempnumcount++
		return Address(result)
	// Char
	case "2":
		result := vm.tempnumcount + charOffset
		vm.tempcharcount++
		return Address(result)
	// Bool
	case "3":
		result := vm.tempnumcount + boolOffset
		vm.tempboolcount++
		return Address(result)
	}

	if t.List() > 0 {
		result := vm.templistcount + listOffset
		vm.templistcount++
		return Address(result)
	}

	if t.Function() {
		result := vm.tempfunctioncount + functionOffset
		vm.tempfunctioncount++
		return Address(result)
	}

	return Address(-1)
}

func (vm *VirtualMemory) getNextConstant(t *types.LambdishType) Address {
	switch t.String() {
	// Num
	case "1":
		result := vm.constantnumcount + numOffset
		vm.constantnumcount++
		return Address(result)
	// Char
	case "2":
		result := vm.constantnumcount + charOffset
		vm.constantcharcount++
		return Address(result)
	// Bool
	case "3":
		result := vm.constantnumcount + boolOffset
		vm.constantboolcount++
		return Address(result)
	}

	if t.List() > 0 {
		result := vm.constantlistcount + listOffset
		vm.constantlistcount++
		return Address(result)
	}

	if t.Function() {
		result := vm.constantfunctioncount + functionOffset
		vm.constantfunctioncount++
		return Address(result)
	}

	return Address(-1)
}

func (vm *VirtualMemory) GetNextScope(t *types.LambdishType) Address {
	switch t.String() {
	// Num
	case "1":
		result := vm.scopenumcount + numOffset
		vm.scopenumcount++
		return Address(result)
	// Char
	case "2":
		result := vm.scopenumcount + charOffset
		vm.scopecharcount++
		return Address(result)
	// Bool
	case "3":
		result := vm.scopenumcount + boolOffset
		vm.scopeboolcount++
		return Address(result)
	}

	if t.List() > 0 {
		result := vm.scopelistcount + listOffset
		vm.scopelistcount++
		return Address(result)
	}

	if t.Function() {
		result := vm.scopefunctioncount + functionOffset
		vm.scopefunctioncount++
		return Address(result)
	}

	return Address(-1)
}

func (vm *VirtualMemory) ResetLocal() {
	vm.localnumcount = localstart
	vm.localboolcount = localstart
	vm.localcharcount = localstart
	vm.locallistcount = localstart
}

func (vm *VirtualMemory) ResetTemp() {
	vm.tempnumcount = tempstart
	vm.tempboolcount = tempstart
	vm.tempcharcount = tempstart
	vm.templistcount = tempstart
}

func (vm *VirtualMemory) ConstantExists(c string) bool {
	_, ok := vm.constantmap[c]
	return ok
}

func (vm *VirtualMemory) AddConstant(c string, t *types.LambdishType) Address {
	if vm.ConstantExists(c) {
		addr := Address(vm.constantmap[c])
		return addr
	}

	// TODO: Determine the type of the constant and address accordingly
	nextAddr := vm.getNextConstant(t)
	vm.constantmap[c] = int(nextAddr)

	return Address(nextAddr)
}

func (vm *VirtualMemory) GetConstantAddress(c string) Address {
	a, ok := vm.constantmap[c]
	if !ok {
		return Address(-1)
	}

	return Address(a)
}
