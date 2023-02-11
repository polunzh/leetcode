package increasingordersearchtree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	root := TreeNode{Val: 5, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 7}}

	tree := increasingBST(&root)

	assert.Equal(t, 1, tree.Val)
	assert.Equal(t, 5, tree.Right.Val)
	assert.Equal(t, 7, tree.Right.Right.Val)
}
