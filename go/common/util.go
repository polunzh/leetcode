package common

type ListNode struct {
	Val  int
	Next *ListNode
}

func LinklistToArray(head *ListNode) []int {
	var res []int

	cur := head
	for cur != nil {
		res = append(res, cur.Val)
		cur = cur.Next
	}

	return res
}

func MakeLinkList(arr []int) *ListNode {
	if len(arr) == 0 {
		return nil
	}

	pre := &ListNode{Val: arr[0]}
	head := pre
	for i := 1; i < len(arr); i++ {
		n := &ListNode{Val: arr[i]}
		pre.Next = n
		pre = pre.Next
	}

	return head
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
