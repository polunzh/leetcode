package SameTree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTreeRecursive(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if p == nil || q == nil {
		return false
	}

	if p.Val == q.Val {
		return isSameTreeRecursive(p.Left, q.Left) && isSameTreeRecursive(p.Right, q.Right)
	}

	return false
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSameTree(p *TreeNode, q *TreeNode) bool {
	queue1 := []*TreeNode{p}
	queue2 := []*TreeNode{q}

	for len(queue1) > 0 && len(queue2) > 0 {
		t1 := queue1[0]
		t2 := queue2[0]

		if (t1 == nil && t2 != nil) || (t1 != nil && t2 == nil) {
			return false
		}

		if t1 != nil && t2 != nil {
			if t1.Val != t2.Val {
				return false
			}
		}

		queue1 = queue1[1:]
		if t1 != nil {
			queue1 = append(queue1, t1.Left)
			queue1 = append(queue1, t1.Right)
		}

		queue2 = queue2[1:]
		if t2 != nil {
			queue2 = append(queue2, t2.Left)
			queue2 = append(queue2, t2.Right)
		}
	}

	return len(queue1) == 0 && len(queue2) == 0
}
