package vm

import (
	"fmt"
	"strings"

	"github.com/Loptt/lambdish-compiler/mem"
	"github.com/mewkiz/pkg/errutil"
)

//MemorySegment ...
type MemorySegment struct {
	num      []float64
	char     []rune
	booleans []bool
	function []mem.Address
	list     []mem.Address
	base     mem.Address
	name     string
}

func (ms *MemorySegment) SetValue(v interface{}, addr mem.Address) error {
	baseaddr := addr - ms.base

	switch {
	case baseaddr < mem.NumOffset: // Error
		return errutil.Newf("Address out of scope")
	case baseaddr < mem.CharOffset: // Number
		if n, ok := v.(float64); ok {
			typebaseaddr := int(baseaddr - mem.NumOffset)
			// If the specified address is bigger than the current size of the array
			// we need to grow the array to that size
			if len(ms.num) <= typebaseaddr {
				// First we create a new slice with the extra cells we need
				newslice := make([]float64, typebaseaddr-len(ms.num)+1)
				ms.num = append(ms.num, newslice...)
				// Now we set the value to the specified address
				ms.num[typebaseaddr] = n
			} else {
				ms.num[typebaseaddr] = n
			}
			return nil
		}
		return errutil.Newf("Cannot set non-number in number address range")
	case baseaddr < mem.BoolOffset: // Character
		if c, ok := v.(rune); ok {
			typebaseaddr := int(baseaddr - mem.CharOffset)
			// If the specified address is bigger than the current size of the array
			// we need to grow the array to that size
			if len(ms.char) <= typebaseaddr {
				// First we create a new slice with the extra cells we need
				newslice := make([]rune, typebaseaddr-len(ms.char)+1)
				ms.char = append(ms.char, newslice...)
				// Now we set the value to the specified address
				ms.char[typebaseaddr] = c
			} else {
				ms.char[typebaseaddr] = c
			}
			return nil
		}
		return errutil.Newf("Cannot set non-char in char address range")
	case baseaddr < mem.FunctionOffset: // Boolean
		if b, ok := v.(bool); ok {
			typebaseaddr := int(baseaddr - mem.BoolOffset)
			// If the specified address is bigger than the current size of the array
			// we need to grow the array to that size
			if len(ms.booleans) <= typebaseaddr {
				// First we create a new slice with the extra cells we need
				newslice := make([]bool, typebaseaddr-len(ms.booleans)+1)
				ms.booleans = append(ms.booleans, newslice...)
				// Now we set the value to the specified address
				ms.booleans[typebaseaddr] = b
			} else {
				ms.booleans[typebaseaddr] = b
			}
			return nil
		}
		return errutil.Newf("Cannot set non-boolean in boolean address range")
	case baseaddr < mem.ListOffset: //Function
		if a, ok := v.(mem.Address); ok {
			typebaseaddr := int(baseaddr - mem.FunctionOffset)
			// If the specified address is bigger than the current size of the array
			// we need to grow the array to that size
			if len(ms.function) <= typebaseaddr {
				// First we create a new slice with the extra cells we need
				newslice := make([]mem.Address, typebaseaddr-len(ms.function)+1)
				ms.function = append(ms.function, newslice...)
				// Now we set the value to the specified address
				ms.function[typebaseaddr] = a
			} else {
				ms.function[typebaseaddr] = a
			}
			return nil
		}
		return errutil.Newf("Cannot set non-function in function address range")
	case baseaddr < mem.ListOffset+1000: //List
		if a, ok := v.(mem.Address); ok {
			typebaseaddr := int(baseaddr - mem.ListOffset)
			// If the specified address is bigger than the current size of the array
			// we need to grow the array to that size
			if len(ms.list) <= typebaseaddr {
				// First we create a new slice with the extra cells we need
				newslice := make([]mem.Address, typebaseaddr-len(ms.list)+1)
				ms.list = append(ms.list, newslice...)
				// Now we set the value to the specified address
				ms.list[typebaseaddr] = a
			} else {
				ms.list[typebaseaddr] = a
			}
			return nil
		}
		return errutil.Newf("Cannot set non-list in list address range")
	default: // Error
		return errutil.Newf("Address out of scope")
	}

}

func (ms *MemorySegment) GetValue(addr mem.Address) (interface{}, error) {
	baseaddr := addr - ms.base

	switch {
	case baseaddr < mem.NumOffset: // Error
		return nil, errutil.Newf("Address out of scope")
	case baseaddr < mem.CharOffset: // Number
		typebaseaddr := int(baseaddr - mem.NumOffset)
		if len(ms.num) <= typebaseaddr {
			return nil, errutil.Newf("Referencing address out of scope")
		}
		return ms.num[typebaseaddr], nil
	case baseaddr < mem.BoolOffset: // Character
		typebaseaddr := int(baseaddr - mem.CharOffset)
		if len(ms.char) <= typebaseaddr {
			return nil, errutil.Newf("Referencing address out of scope")
		}
		return ms.char[typebaseaddr], nil
	case baseaddr < mem.FunctionOffset: // Boolean
		typebaseaddr := int(baseaddr - mem.BoolOffset)
		if len(ms.booleans) <= typebaseaddr {
			return nil, errutil.Newf("Referencing address out of scope")
		}
		return ms.booleans[typebaseaddr], nil
	case baseaddr < mem.ListOffset: //Function
		typebaseaddr := int(baseaddr - mem.FunctionOffset)
		if len(ms.function) <= typebaseaddr {
			return nil, errutil.Newf("Referencing address out of scope")
		}
		return ms.function[typebaseaddr], nil
	case baseaddr < mem.ListOffset+1000: //List
		typebaseaddr := int(baseaddr - mem.ListOffset)
		if len(ms.list) <= typebaseaddr {
			return nil, errutil.Newf("Referencing address out of scope")
		}
		return ms.list[typebaseaddr], nil
	default: // Error
		return nil, errutil.Newf("Address out of scope")
	}

}

func (ms *MemorySegment) String() string {
	var builder strings.Builder

	builder.WriteString(fmt.Sprintf("  %s:\n", ms.name))

	builder.WriteString("    Num:\n")
	for i, v := range ms.num {
		builder.WriteString(fmt.Sprintf("      %d: %f\n", i, v))
	}
	builder.WriteString("    Char:\n")
	for i, v := range ms.char {
		builder.WriteString(fmt.Sprintf("      %d: %c\n", i, v))
	}
	builder.WriteString("    Bool:\n")
	for i, v := range ms.booleans {
		builder.WriteString(fmt.Sprintf("      %d: %t\n", i, v))
	}
	builder.WriteString("    Function:\n")
	for i, v := range ms.function {
		builder.WriteString(fmt.Sprintf("      %d: %d\n", i, v))
	}
	builder.WriteString("    List:\n")
	for i, v := range ms.list {
		builder.WriteString(fmt.Sprintf("      %d: %d\n", i, v))
	}

	return builder.String()
}

func NewMemorySegment(base int, name string) *MemorySegment {
	return &MemorySegment{
		make([]float64, 0),
		make([]rune, 0), make([]bool, 0),
		make([]mem.Address, 0),
		make([]mem.Address, 0),
		mem.Address(base),
		name,
	}
}

type Memory struct {
	memglobal   *MemorySegment
	memlocal    *MemorySegment
	memtemp     *MemorySegment
	memconstant *MemorySegment
	memscope    *MemorySegment
}

func NewMemory() *Memory {
	return &Memory{
		NewMemorySegment(mem.Globalstart, "Global"),
		NewMemorySegment(mem.Localstart, "Local"),
		NewMemorySegment(mem.Tempstart, "Temp"),
		NewMemorySegment(mem.Constantstart, "Constant"),
		NewMemorySegment(mem.Scopestart, "Scope"),
	}
}

func (m *Memory) GetValue(addr mem.Address) (interface{}, error) {
	switch {
	case addr < mem.Globalstart: // Error
		return false, errutil.Newf("Address out of scope")
	case addr < mem.Localstart: // Global
		v, err := m.memglobal.GetValue(addr)
		if err != nil {
			return nil, err
		}
		return v, nil
	case addr < mem.Tempstart: // Local
		v, err := m.memlocal.GetValue(addr)
		if err != nil {
			return nil, err
		}
		return v, nil
	case addr < mem.Constantstart: // Temp
		v, err := m.memtemp.GetValue(addr)
		if err != nil {
			return nil, err
		}
		return v, nil
	case addr < mem.Scopestart: // Constant
		v, err := m.memconstant.GetValue(addr)
		if err != nil {
			return nil, err
		}
		return v, nil
	case addr < mem.Scopestart+5000: // Scope
		v, err := m.memscope.GetValue(addr)
		if err != nil {
			return nil, err
		}
		return v, nil
	default: // Error
		return nil, errutil.Newf("Address out of scope")
	}
}

func (m *Memory) SetValue(v interface{}, addr mem.Address) error {
	switch {
	case addr < mem.Globalstart: // Error
		return errutil.Newf("Address out of scope")
	case addr < mem.Localstart: // Global
		if err := m.memglobal.SetValue(v, addr); err != nil {
			return err
		}
		return nil
	case addr < mem.Tempstart: // Local
		if err := m.memlocal.SetValue(v, addr); err != nil {
			return err
		}
		return nil
	case addr < mem.Constantstart: // Temp
		if err := m.memtemp.SetValue(v, addr); err != nil {
			return err
		}
		return nil
	case addr < mem.Scopestart: // Constant
		if err := m.memconstant.SetValue(v, addr); err != nil {
			return err
		}
		return nil
	case addr < mem.Scopestart+5000: // Scope
		if err := m.memscope.SetValue(v, addr); err != nil {
			return err
		}
		return nil
	default: // Error
		return errutil.Newf("Address out of scope")
	}
}

func (m *Memory) String() string {
	var builder strings.Builder

	builder.WriteString("Memory:\n")
	builder.WriteString(m.memglobal.String())
	builder.WriteString(m.memlocal.String())
	builder.WriteString(m.memtemp.String())
	builder.WriteString(m.memconstant.String())
	builder.WriteString(m.memscope.String())

	return builder.String()
}
