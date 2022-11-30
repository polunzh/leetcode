package plusone

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddStrings(t *testing.T) {
	assert.Equal(t, []int{1, 2, 4}, plusOne([]int{1, 2, 3}))
	assert.Equal(t, []int{4, 3, 2, 2}, plusOne([]int{4, 3, 2, 1}))
	assert.Equal(t, []int{1}, plusOne([]int{0}))
	assert.Equal(t, []int{1, 0}, plusOne([]int{9}))
	assert.Equal(t, []int{1, 0, 0}, plusOne([]int{9, 9}))
	assert.Equal(t, []int{9, 0, 0, 0}, plusOne([]int{8, 9, 9, 9}))
}
