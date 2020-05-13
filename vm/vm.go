package vm

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Loptt/lambdish-compiler/ic"
	"github.com/Loptt/lambdish-compiler/mem"
)

type VirtualMachine struct {
	ip    int
	quads []*ic.Quadruple
}

func (vm *VirtualMachine) String() string {
	var builder strings.Builder
	builder.WriteString("VirtualMachine:\n")
	builder.WriteString(fmt.Sprintf("  IP: %d\n", vm.ip))
	builder.WriteString("  Quads:\n")

	for i, q := range vm.quads {
		builder.WriteString(fmt.Sprintf("    %d: %s\n", i, q))
	}

	return builder.String()
}

func (vm *VirtualMachine) loadProgram(path string) error {
	input, err := readFile(path)
	if err != nil {
		return err
	}

	fields := strings.Fields(string(input))

	for i, f := range fields {
		var op ic.Operation
		var rop mem.Address
		var lop mem.Address
		var r mem.Address

		switch i % 4 {
		case 0:
			op = ic.StringToOperation(f)
		case 1:
			addr, err := strconv.Atoi(f)
			if err != nil {
				return err
			}
			lop = mem.Address(mem.Address(addr))
		case 2:
			addr, err := strconv.Atoi(f)
			if err != nil {
				return err
			}
			rop = mem.Address(mem.Address(addr))
		case 3:
			addr, err := strconv.Atoi(f)
			if err != nil {
				return err
			}
			r = mem.Address(mem.Address(addr))
			vm.quads = append(vm.quads, ic.NewQuadruple(op, lop, rop, r))
		}
	}

	return nil
}

func NewVirtualMachine() *VirtualMachine {
	return &VirtualMachine{0, make([]*ic.Quadruple, 0)}
}
