package swapnodesinpairs

import (
	"testing"
	. "leetcode/common"

	"github.com/stretchr/testify/assert"
)

func Test1(t *testing.T) {
	t1 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 5,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 6,
						Next: &ListNode{
							Val: 2,
							Next: &ListNode{
								Val: 2,
							},
						},
					},
				},
			},
		},
	}

	r1 := swapPairs(t1)
	assert.Equal(t, []int{5, 2, 4, 3, 2, 6, 2}, LinklistToArray(r1))
}

func Test2(t *testing.T) {
	t1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
			},
		},
	}

	r1 := swapPairs(t1)
	assert.Equal(t, []int{2, 1, 3}, LinklistToArray(r1))
}

func Test3(t *testing.T) {
	t1 := &ListNode{
		Val: 1,
	}

	r1 := swapPairs(t1)
	assert.Equal(t, []int{1}, LinklistToArray(r1))
}

func Test4(t *testing.T) {
	var tree *ListNode
	r1 := swapPairs(tree)
	assert.Equal(t, tree, r1)
}
