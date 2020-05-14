package ar

import "github.com/Loptt/lambdish-compiler/mem"

type NumParam struct {
	value float64
	addr  mem.Address
}

func (np NumParam) Value() float64 {
	return np.value
}

func (np NumParam) Addr() mem.Address {
	return np.addr
}

type CharParam struct {
	value rune
	addr  mem.Address
}

func (np CharParam) Value() rune {
	return np.value
}

func (np CharParam) Addr() mem.Address {
	return np.addr
}

type BoolParam struct {
	value bool
	addr  mem.Address
}

func (np BoolParam) Value() bool {
	return np.value
}

func (np BoolParam) Addr() mem.Address {
	return np.addr
}

type FuncParam struct {
	value int
	addr  mem.Address
}

func (np FuncParam) Value() int {
	return np.value
}

func (np FuncParam) Addr() mem.Address {
	return np.addr
}

type ListParam struct {
	value int
	addr  mem.Address
}

func (np ListParam) Value() int {
	return np.value
}

func (np ListParam) Addr() mem.Address {
	return np.addr
}

type ActivationRecord struct {
	retip      int
	numparams  []NumParam
	charparams []CharParam
	boolparams []BoolParam
	funcparams []FuncParam
	listparams []ListParam
	numtemps   []NumParam
	chartemps  []CharParam
	booltemps  []BoolParam
	functemps  []FuncParam
	listtemps  []ListParam
	retnum     float64
	retchar    rune
	retbool    bool
	retfunc    int
	retlist    int
	numcount   int
	charcount  int
	boolcount  int
	funccount  int
	listcount  int
}

func (a *ActivationRecord) SetRetIp(ip int) {
	a.retip = ip
}

func (a *ActivationRecord) AddNumParam(num float64) {
	addr := mem.Address(a.numcount + mem.Localstart + mem.NumOffset)
	a.numparams = append(a.numparams, NumParam{num, addr})
	a.numcount++
}

func (a *ActivationRecord) AddCharParam(char rune) {
	addr := mem.Address(a.charcount + mem.Localstart + mem.CharOffset)
	a.charparams = append(a.charparams, CharParam{char, addr})
	a.charcount++
}

func (a *ActivationRecord) AddBoolParam(b bool) {
	addr := mem.Address(a.boolcount + mem.Localstart + mem.BoolOffset)
	a.boolparams = append(a.boolparams, BoolParam{b, addr})
	a.boolcount++
}

func (a *ActivationRecord) AddFuncParam(f int) {
	addr := mem.Address(a.funccount + mem.Localstart + mem.FunctionOffset)
	a.funcparams = append(a.funcparams, FuncParam{f, addr})
	a.funccount++
}

func (a *ActivationRecord) AddListParam(l int) {
	addr := mem.Address(a.listcount + mem.Localstart + mem.ListOffset)
	a.listparams = append(a.listparams, ListParam{l, addr})
	a.listcount++
}

func (a *ActivationRecord) AddNumTemp(num float64, addr mem.Address) {
	a.numtemps = append(a.numtemps, NumParam{num, addr})
}

func (a *ActivationRecord) AddCharTemp(char rune, addr mem.Address) {
	a.chartemps = append(a.chartemps, CharParam{char, addr})
}

func (a *ActivationRecord) AddBoolTemp(b bool, addr mem.Address) {
	a.booltemps = append(a.booltemps, BoolParam{b, addr})
}

func (a *ActivationRecord) AddFuncTemp(f int, addr mem.Address) {
	a.functemps = append(a.functemps, FuncParam{f, addr})
}

func (a *ActivationRecord) AddListTemp(l int, addr mem.Address) {
	a.listtemps = append(a.listtemps, ListParam{l, addr})
}

func (a *ActivationRecord) SetRetNum(num float64) {
	a.retnum = num
}

func (a *ActivationRecord) SetRetChar(char rune) {
	a.retchar = char
}

func (a *ActivationRecord) SetRetBool(boolean bool) {
	a.retbool = boolean
}

func (a *ActivationRecord) SetRetFunc(function int) {
	a.retfunc = function
}

func (a *ActivationRecord) SetRetList(list int) {
	a.retlist = list
}

func (a *ActivationRecord) ResetTemps() {
	a.numtemps = make([]NumParam, 0)
	a.chartemps = make([]CharParam, 0)
	a.booltemps = make([]BoolParam, 0)
	a.functemps = make([]FuncParam, 0)
	a.listtemps = make([]ListParam, 0)
}

func (a *ActivationRecord) ResetParams() {
	a.numparams = make([]NumParam, 0)
	a.charparams = make([]CharParam, 0)
	a.boolparams = make([]BoolParam, 0)
	a.funcparams = make([]FuncParam, 0)
	a.listparams = make([]ListParam, 0)
}

func NewActivationRecord() *ActivationRecord {
	return &ActivationRecord{
		0,
		make([]NumParam, 0),
		make([]CharParam, 0),
		make([]BoolParam, 0),
		make([]FuncParam, 0),
		make([]ListParam, 0),
		make([]NumParam, 0),
		make([]CharParam, 0),
		make([]BoolParam, 0),
		make([]FuncParam, 0),
		make([]ListParam, 0),
		0,
		0,
		false,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
	}
}

func (a *ActivationRecord) Listcount() int {
	return a.listcount
}

func (a *ActivationRecord) SetListcount(listcount int) {
	a.listcount = listcount
}

func (a *ActivationRecord) Funccount() int {
	return a.funccount
}

func (a *ActivationRecord) SetFunccount(funccount int) {
	a.funccount = funccount
}

func (a *ActivationRecord) Boolcount() int {
	return a.boolcount
}

func (a *ActivationRecord) SetBoolcount(boolcount int) {
	a.boolcount = boolcount
}

func (a *ActivationRecord) Charcount() int {
	return a.charcount
}

func (a *ActivationRecord) SetCharcount(charcount int) {
	a.charcount = charcount
}

func (a *ActivationRecord) Numcount() int {
	return a.numcount
}

func (a *ActivationRecord) SetNumcount(numcount int) {
	a.numcount = numcount
}

func (a *ActivationRecord) Retlist() int {
	return a.retlist
}

func (a *ActivationRecord) SetRetlist(retlist int) {
	a.retlist = retlist
}

func (a *ActivationRecord) Retfunc() int {
	return a.retfunc
}

func (a *ActivationRecord) SetRetfunc(retfunc int) {
	a.retfunc = retfunc
}

func (a *ActivationRecord) Retbool() bool {
	return a.retbool
}

func (a *ActivationRecord) SetRetbool(retbool bool) {
	a.retbool = retbool
}

func (a *ActivationRecord) Retchar() rune {
	return a.retchar
}

func (a *ActivationRecord) SetRetchar(retchar rune) {
	a.retchar = retchar
}

func (a *ActivationRecord) Retnum() float64 {
	return a.retnum
}

func (a *ActivationRecord) SetRetnum(retnum float64) {
	a.retnum = retnum
}

func (a *ActivationRecord) Listtemps() []ListParam {
	return a.listtemps
}

func (a *ActivationRecord) SetListtemps(listtemps []ListParam) {
	a.listtemps = listtemps
}

func (a *ActivationRecord) Functemps() []FuncParam {
	return a.functemps
}

func (a *ActivationRecord) SetFunctemps(functemps []FuncParam) {
	a.functemps = functemps
}

func (a *ActivationRecord) Booltemps() []BoolParam {
	return a.booltemps
}

func (a *ActivationRecord) SetBooltemps(booltemps []BoolParam) {
	a.booltemps = booltemps
}

func (a *ActivationRecord) Chartemps() []CharParam {
	return a.chartemps
}

func (a *ActivationRecord) SetChartemps(chartemps []CharParam) {
	a.chartemps = chartemps
}

func (a *ActivationRecord) Numtemps() []NumParam {
	return a.numtemps
}

func (a *ActivationRecord) SetNumtemps(numtemps []NumParam) {
	a.numtemps = numtemps
}

func (a *ActivationRecord) Listparams() []ListParam {
	return a.listparams
}

func (a *ActivationRecord) SetListparams(listparams []ListParam) {
	a.listparams = listparams
}

func (a *ActivationRecord) Funcparams() []FuncParam {
	return a.funcparams
}

func (a *ActivationRecord) SetFuncparams(funcparams []FuncParam) {
	a.funcparams = funcparams
}

func (a *ActivationRecord) Boolparams() []BoolParam {
	return a.boolparams
}

func (a *ActivationRecord) SetBoolparams(boolparams []BoolParam) {
	a.boolparams = boolparams
}

func (a *ActivationRecord) Charparams() []CharParam {
	return a.charparams
}

func (a *ActivationRecord) SetCharparams(charparams []CharParam) {
	a.charparams = charparams
}

func (a *ActivationRecord) Numparams() []NumParam {
	return a.numparams
}

func (a *ActivationRecord) SetNumparams(numparams []NumParam) {
	a.numparams = numparams
}

func (a *ActivationRecord) Retip() int {
	return a.retip
}

func (a *ActivationRecord) SetRetip(retip int) {
	a.retip = retip
}
