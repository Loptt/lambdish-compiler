package vm

import (
	"github.com/Loptt/lambdish-compiler/mem"
)

//MemorySegment ...
type MemorySegment struct {
	num      []float64
	char     []rune
	booleans []bool
	function []mem.Address
	list     []mem.Address
}

func NewMemorySegment() *MemorySegment {
	return &MemorySegment{make([]float64, 0), make([]rune, 0), make([]bool, 0), make([]mem.Address, 0), make([]mem.Address, 0)}
}

type Memory struct {
	memglobal   *MemorySegment
	memlocal    *MemorySegment
	memtemp     *MemorySegment
	memconstant *MemorySegment
	memscope    *MemorySegment
}

func NewMemory() *Memory {
	return &Memory{NewMemorySegment(), NewMemorySegment(), NewMemorySegment(), NewMemorySegment(), NewMemorySegment()}
}

/*
func GetNum(addr mem.Address) float64 {
}

func GetChar(addr mem.Address) rune {

}

func GetBool(addr mem.Address) bool {

}

func GetFunction(addr mem.Address) mem.Address {

}

func GetList(addr mem.Address) mem.Address {

}
*/
