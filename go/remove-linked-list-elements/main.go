package main

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	var current *ListNode = head
	var pre *ListNode
	var header = pre

	for current != nil {
		if current.Val == val {
			if current.Next == nil {
				if pre == nil {
					return header // 第一个元素是目标元素
				}

				pre.Next = nil
				break
			}

			if current.Next.Val != val {
				if pre == nil {
					header = current.Next
				} else {
					pre.Next = current.Next
				}
			}

			current = current.Next
			continue
		}

		pre = current
		current = current.Next
		if header == nil {
			header = pre
		}
	}

	return header
}

func print(node *ListNode) {
	for node != nil {
		fmt.Println(node.Val)
		node = node.Next
	}
}

func main() {
	head := &ListNode{Val: 6}
	head.Next = &ListNode{Val: 6}
	head.Next.Next = &ListNode{Val: 6}
	head.Next.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next.Next = &ListNode{Val: 4}
	head.Next.Next.Next.Next.Next = &ListNode{Val: 5}
	head.Next.Next.Next.Next.Next.Next = &ListNode{Val: 6}

	print(removeElements(head, 6))
	// removeElements(head, 6)
}
