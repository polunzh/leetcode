package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 */
func inorderTraversal(root *TreeNode) []int {
	var res []int

	if root != nil {
		res = append(res, inorderTraversal(root.Left)...)
		res = append(res, root.Val)
		res = append(res, inorderTraversal(root.Right)...)
	}

	return res
}

func main() {
	node := TreeNode{Val: 1}
	node.Right = &TreeNode{Val: 2}
	node.Right.Left = &TreeNode{Val: 3}

	fmt.Println(inorderTraversal(&node))
}
