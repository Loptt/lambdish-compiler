package vm

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Loptt/lambdish-compiler/mem"
	"github.com/Loptt/lambdish-compiler/quad"
	"github.com/Loptt/lambdish-compiler/vm/ar"
	"github.com/Loptt/lambdish-compiler/vm/list"
	"github.com/mewkiz/pkg/errutil"
)

// VirtualMachine contains the necesary attributes of a virtual machine to execute sequential code,
// manage memory, and call submodules
type VirtualMachine struct {
	ip           int
	quads        []*quad.Quadruple
	mm           *Memory
	ar           *ar.ArStack
	pendingcalls *ar.ArStack
	pendinglists *list.ListStack
	output       interface{}
}

// String represents the vm in a strctured format so that it can be easily debugged
func (vm *VirtualMachine) String() string {
	var builder strings.Builder

	builder.WriteString("VirtualMachine:\n")
	builder.WriteString(fmt.Sprintf("  IP: %d\n", vm.ip))
	builder.WriteString("  Quads:\n")

	for i, q := range vm.quads {
		builder.WriteString(fmt.Sprintf("    %d: %s\n", i, q))
	}

	builder.WriteString("\n")

	builder.WriteString(vm.mm.String())

	return builder.String()
}

// loadInstructions takes an array on strings and converts them to the corresponding quadruples
func (vm *VirtualMachine) loadInstructions(lines []string) error {
	var (
		op  quad.Operation
		rop mem.Address
		lop mem.Address
		r   mem.Address
	)

	// We iterate on the next iamount lines to get and parse each instruction
	for _, l := range lines {
		fields := strings.Fields(l)
		op = quad.StringToOperation(fields[0])

		addr, err := strconv.Atoi(fields[1])
		if err != nil {
			return err
		}
		lop = mem.Address(mem.Address(addr))

		addr, err = strconv.Atoi(fields[2])
		if err != nil {
			return err
		}
		rop = mem.Address(mem.Address(addr))

		addr, err = strconv.Atoi(fields[3])
		if err != nil {
			return err
		}
		r = mem.Address(mem.Address(addr))
		vm.quads = append(vm.quads, quad.NewQuadruple(op, lop, rop, r))
	}

	return nil
}

// loadConstants takes an array of strings and converts them to the corresponding constants in memory
func (vm *VirtualMachine) loadConstants(lines []string) error {
	for _, l := range lines {
		fields := strings.Fields(l)

		addr, err := strconv.Atoi(fields[1])
		if err != nil {
			return err
		}

		switch {
		case addr < mem.Constantstart+mem.CharOffset: // Number
			num, err := strconv.ParseFloat(fields[0], 64)
			if err != nil {
				return err
			}

			if err := vm.mm.SetValue(num, mem.Address(addr)); err != nil {
				return err
			}
		case addr < mem.Constantstart+mem.BoolOffset: // Char
			char := rune(fields[0][1])

			if err := vm.mm.SetValue(char, mem.Address(addr)); err != nil {
				return err
			}
		case addr < mem.Constantstart+mem.FunctionOffset: // Bool
			boolean, err := strconv.ParseBool(fields[0])
			if err != nil {
				return err
			}

			if err := vm.mm.SetValue(boolean, mem.Address(addr)); err != nil {
				return err
			}
		default:
			return errutil.NewNoPosf("Cannot set non-constant value")
		}
	}

	return nil
}

// LoadProgram takes a path to a file and parses its contents to quadruples and constants
func (vm *VirtualMachine) LoadProgram(path string) error {
	input, err := readFile(path)
	if err != nil {
		return err
	}

	lines := strings.Split(string(input), "\n")

	// Get the Amount of instructions from the top of the file
	iamount, err := strconv.Atoi(lines[0])
	if err != nil {
		return err
	}

	err = vm.loadInstructions(lines[1:(iamount + 1)])
	if err != nil {
		return err
	}

	// We get the constant amount at the end of the instructions
	camount, err := strconv.Atoi(lines[iamount+1])
	if err != nil {
		return err
	}

	cstart := iamount + 2

	err = vm.loadConstants(lines[cstart:(cstart + camount)])
	if err != nil {
		return err
	}

	return nil
}

func (vm *VirtualMachine) printOutput() {
	if c, ok := vm.output.(rune); ok {
		fmt.Printf("%c\n", c)
	} else if f, ok := vm.output.(float64); ok {
		if f == float64(int64(f)) {
			fmt.Printf("%d\n", int64(f))
		} else {
			fmt.Printf("%f\n", f)
		}
	} else {
		fmt.Printf("%v\n", vm.output)
	}
}

// executeNextInstruction indexes the next quadruple with the instruction pointer and
// proceeds to execute that instruction
func (vm *VirtualMachine) executeNextInstruction() error {
	q := vm.quads[vm.ip]

	//fmt.Printf("%d: Operation: %s\n", vm.ip, q.Op().String())

	switch q.Op() {
	case quad.Add:
		if err := vm.operationAdd(q.Lop(), q.Rop(), q.R()); err != nil {
			return err
		}
		vm.ip++
	case quad.Sub:
		if err := vm.operationSub(q.Lop(), q.Rop(), q.R()); err != nil {
			return err
		}
		vm.ip++
	case quad.Mult:
		if err := vm.operationMult(q.Lop(), q.Rop(), q.R()); err != nil {
			return err
		}
		vm.ip++
	case quad.Div:
		if err := vm.operationDiv(q.Lop(), q.Rop(), q.R()); err != nil {
			return err
		}
		vm.ip++
	case quad.Mod:
		if err := vm.operationMod(q.Lop(), q.Rop(), q.R()); err != nil {
			return err
		}
		vm.ip++
	case quad.And:
		if err := vm.operationAnd(q.Lop(), q.Rop(), q.R()); err != nil {
			return err
		}
		vm.ip++
	case quad.Or:
		if err := vm.operationOr(q.Lop(), q.Rop(), q.R()); err != nil {
			return err
		}
		vm.ip++
	case quad.Not:
		if err := vm.operationNot(q.Lop(), q.Rop(), q.R()); err != nil {
			return err
		}
		vm.ip++
	case quad.Gt:
		if err := vm.operationGt(q.Lop(), q.Rop(), q.R()); err != nil {
			return err
		}
		vm.ip++
	case quad.Lt:
		if err := vm.operationLt(q.Lop(), q.Rop(), q.R()); err != nil {
			return err
		}
		vm.ip++
	case quad.Equal:
		if err := vm.operationEqual(q.Lop(), q.Rop(), q.R()); err != nil {
			return err
		}
		vm.ip++
	case quad.Head:
		if err := vm.operationHead(q.Lop(), q.Rop(), q.R()); err != nil {
			return err
		}
		vm.ip++
	case quad.Tail:
		if err := vm.operationTail(q.Lop(), q.Rop(), q.R()); err != nil {
			return err
		}
		vm.ip++
	case quad.Ins:
		if err := vm.operationIns(q.Lop(), q.Rop(), q.R()); err != nil {
			return err
		}
		vm.ip++
	case quad.App:
		if err := vm.operationApp(q.Lop(), q.Rop(), q.R()); err != nil {
			return err
		}
		vm.ip++
	case quad.Emp:
		if err := vm.operationEmp(q.Lop(), q.Rop(), q.R()); err != nil {
			return err
		}
		vm.ip++
	case quad.Lst:
		if err := vm.operationLst(q.Lop(), q.Rop(), q.R()); err != nil {
			return err
		}
		vm.ip++
	case quad.PaLst:
		if err := vm.operationPalst(q.Lop(), q.Rop(), q.R()); err != nil {
			return err
		}
		vm.ip++
	case quad.GeLst:
		if err := vm.operationGelst(q.Lop(), q.Rop(), q.R()); err != nil {
			return err
		}
		vm.ip++
	case quad.Print:
		if err := vm.operationPrint(q.Lop(), q.Rop(), q.R()); err != nil {
			return err
		}
		vm.ip++
	case quad.Era:
		if err := vm.operationEra(q.Lop(), q.Rop(), q.R()); err != nil {
			return err
		}
		vm.ip++
	case quad.Param:
		if err := vm.operationParam(q.Lop(), q.Rop(), q.R()); err != nil {
			return err
		}
		vm.ip++
	case quad.Call:
		if err := vm.operationCall(q.Lop(), q.Rop(), q.R()); err != nil {
			return err
		}
	case quad.Ret:
		if err := vm.operationRet(q.Lop(), q.Rop(), q.R()); err != nil {
			return err
		}
	case quad.Assign:
		if err := vm.operationAssign(q.Lop(), q.Rop(), q.R()); err != nil {
			return err
		}
		vm.ip++
	case quad.Goto:
		if err := vm.operationGoto(q.Lop(), q.Rop(), q.R()); err != nil {
			return err
		}
	case quad.GotoT:
		if err := vm.operationGotoT(q.Lop(), q.Rop(), q.R()); err != nil {
			return err
		}
	case quad.GotoF:
		if err := vm.operationGotoF(q.Lop(), q.Rop(), q.R()); err != nil {
			return err
		}
	}

	return nil
}

// Run starts executing the instructions in the virtual machine
func (vm *VirtualMachine) Run() error {
	if len(vm.quads) < 1 {
		return errutil.NewNoPosf("No instructions to execute")
	}

	// Push the main activation record
	vm.ar.Push(ar.NewActivationRecord())

	for vm.ip < len(vm.quads) {
		if err := vm.executeNextInstruction(); err != nil {
			return err
		}
	}

	vm.printOutput()

	return nil
}

//NewVirtualMachine ...
func NewVirtualMachine() *VirtualMachine {
	return &VirtualMachine{0, make([]*quad.Quadruple, 0), NewMemory(), ar.NewArStack(), ar.NewArStack(), list.NewListStack(), 0}
}
