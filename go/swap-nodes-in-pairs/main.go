package swapnodesinpairs

type ListNode struct {
	Val  int
	Next *ListNode
}

func treeToArray(head *ListNode) []int {
	var res []int

	cur := head
	for cur != nil {
		res = append(res, cur.Val)
		cur = cur.Next
	}

	return res
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	cur := head
	var pre *ListNode

	for cur != nil && cur.Next != nil {
		t1 := cur
		t2 := cur.Next
		next := t2.Next

		if pre == nil {
			head = t2

			head.Next = t1
			head.Next.Next = next

			pre = cur
		} else {
			pre.Next = t2
			pre.Next.Next = t1
			t1.Next = next
			pre = pre.Next.Next
		}

		cur = next
	}

	return head
}
