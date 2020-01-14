package concurrency

import "testing"

import "github.com/stretchr/testify/assert"

func TestSum(t *testing.T) {
	sum := Sum([]int{2, 50, 99, 14, 21})
	assert.Equal(t, 186, sum, "It should sum a list of numbers")

	sum = Sum([]int{})
	assert.Equal(t, 0, sum, "An empty list should yield 0")

	longList := []int{}
	for i := 0; i < 10000; i++ {
		longList = append(longList, 3)
	}
	assert.Equal(t, 30000, Sum(longList), "It should handle long lists.")
}
