package findAllNumbersDisappearedInAnArray

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsSubsequence(t *testing.T) {
	assert.Equal(t, []int{5, 6}, findDisappearedNumbers([]int{4, 3, 2, 7, 8, 2, 3, 1}))
	assert.Equal(t, []int{2}, findDisappearedNumbers([]int{1, 1}))
}
