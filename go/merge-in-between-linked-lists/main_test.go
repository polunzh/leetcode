package mergeinbetweenlinkedlists

import (
	. "leetcode/common"
	"testing"

	"github.com/stretchr/testify/assert" // "github.com/stretchr/testify/assert"
)

func Test1(t *testing.T) {
	l1 := MakeLinkList([]int{0, 1, 2, 3, 4, 5})
	l2 := MakeLinkList([]int{1000000, 1000001, 1000002})

	r := mergeInBetween(l1, 3, 4, l2)
	assert.Equal(t, []int{0, 1, 2, 1000000, 1000001, 1000002, 5}, LinklistToArray(r))
}

func Test2(t *testing.T) {
	l1 := MakeLinkList([]int{0, 1, 2, 3, 4, 5, 6})
	l2 := MakeLinkList([]int{1000000, 1000001, 1000002, 1000003, 1000004})

	r := mergeInBetween(l1, 2, 5, l2)
	assert.Equal(t, []int{0, 1, 1000000, 1000001, 1000002, 1000003, 1000004, 6}, LinklistToArray(r))
}

func Test3(t *testing.T) {
	l1 := MakeLinkList([]int{0, 3, 2, 1, 4, 5})
	l2 := MakeLinkList([]int{1000000, 1000001, 1000002})

	r := mergeInBetween(l1, 3, 4, l2)
	assert.Equal(t, []int{0, 3, 2, 1000000, 1000001, 1000002, 5}, LinklistToArray(r))
}
