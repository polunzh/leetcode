package relativeSortArray

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRelativeSortArray(t *testing.T) {
	assert.Equal(t, []int{2, 2, 2, 1, 4, 3, 3, 9, 6, 7, 19}, relativeSortArray([]int{2, 3, 1, 3, 2, 4, 6, 7, 9, 2, 19}, []int{2, 1, 4, 3, 9, 6}))
	assert.Equal(t, []int{22, 28, 8, 6, 17, 44}, relativeSortArray([]int{28, 6, 22, 8, 44, 17}, []int{22, 28, 8, 6}))
}
