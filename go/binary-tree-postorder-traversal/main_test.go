package binarytreepostordertraversal

import (
	"github.com/stretchr/testify/assert"
	. "leetcode/common"
	"testing"
)

func Test1(t *testing.T) {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{Val: 2,
			Left:  &TreeNode{Val: 4},
			Right: &TreeNode{Val: 5},
		},
		Right: &TreeNode{Val: 3,
			Left:  &TreeNode{Val: 6},
			Right: &TreeNode{Val: 7},
		},
	}
	assert.Equal(t, []int{4, 5, 2, 6, 7, 3, 1}, postorderTraversal2(root))
}

func Test2(t *testing.T) {
	root := &TreeNode{
		Val: 1,
		Right: &TreeNode{Val: 2,
			Left: &TreeNode{Val: 3},
		},
	}
	assert.Equal(t, []int{3, 2, 1}, postorderTraversal2(root))
}
