package increasingordersearchtree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func increasingBST(root *TreeNode) *TreeNode {
	var head *TreeNode
	var pre *TreeNode // pre 是新树的 pre 节点
	var cur *TreeNode = root
	var stack []*TreeNode

	for cur != nil || len(stack) > 0 {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}

		cur = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if pre != nil {
			pre.Right = cur
		} else {
			head = cur
		}

		pre = cur
		cur.Left = nil
		cur = cur.Right
	}

	return head
}
