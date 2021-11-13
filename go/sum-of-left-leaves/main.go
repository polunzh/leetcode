package sumOfLeftLeaves

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumOfLeftLeaves(root *TreeNode) int {
	result := 0

	if root == nil {
		return 0
	}

	stack := []*TreeNode{}
	stack = append(stack, root)
	for len(stack) != 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if node.Left != nil {
			if node.Left.Left == nil && node.Left.Right == nil {
				result += node.Left.Val
			} else {
				stack = append(stack, node.Left)
			}
		}

		if node.Right != nil {
			stack = append(stack, node.Right)
		}
	}

	return result
}
