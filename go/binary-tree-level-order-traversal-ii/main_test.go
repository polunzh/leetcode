package binarytreelevelordertraversalii

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	tree := TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 9,
		},
		Right: &TreeNode{
			Val: 20,
			Left:  &TreeNode{Val: 15},
			Right: &TreeNode{Val: 7},
		},
	}

	assert.Equal(t, [][]int{[]int{15, 7}, []int{9, 20}, []int{3}}, levelOrderBottom(&tree))
}
