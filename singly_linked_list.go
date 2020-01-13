package main

import (
	"errors"
	"fmt"
)

type node struct {
	value int
	next  *node
}

// SinglyLinkedList is a list implementation with links in only the forward direction
type SinglyLinkedList struct {
	first *node
}

// Get retrieves the node at the index of a list
// It will return nil if the index is out of bounds
func (l SinglyLinkedList) get(index int) *node {
	if index < 0 {
		return nil
	}

	n := l.first

	i := 0

	for n != nil {
		if i == index {
			return n
		}
		n = n.next
		i++
	}

	return nil
}

func (l SinglyLinkedList) last() *node {
	n := l.first
	for n.next != nil {
		n = n.next
	}
	return n
}

// CreateSinglyLinkedList constructs a new SinglyLinkedList with a first value
func CreateSinglyLinkedList(firstValue int) SinglyLinkedList {
	return SinglyLinkedList{&node{firstValue, nil}}
}

// NewSinglyLinkedList constructs an empty list
func NewSinglyLinkedList() SinglyLinkedList {
	return SinglyLinkedList{nil}
}

// Append adds an element to the end of a list
func (l *SinglyLinkedList) Append(value int) {
	if l.IsEmpty() {
		fmt.Printf("l.isEmpty(). Assigning first.\n")
		l.first = &node{value, nil}
		fmt.Printf("l.first %v\n", l.first)
	} else {
		n := l.last()
		n.next = &node{value, nil}
	}
}

// ToSlice converts the list to a slice
func (l SinglyLinkedList) ToSlice() []int {
	var s []int

	n := l.first

	for n.next != nil {
		s = append(s, n.value)
		n = n.next
	}

	// Didn't get an elegant way to get the last value inside the loop
	s = append(s, n.value)

	return s
}

// Get retrieves the element at the index of a list
func (l SinglyLinkedList) Get(index int) (int, error) {
	n := l.get(index)

	if n == nil {
		return 0, errors.New("Index out of bounds")
	}

	return n.value, nil
}

// Length returns the number of elements in a linked list
func (l SinglyLinkedList) Length() int {
	n := l.first
	length := 0

	for n != nil {
		length++
		n = n.next
	}

	return length
}

// IsEmpty returns true if there are no items in the list
func (l SinglyLinkedList) IsEmpty() bool {
	return l.first == nil
}

// DeleteAtIndex removes an element at the specified index
func (l *SinglyLinkedList) DeleteAtIndex(index int) (int, error) {
	length := l.Length()
	if index >= length || index < 0 {
		return 0, errors.New("Index out of bounds")
	}

	if length == 1 {
		n := l.first
		val := n.value
		l.first = nil
		return val, nil
	}

	reconnectMe := l.get(index - 1)
	deleteMe := l.get(index)
	reconnectMe.next = deleteMe.next
	return deleteMe.value, nil
}

// RemoveTail deletes the last element from the list and returns its value
func (l *SinglyLinkedList) RemoveTail() (int, error) {
	return l.DeleteAtIndex(l.Length() - 1)
}
