package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	if root.Left == nil {
		return minDepth(root.Right) + 1
	}

	if root.Right == nil {
		return minDepth(root.Left) + 1
	}

	return min(minDepth(root.Left), minDepth(root.Right)) + 1
}

func min(x, y int) int {
	if x < y {
		return x
	}

	return y
}

func main() {
	node := &TreeNode{Val: 3}
	node.Left = &TreeNode{Val: 9}
	node.Right = &TreeNode{Val: 20}
	node.Right.Left = &TreeNode{Val: 15}
	node.Right.Right = &TreeNode{Val: 7}

	fmt.Println(minDepth(node))
}
