package ar

type node struct {
	val  *ActivationRecord
	next *node
}

type ArStack struct {
	head *node
}

// Empty returns true if ArStack is empty
func (s *ArStack) Empty() bool {
	return s.head == nil
}

// Pop removes the first element in the container
func (s *ArStack) Pop() {
	if s.Empty() {
		return
	}

	s.head = s.head.next
}

// Top returns the first element in the container
func (s *ArStack) Top() *ActivationRecord {
	if s.Empty() {
		return nil
	}
	return s.head.val
}

// Push adds an element to the top of the container
func (s *ArStack) Push(val *ActivationRecord) {
	newHead := &node{val, s.head}
	s.head = newHead
}

func NewArStack() *ArStack {
	return &ArStack{nil}
}
