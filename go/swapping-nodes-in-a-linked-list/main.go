package swappingnodesinalinkedlist

import (
	. "leetcode/common"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// func swapNodes(head *ListNode, k int) *ListNode {
// 	length := 0
// 	cur := head
// 	for cur != nil {
// 		length++
// 		cur = cur.Next
// 	}

// 	if length == 1 || k < 1 || k > length {
// 		return head
// 	}

// 	var firstNode, lastNode *ListNode
// 	j := length - k + 1
// 	cur = head

// 	index := 1
// 	for ; cur != nil; index++ {
// 		if index == k {
// 			firstNode = cur
// 		}

// 		if index == j {
// 			lastNode = cur
// 		}

// 		if firstNode != nil && lastNode != nil {
// 			break
// 		}

// 		cur = cur.Next
// 	}

// 	firstNode.Val, lastNode.Val = lastNode.Val, firstNode.Val

// 	return head
// }

func swapNodes(head *ListNode, k int) *ListNode {
	cur := head
	firstNode := head
	lastNode := head

	for ; k > 1; k-- {
		if cur == nil {
			return head
		}

		cur = cur.Next
	}

	if cur == nil {
		return head
	}

	firstNode = cur

	for cur.Next != nil {
		cur = cur.Next
		lastNode = lastNode.Next
	}

	firstNode.Val, lastNode.Val = lastNode.Val, firstNode.Val
	return head
}
