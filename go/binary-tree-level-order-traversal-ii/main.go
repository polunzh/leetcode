package binarytreelevelordertraversalii

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrderBottom(root *TreeNode) [][]int {
	var tempResult [][]int

	if root == nil {
		return tempResult
	}

	var nodes []*TreeNode = []*TreeNode{root}
	for len(nodes) != 0 {
		var r []int
		var tempNodes []*TreeNode
		for _, n := range nodes {
			r = append(r, n.Val)
			if n.Left != nil {
				tempNodes = append(tempNodes, n.Left)
			}

			if n.Right != nil {
				tempNodes = append(tempNodes, n.Right)
			}
		}

		tempResult = append(tempResult, r)
		nodes = tempNodes
	}

	var result [][]int
	for i := len(tempResult) - 1; i >= 0; i-- {
		result = append(result, tempResult[i])
	}

	return result
}
