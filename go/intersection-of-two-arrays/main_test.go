package intersectionOfTwoArrays

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIntersection(t *testing.T) {
	assert.Equal(t, []int{2}, intersection([]int{1, 2, 2, 1}, []int{2, 2}))
	assert.Equal(t, []int{4, 9}, intersection([]int{4, 9, 5}, []int{9, 4, 9, 8, 4}))
}
