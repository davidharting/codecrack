package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func createList(values []int) SinglyLinkedList {
	l := CreateSinglyLinkedList(values[0])
	for _, n := range values[1:] {
		l.Append(n)
	}
	return l
}

func sliceEq(s1 []int, s2 []int) bool {
	if len(s1) != len(s2) {
		return false
	}

	for i, n := range s1 {
		if n != s2[i] {
			return false
		}
	}

	return true
}

func TestToSlice(t *testing.T) {
	list := CreateSinglyLinkedList(3)
	list.Append(99)
	list.Append(31)
	list.Append(4321)
	got := list.ToSlice()
	expected := []int{3, 99, 31, 4321}
	if len(got) != len(expected) {
		t.Errorf("Expected %v, got %v", expected, got)
	}

	for i, g := range got {
		e := expected[i]
		if g != e {
			t.Errorf("Expected %v, got %v", expected, got)
			break
		}
	}
}

func TestAppend(t *testing.T) {
	list := CreateSinglyLinkedList(4)
	list.Append(99)
	list.Append(42)
	got := list.ToSlice()
	expected := []int{4, 99, 42}
	if !sliceEq(got, expected) {
		t.Errorf("Expected %v, got %v", expected, got)
	}
}

func TestGet(t *testing.T) {
	values := []int{100, 98, 86, 200, 33, 1001}

	l := createList(values)

	for i, v := range values {
		got, err := l.Get(i)
		assert.Nil(t, err, "There should not be an error")
		assert.Equal(t, v, got, "Retrieved element should match expected")
	}

	got, err := l.Get(6)
	assert.Zero(t, got, "In error condition, value should be 0")
	assert.EqualError(t, err, "Index out of bounds", "Should return Index out of bounds error if the index is too high")
}

func TestLength(t *testing.T) {
	assert.Zero(t, SinglyLinkedList{})

	l := createList([]int{1})
	assert.Equal(t, 1, l.Length())

	l = createList([]int{1, 2, 3, 4})
	assert.Equal(t, 4, l.Length())
}
