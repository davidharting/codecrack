package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStack(t *testing.T) {
	s := NewStack()

	for i := 0; i <= 20; i++ {
		s.Push(i * 2)
	}

	for i := 20; i >= 0; i-- {
		got, err := s.Pop()
		assert.Nil(t, err)
		expected := i * 2
		assert.Equal(t, expected, got)
	}
}
