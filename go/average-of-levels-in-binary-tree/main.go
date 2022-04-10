package averageoflevelsinbinarytree

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
func averageOfLevels(root *TreeNode) []float64 {
	var result []float64

	if root == nil {
		return result
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

		// result = append(result, )
		var total int
		for _, v := range r {
			total += v
		}

		result = append(result, float64(total)/float64(len(r)))
		nodes = tempNodes
	}

	return result
}
