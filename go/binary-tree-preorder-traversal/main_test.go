package binarytreepreordertraversal

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test1(t *testing.T) {
	t1 := &TreeNode{Val: 1, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}}
	assert.Equal(t, []int{1, 2, 3}, preorderTraversal(t1))

}

func Test2(t *testing.T) {
	t1 := &TreeNode{Val: 3, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 2}}
	assert.Equal(t, []int{3, 1, 2}, preorderTraversal(t1))
}

func Test3(t *testing.T) {
	assert.Equal(t, []int{}, preorderTraversal(nil))
}

func Test4(t *testing.T) {
	t1 := &TreeNode{Val: 1, Left: &TreeNode{Val: 4, Left: &TreeNode{Val: 2}}, Right: &TreeNode{Val: 3}}
	assert.Equal(t, []int{1, 4, 2, 3}, preorderTraversal(t1))
}
