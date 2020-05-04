//Package ic provides the generation of intermediate code
package ic

import (
	"fmt"
	"strings"

	"github.com/Loptt/lambdish-compiler/mem"
)

// node is a data container
type node struct {
	val  mem.Address
	next *node
}

// AddressStack implements a stack for the FuncEntry data type
type AddressStack struct {
	head *node
}

// Empty returns true if AddressStack is empty
func (s *AddressStack) Empty() bool {
	return s.head == nil
}

// String returns the string representation of the stack
func (s *AddressStack) String() string {
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
func (s *AddressStack) Pop() {
	if s.Empty() {
		return
	}

	s.head = s.head.next
}

// Top returns the first element in the container
func (s *AddressStack) Top() mem.Address {
	if s.Empty() {
		return 0
	}
	return s.head.val
}

// Push adds an element to the top of the container
func (s *AddressStack) Push(val mem.Address) {
	newHead := &node{val, s.head}
	s.head = newHead
}

// NewAddressStack ...
func NewAddressStack() *AddressStack {
	return &AddressStack{nil}
}
