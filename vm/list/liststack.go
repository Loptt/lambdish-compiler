package list

type node struct {
	val  *ListManager
	next *node
}

type ListStack struct {
	head *node
}

// Empty returns true if ListStack is empty
func (s *ListStack) Empty() bool {
	return s.head == nil
}

// Pop removes the first element in the container
func (s *ListStack) Pop() {
	if s.Empty() {
		return
	}

	s.head = s.head.next
}

// Top returns the first element in the container
func (s *ListStack) Top() *ListManager {
	if s.Empty() {
		return nil
	}
	return s.head.val
}

// Push adds an element to the top of the container
func (s *ListStack) Push(val *ListManager) {
	newHead := &node{val, s.head}
	s.head = newHead
}

func NewListStack() *ListStack {
	return &ListStack{nil}
}
