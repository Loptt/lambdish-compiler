package ic

import (
	"fmt"
	"strings"
)

// nodeint is a data container
type nodeint struct {
	val  int
	next *nodeint
}

// FuncEntryStack implements a stack for the FuncEntry data type
type JumpStack struct {
	head *nodeint
}

// Empty returns true if JumpStack is empty
func (s *JumpStack) Empty() bool {
	return s.head == nil
}

// String returns the string representation of the stack
func (s *JumpStack) String() string {
	if s.Empty() {
		return "<Empty Stack>"
	}

	curr := s.head
	var result strings.Builder

	result.WriteString("< ")

	for curr != nil {
		value := fmt.Sprintf("%v", curr.val)
		if curr.next == nil {
			result.WriteString(value)
		} else {
			result.WriteString(value + ", ")
		}
		curr = curr.next
	}

	result.WriteString(" >")

	return result.String()
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
	newHead := &nodeint{val, s.head}
	s.head = newHead
}

func NewJumpStack() *JumpStack {
	return &JumpStack{nil}
}
