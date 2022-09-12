package binarytreepostordertraversal

import (
	. "leetcode/common"
)

func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	res := []int{}

	res = append(res, postorderTraversal(root.Left)...)
	res = append(res, postorderTraversal(root.Right)...)
	res = append(res, root.Val)

	return res
}

func postorderTraversal2(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	stack := []*TreeNode{}
	res := []int{}

	cur := root

	for cur != nil || len(stack) > 0 {
		if cur != nil {
			if cur.Right != nil {
				stack = append(stack, cur.Right, cur)
				cur = cur.Left
				continue
			} else if cur.Left != nil {
				stack = append(stack, cur)
				cur = cur.Left
				continue
			} else {
				stack = append(stack, cur)
				cur = nil
			}
		}

		if cur == nil {
			cur = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}

		if cur.Right != nil && (len(stack) > 0 && cur.Right == stack[len(stack)-1]) {
			stack = stack[:len(stack)-1]
			stack = append(stack, cur)
			cur = cur.Right
		} else if cur.Left == nil {
			res = append(res, cur.Val)
			cur = nil
		} else {
			res = append(res, cur.Val)
			cur = nil
		}
	}

	return res
}
