package mergeinbetweenlinkedlists

import (
	. "leetcode/common"
)

func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
	head := list1
	index := 0

	for head != nil {
		if head.Next != nil && index == a-1 {
			tmp := head.Next
			head.Next = list2

			for list2 != nil {
				head = list2
				list2 = list2.Next
			}

			for tmp != nil && index != b-1 {
				tmp = tmp.Next
				index++
			}

			head.Next = tmp.Next
		}

		head = head.Next
		index++
	}

	return list1
}
