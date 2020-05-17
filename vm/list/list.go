package list

import (
	"fmt"
	"strings"

	"github.com/mewkiz/pkg/errutil"
)

type ListNum struct {
	list []float64
	size int
}

func (l *ListNum) List() []float64 {
	return l.list
}

func (l *ListNum) Size() int {
	return l.size
}

func (l *ListNum) Insert(n float64) {
	l.list = append([]float64{n}, l.list...)
	l.size++
}

func (l *ListNum) Head() (float64, error) {
	if len(l.list) < 1 {
		return 0, errutil.Newf("Attempting to call head on empty list")
	}

	return l.list[0], nil
}

func (l *ListNum) Tail() ([]float64, error) {
	if len(l.list) < 1 {
		return l.list, errutil.Newf("Attempting to call tail on empty list")
	}

	return l.list[1:], nil
}

func (l *ListNum) Copy() *ListNum {
	return &ListNum{l.list, l.size}
}

type ListChar struct {
	list []rune
	size int
}

func (l *ListChar) List() []rune {
	return l.list
}

func (l *ListChar) Size() int {
	return l.size
}

func (l *ListChar) Insert(n rune) {
	l.list = append([]rune{n}, l.list...)
	l.size++
}

func (l *ListChar) Head() (rune, error) {
	if len(l.list) < 1 {
		return 0, errutil.Newf("Attempting to call head on empty list")
	}

	return l.list[0], nil
}

func (l *ListChar) Tail() ([]rune, error) {
	if len(l.list) < 1 {
		return l.list, errutil.Newf("Attempting to call tail on empty list")
	}

	return l.list[1:], nil
}

func (l *ListChar) Copy() *ListChar {
	return &ListChar{l.list, l.size}
}

type ListBool struct {
	list []bool
	size int
}

func (l *ListBool) List() []bool {
	return l.list
}

func (l *ListBool) Size() int {
	return l.size
}

func (l *ListBool) Insert(n bool) {
	l.list = append([]bool{n}, l.list...)
	l.size++
}

func (l *ListBool) Head() (bool, error) {
	if len(l.list) < 1 {
		return false, errutil.Newf("Attempting to call head on empty list")
	}

	return l.list[0], nil
}

func (l *ListBool) Tail() ([]bool, error) {
	if len(l.list) < 1 {
		return l.list, errutil.Newf("Attempting to call tail on empty list")
	}

	return l.list[1:], nil
}

func (l *ListBool) Copy() *ListBool {
	return &ListBool{l.list, l.size}
}

type ListFunc struct {
	list []int
	size int
}

func (l *ListFunc) List() []int {
	return l.list
}

func (l *ListFunc) Size() int {
	return l.size
}

func (l *ListFunc) Insert(n int) {
	l.list = append([]int{n}, l.list...)
	l.size++
}

func (l *ListFunc) Head() (int, error) {
	if len(l.list) < 1 {
		return 0, errutil.Newf("Attempting to call head on empty list")
	}

	return l.list[0], nil
}

func (l *ListFunc) Tail() ([]int, error) {
	if len(l.list) < 1 {
		return l.list, errutil.Newf("Attempting to call tail on empty list")
	}

	return l.list[1:], nil
}

func (l *ListFunc) Copy() *ListFunc {
	return &ListFunc{l.list, l.size}
}

type ListList struct {
	list []*ListManager
	size int
}

func (l *ListList) List() []*ListManager {
	return l.list
}

func (l *ListList) Size() int {
	return l.size
}

func (l *ListList) Insert(n *ListManager) {
	l.list = append([]*ListManager{n}, l.list...)
	l.size++
}

func (l *ListList) Head() (*ListManager, error) {
	if len(l.list) < 1 {
		return nil, errutil.Newf("Attempting to call head on empty list")
	}

	return l.list[0], nil
}

func (l *ListList) Tail() ([]*ListManager, error) {
	if len(l.list) < 1 {
		return l.list, errutil.Newf("Attempting to call tail on empty list")
	}

	return l.list[1:], nil
}

func (l *ListList) Copy() *ListList {
	newlist := make([]*ListManager, 0)

	for _, e := range l.list {
		newlist = append(newlist, e.Copy())
	}
	return &ListList{newlist, l.size}
}

type ListManager struct {
	lnum  *ListNum
	lchar *ListChar
	lbool *ListBool
	lfunc *ListFunc
	llist *ListList
}

func (lm *ListManager) String() string {
	var builder strings.Builder

	builder.WriteString("[")

	if lm.lnum != nil {
		for i, n := range lm.lnum.list {
			if i != len(lm.lnum.list)-1 {
				builder.WriteString(fmt.Sprintf("%f, ", n))
			} else {
				builder.WriteString(fmt.Sprintf("%f", n))
			}
		}
	} else if lm.lchar != nil {
		for i, n := range lm.lchar.list {
			if i != len(lm.lchar.list)-1 {
				builder.WriteString(fmt.Sprintf("%c, ", n))
			} else {
				builder.WriteString(fmt.Sprintf("%c", n))
			}
		}
	} else if lm.lbool != nil {
		for i, n := range lm.lbool.list {
			if i != len(lm.lbool.list)-1 {
				builder.WriteString(fmt.Sprintf("%t, ", n))
			} else {
				builder.WriteString(fmt.Sprintf("%t", n))
			}
		}
	} else if lm.lfunc != nil {
		for i, n := range lm.lfunc.list {
			if i != len(lm.lfunc.list)-1 {
				builder.WriteString(fmt.Sprintf("%d, ", n))
			} else {
				builder.WriteString(fmt.Sprintf("%d", n))
			}
		}
	} else if lm.llist != nil {
		for i, n := range lm.llist.list {
			if i != len(lm.llist.list)-1 {
				builder.WriteString(fmt.Sprintf("%s, ", n))
			} else {
				builder.WriteString(fmt.Sprintf("%s", n))
			}
		}
	}

	builder.WriteString("]")

	return builder.String()
}

func (lm *ListManager) Copy() *ListManager {
	return &ListManager{
		lm.lnum.Copy(),
		lm.lchar.Copy(),
		lm.lbool.Copy(),
		lm.lfunc.Copy(),
		lm.llist.Copy(),
	}
}

func (lm *ListManager) Add(n interface{}) error {
	if f, ok := n.(float64); ok {
		if lm.lnum == nil {
			return errutil.Newf("Cannot set num in non-num list")
		}
		lm.lnum.list = append(lm.lnum.list, f)
		return nil
	} else if c, ok := n.(rune); ok {
		if lm.lchar == nil {
			return errutil.Newf("Cannot set num in non-char list")
		}
		lm.lchar.list = append(lm.lchar.list, c)
		return nil
	} else if b, ok := n.(bool); ok {
		if lm.lbool == nil {
			return errutil.Newf("Cannot set num in non-bool list")
		}

		lm.lbool.list = append(lm.lbool.list, b)
		return nil
	} else if f, ok := n.(int); ok {
		if lm.lfunc == nil {
			return errutil.Newf("Cannot set num in non-func list")
		}
		lm.lfunc.list = append(lm.lfunc.list, f)
		return nil
	} else if l, ok := n.(*ListManager); ok {
		if lm.llist == nil {
			return errutil.Newf("Cannot set num in non-list list")
		}
		lm.llist.list = append(lm.llist.list, l)
		return nil
	}

	return errutil.Newf("Cannot cast element to valid form to add to list")
}

func (lm *ListManager) GetHeadNum() (float64, error) {
	if lm.lnum == nil {
		return 0, errutil.Newf("Cannot get head number from non-number list")
	}

	return lm.lnum.Head()
}

func (lm *ListManager) GetHeadChar() (rune, error) {
	if lm.lchar == nil {
		return 0, errutil.Newf("Cannot get head char from non-char list")
	}

	return lm.lchar.Head()
}

func (lm *ListManager) GetHeadBool() (bool, error) {
	if lm.lbool == nil {
		return false, errutil.Newf("Cannot get head bool from non-bool list")
	}

	return lm.lbool.Head()
}

func (lm *ListManager) GetHeadFunc() (int, error) {
	if lm.lfunc == nil {
		return 0, errutil.Newf("Cannot get head func from non-func list")
	}

	return lm.lfunc.Head()
}

func (lm *ListManager) GetHeadList() (*ListManager, error) {
	if lm.llist == nil {
		return nil, errutil.Newf("Cannot get head list from non-list list")
	}

	return lm.llist.Head()
}

func (lm *ListManager) IsNum() bool {
	return lm.lnum != nil
}

func (lm *ListManager) IsChar() bool {
	return lm.lchar != nil
}

func (lm *ListManager) IsBool() bool {
	return lm.lbool != nil
}

func (lm *ListManager) IsFunc() bool {
	return lm.lfunc != nil
}

func (lm *ListManager) IsList() bool {
	return lm.llist != nil
}

func NewListManager(t int) *ListManager {
	switch t {
	case 1:
		return &ListManager{&ListNum{make([]float64, 0), 0}, nil, nil, nil, nil}
	case 2:
		return &ListManager{nil, &ListChar{make([]rune, 0), 0}, nil, nil, nil}
	case 3:
		return &ListManager{nil, nil, &ListBool{make([]bool, 0), 0}, nil, nil}
	case 4:
		return &ListManager{nil, nil, nil, &ListFunc{make([]int, 0), 0}, nil}
	case 5:
		return &ListManager{nil, nil, nil, nil, &ListList{make([]*ListManager, 0), 0}}
	}

	return nil
}
