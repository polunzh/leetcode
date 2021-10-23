package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var depth int = 0
	var stack []*TreeNode
	stack = append(stack, root)
	for len(stack) > 0 {
		var internalStack []*TreeNode
		depth++
		for _, node := range stack {
			if node.Left != nil {
				internalStack = append(internalStack, node.Left)
			}

			if node.Right != nil{
				internalStack = append(internalStack, node.Right)
			}
		}

		stack = internalStack
	}

	return depth
}

func main() {
	tree1 := &TreeNode{Val: 3}
	tree1.Left = &TreeNode{Val: 9}
	tree1.Right = &TreeNode{Val: 20}
	tree1.Right.Left = &TreeNode{Val: 15}
	tree1.Right.Right = &TreeNode{Val: 7}

	fmt.Println(maxDepth(tree1))
	fmt.Println(maxDepth(nil))
}
