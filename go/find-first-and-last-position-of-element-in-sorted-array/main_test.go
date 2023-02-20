package findfirstandlastpositionofelementinsortedarray

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	assert.Equal(t, []int{-1, -1}, searchRange([]int{}, 0))
	assert.Equal(t, []int{-1, -1}, searchRange([]int{5, 7, 7, 8, 8, 10}, 6))
	assert.Equal(t, []int{3, 4}, searchRange([]int{5, 7, 7, 8, 8, 10}, 8))
}

func TestBinarySearch(t *testing.T) {
	assert.Equal(t, 2, binarySearch([]int{1, 2, 3, 5, 6}, 0, 3))
	assert.Equal(t, 2, binarySearch([]int{1, 2, 3, 5, 6}, 2, 3))
	assert.Equal(t, 1, binarySearch([]int{1, 2, 2, 5, 6}, 0, 2))
	assert.Equal(t, 5, binarySearch([]int{1, 2, 2, 5, 6}, 0, 8))
	assert.Equal(t, 0, binarySearch([]int{2, 2, 2, 5, 6}, 0, 1))
}
