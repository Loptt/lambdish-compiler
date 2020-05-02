package ic

// node is a data container
type node struct {
	val  int
	next *node
}

// FuncEntryStack implements a stack for the FuncEntry data type
type JumpStack struct {
	head *node
}

// Empty returns true if JumpStack is empty
func (s *JumpStack) Empty() bool {
	return s.head == nil
}

// Pop removes the first element in the container
func (s *JumpStack) Pop() {
	if s.Empty() {
		return
	}

	s.head = s.head.next
}

// Top returns the first element in the container
func (s *JumpStack) Top() int {
	if s.Empty() {
		return 0
	}
	return s.head.val
}

// Push adds an element to the top of the container
func (s *JumpStack) Push(val int) {
	newHead := &node{val, s.head}
	s.head = newHead
}

func NewJumpStack() *JumpStack {
	return &JumpStack{nil}
}
