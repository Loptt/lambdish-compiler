package dir

// Node is a data container
type Node struct {
	val  *FuncEntry
	next *Node
}

// FuncEntryStack implements a stack for the FuncEntry data type
type FuncEntryStack struct {
	head *Node
}

// Empty returns true if FuncEntryStack is empty
func (s *FuncEntryStack) Empty() bool {
	return s.head == nil
}

// Pop removes the first element in the container
func (s *FuncEntryStack) Pop() {
	if s.Empty() {
		return
	}

	s.head = s.head.next
}

// Top returns the first element in the container
func (s *FuncEntryStack) Top() *FuncEntry {
	if s.Empty() {
		return nil
	}
	return s.head.val
}

// Push adds an element to the top of the container
func (s *FuncEntryStack) Push(val *FuncEntry) {
	newHead := &Node{val, s.head}
	s.head = newHead
}

func NewFuncEntryStack() *FuncEntryStack {
	return &FuncEntryStack{nil}
}