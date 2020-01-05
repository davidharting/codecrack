package main

import "errors"

type node struct {
	value int
	next  *node
}

// SinglyLinkedList is a list implementation with links in only the forward direction
type SinglyLinkedList struct {
	first *node
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

// Append adds an element to the end of a list
func (l SinglyLinkedList) Append(value int) {
	n := l.last()
	n.next = &node{value, nil}
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
	if index < 0 {
		return 0, errors.New("Index out of bounds")
	}

	n := l.first

	i := 0

	for n != nil {
		if i == index {
			return n.value, nil
		}
		n = n.next
		i++
	}

	return 0, errors.New("Index out of bounds")
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
