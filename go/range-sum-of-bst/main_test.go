package rangeSumOfBst

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRangeSumBST(t *testing.T){
	tree1 := &TreeNode{
		Val: 10,
		Left: &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val:3,
			},
			Right: &TreeNode{
				Val:7,
			},
		},
		Right: &TreeNode{
			Val: 15,
			Right: &TreeNode{
				Val: 18,
			},
		},
	}

	assert.Equal(t, 0, rangeSumBST(nil, 7, 15))
	assert.Equal(t, 32, rangeSumBST(tree1, 7, 15))
}