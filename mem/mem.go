package mem

type Address int

const globalstart = 0
const localstart = 5000
const tempstart = 10000
const constantstart = 15000
const scopestart = 20000

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
	globalcount   int
	localcount    int
	tempcount     int
	constantcount int
	scopecount    int
	constantmap   map[string]int
}

// NewVirtualMemory creates a new virtual memory
func NewVirtualMemory() *VirtualMemory {
	return &VirtualMemory{globalstart, localstart, tempstart, constantstart, scopestart, make(map[string]int)}
}

func (vm *VirtualMemory) GetNextLocal() Address {
	result := vm.localcount
	vm.localcount += 1
	return Address(result)
}

func (vm *VirtualMemory) GetNextTemp() Address {
	result := vm.tempcount
	vm.tempcount += 1
	return Address(result)
}

func (vm *VirtualMemory) GetNextConstant() Address {
	result := vm.constantcount
	vm.constantcount += 1
	return Address(result)
}

func (vm *VirtualMemory) GetNextScope() Address {
	result := vm.scopecount
	vm.scopecount += 1
	return Address(result)
}

func (vm *VirtualMemory) ResetLocal() {
	vm.localcount = localstart
}

func (vm *VirtualMemory) ConstantExists(c string) bool {
	_, ok := vm.constantmap[c]
	return ok
}

func (vm *VirtualMemory) AddConstant(c string) Address {
	if vm.ConstantExists(c string) {
		addr := Address(vm.constantmap[c])
		return addr
	}

	// TODO: Determine the type of the constant and address accordingly
	nextAddr := vm.GetNextConstant()
	vm.constantmap[c] = nextAddr

	return Address(nextAddr)
}

func (vm *VirtualMemory) GetConstantAddress(c string) Address {
	a, ok := vm.constantmap[c]
	if !ok {
		return Address(-1)
	}

	return Address(a)
}
