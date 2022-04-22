package partitionlist

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition(head *ListNode, x int) *ListNode {
	l1, l2 := new(ListNode), new(ListNode)
	c1, c2 := l1, l2

	for head != nil {
		if head.Val < x {
			c1.Next = head
			c1 = c1.Next
		} else {
			c2.Next = head
			c2 = c2.Next
		}

		head = head.Next
	}

	c2.Next = nil
	c1.Next = l2.Next

	return l1.Next
}
