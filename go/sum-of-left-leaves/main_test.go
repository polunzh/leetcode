package sumOfLeftLeaves

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSumOfLeftLeaves(t *testing.T) {
	tree1 := TreeNode{
		Val:   3,
		Left:  &TreeNode{Val: 9},
		Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}},
	}

	tree2 := TreeNode{
		Val: 3,
	}

	tree3 := TreeNode{
		Val: 3,
		Left:  &TreeNode{Val: 9},
	}

	assert.Equal(t, 24, sumOfLeftLeaves(&tree1))
	assert.Equal(t, 0, sumOfLeftLeaves(&tree2))
	assert.Equal(t, 9, sumOfLeftLeaves(&tree3))
	assert.Equal(t, 0, sumOfLeftLeaves(nil))
}
