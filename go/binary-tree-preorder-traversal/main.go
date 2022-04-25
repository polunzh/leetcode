package binarytreepreordertraversal

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// for test
func print(nodes []*TreeNode) {
	for i := 0; i < len(nodes); i++ {
		fmt.Printf("%d,", nodes[i].Val)
	}
	fmt.Println()
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 递归版本
// func preorderTraversal(root *TreeNode) []int {
// 	if root == nil {
// 		return []int{}
// 	}

// 	res := []int{root.Val}
// 	if root.Left != nil {
// 		res = append(res, preorderTraversal(root.Left)...)
// 	}

// 	if root.Right != nil {
// 		res = append(res, preorderTraversal(root.Right)...)
// 	}

// 	return res
// }

func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	res := []int{}
	stack := []*TreeNode{root}

	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[0 : len(stack)-1]
		res = append(res, node.Val)

		if (*node).Right != nil {
			stack = append(stack, (*node).Right)
		}

		if (*node).Left != nil {
			stack = append(stack, (*node).Left)
		}
	}

	return res
}
