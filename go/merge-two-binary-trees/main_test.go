package mergetwobinarytrees

import (
	. "leetcode/common"
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	root1 := TreeNode{Val: 2}
	root1.Left = &TreeNode{Val: 1}
	root1.Right = &TreeNode{Val: 3}

	root2 := TreeNode{Val: 2}
	// root2.Left = &TreeNode{Val: 1}
	root2.Right = &TreeNode{Val: 3}

	root := mergeTrees(&root1, &root2)
	assert.Equal(t, 4, root.Val)
	assert.Equal(t, 1, root.Left.Val)
	assert.Equal(t, 6, root.Right.Val)
}
