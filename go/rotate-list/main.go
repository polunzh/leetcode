package rotatelist

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
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || k == 0 {
		return head
	}

	fast := head
	slow := head
	count := 1
	for fast.Next != nil {
		fast = fast.Next
		count++
	}

	if k%count == 0 {
		return head
	}

	for i := 0; i < count-k%count-1; i++ {
		slow = slow.Next
	}

	newHead := slow.Next
	fast.Next = head
	slow.Next = nil

	return newHead
}
