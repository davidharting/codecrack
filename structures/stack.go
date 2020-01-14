package structures

import "errors"

// Stack provides a last-in-first-out in data structure
type Stack struct {
	list SinglyLinkedList
}

// NewStack creates a new stack instance, backed by a linked-list
// You may omit passing an integer, in which case the stack will be empty
func NewStack() Stack {
	list := NewSinglyLinkedList()
	return Stack{list}
}

// Push adds an element to the Stack
// It will be the first out if no further Pushes are made
func (s *Stack) Push(value int) {
	s.list.Append(value)
}

// Pop removes and returns the latest item from the Stack
// If there are no items in the stack, then an error is returned
func (s *Stack) Pop() (int, error) {
	if s.list.IsEmpty() {
		return 0, errors.New("cannot Pop an empty Stack")
	}

	return s.list.RemoveTail()
}
