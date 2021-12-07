package convertBinaryNumberInALinkedListToInteger

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
func getDecimalValue(head *ListNode) int {
	if head == nil {
		return 0
	}

	result := 0
	for head != nil {
		result = (2*result + head.Val)
		head = head.Next
	}

	return result
}
