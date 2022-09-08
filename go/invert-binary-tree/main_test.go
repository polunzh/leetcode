package invertBinaryTree

import (
	. "leetcode/common"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	root := TreeNode{Val: 2}
	root.Left = &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 3}

	tree := invertTree(&root)
	assert.Equal(t, 2, tree.Val)
	assert.Equal(t, 3, tree.Left.Val)
	assert.Equal(t, 1, tree.Right.Val)
}
