package vm

import (
	"os"

	"github.com/Loptt/lambdish-compiler/mem"
	"github.com/Loptt/lambdish-compiler/vm/ar"
	"github.com/Loptt/lambdish-compiler/vm/list"
	"github.com/mewkiz/pkg/errutil"
)

func readFile(path string) ([]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	fileinfo, err := file.Stat()
	if err != nil {
		return nil, err
	}

	filesize := fileinfo.Size()
	buffer := make([]byte, filesize)

	_, err = file.Read(buffer)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

func getNum(v interface{}) (float64, error) {
	num, ok := v.(float64)
	if !ok {
		return 0, errutil.Newf("Cannot convert current value to num")
	}

	return num, nil
}

func getNums(v1, v2 interface{}) (float64, float64, error) {
	num1, ok := v1.(float64)
	if !ok {
		return 0, 0, errutil.Newf("Cannot convert current value to num")
	}

	num2, ok := v2.(float64)
	if !ok {
		return 0, 0, errutil.Newf("Cannot convert current value to num")
	}

	return num1, num2, nil
}

func getChar(v interface{}) (rune, error) {
	char, ok := v.(rune)
	if !ok {
		return 0, errutil.Newf("Cannot convert current value to char")
	}

	return char, nil
}

func getChars(v1, v2 interface{}) (rune, rune, error) {
	char1, ok := v1.(rune)
	if !ok {
		return 0, 0, errutil.Newf("Cannot convert current value to char")
	}

	char2, ok := v2.(rune)
	if !ok {
		return 0, 0, errutil.Newf("Cannot convert current value to char")
	}

	return char1, char2, nil
}

func getBool(v interface{}) (bool, error) {
	boo, ok := v.(bool)
	if !ok {
		return false, errutil.Newf("Cannot convert current value to bool")
	}

	return boo, nil
}

func getBools(v1, v2 interface{}) (bool, bool, error) {
	bool1, ok := v1.(bool)
	if !ok {
		return false, false, errutil.Newf("Cannot convert current value to bool")
	}

	bool2, ok := v2.(bool)
	if !ok {
		return false, false, errutil.Newf("Cannot convert current value to bool")
	}

	return bool1, bool2, nil
}

func getInt(v interface{}) (int, error) {
	in, ok := v.(int)
	if !ok {
		return 0, errutil.Newf("Cannot convert current value to address")
	}

	return in, nil
}

func getListManager(v interface{}) (*list.ListManager, error) {
	in, ok := v.(*list.ListManager)
	if !ok {
		return nil, errutil.Newf("Cannot convert current value to list manager")
	}

	return in, nil
}

func getTypeAddr(addr mem.Address) int {
	switch {
	case addr < mem.Localstart:
		return int(addr) - mem.Globalstart
	case addr < mem.Tempstart:
		return int(addr) - mem.Localstart
	case addr < mem.Constantstart:
		return int(addr) - mem.Tempstart
	case addr < mem.Scopestart:
		return int(addr) - mem.Constantstart
	default:
		return int(addr) - mem.Scopestart
	}
}

// copyTempoAR copies all the contents of the temp memory to the current activation record
// so that it can be restored later
func (vm *VirtualMachine) copyTempToAR(a *ar.ActivationRecord) error {
	mstemp := vm.mm.memtemp

	a.ResetTemps()

	for i, num := range mstemp.num {
		a.AddNumTemp(num, mem.Address(i+mem.NumOffset+mem.Tempstart))
	}

	for i, char := range mstemp.char {
		a.AddCharTemp(char, mem.Address(i+mem.CharOffset+mem.Tempstart))
	}

	for i, b := range mstemp.booleans {
		a.AddBoolTemp(b, mem.Address(i+mem.BoolOffset+mem.Tempstart))
	}

	for i, f := range mstemp.function {
		a.AddFuncTemp(int(f), mem.Address(i+mem.FunctionOffset+mem.Tempstart))
	}

	for i, l := range mstemp.list {
		a.AddListTemp(l, mem.Address(i+mem.ListOffset+mem.Tempstart))
	}

	return nil
}

func (vm *VirtualMachine) copyTempToMemory(a *ar.ActivationRecord) error {
	for _, p := range a.Numtemps() {
		if err := vm.mm.SetValue(p.Value(), p.Addr()); err != nil {
			return err
		}
	}

	for _, p := range a.Chartemps() {
		if err := vm.mm.SetValue(p.Value(), p.Addr()); err != nil {
			return err
		}
	}

	for _, p := range a.Booltemps() {
		if err := vm.mm.SetValue(p.Value(), p.Addr()); err != nil {
			return err
		}
	}

	for _, p := range a.Functemps() {
		if err := vm.mm.SetValue(p.Value(), p.Addr()); err != nil {
			return err
		}
	}

	for _, p := range a.Listtemps() {
		if err := vm.mm.SetValue(p.Value(), p.Addr()); err != nil {
			return err
		}
	}

	return nil
}

func (vm *VirtualMachine) copyParamsToLocal(a *ar.ActivationRecord) error {
	for _, p := range a.Numparams() {
		if err := vm.mm.SetValue(p.Value(), p.Addr()); err != nil {
			return err
		}
	}

	for _, p := range a.Charparams() {
		if err := vm.mm.SetValue(p.Value(), p.Addr()); err != nil {
			return err
		}
	}

	for _, p := range a.Boolparams() {
		if err := vm.mm.SetValue(p.Value(), p.Addr()); err != nil {
			return err
		}
	}

	for _, p := range a.Funcparams() {
		if err := vm.mm.SetValue(p.Value(), p.Addr()); err != nil {
			return err
		}
	}

	for _, p := range a.Listparams() {
		if err := vm.mm.SetValue(p.Value(), p.Addr()); err != nil {
			return err
		}
	}

	return nil
}
