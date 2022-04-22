package partitionlist

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	l1 := &ListNode{Val: 1,
		Next: &ListNode{Val: 4,
			Next: &ListNode{Val: 3,
				Next: &ListNode{Val: 2,
					Next: &ListNode{Val: 5,
						Next: &ListNode{Val: 2}}}}}}

	res := partition(l1, 3)
	assert.Equal(t, 1, res.Val)
	assert.Equal(t, 2, res.Next.Val)
	assert.Equal(t, 2, res.Next.Next.Val)
	assert.Equal(t, 4, res.Next.Next.Next.Val)
	assert.Equal(t, 3, res.Next.Next.Next.Next.Val)
	assert.Equal(t, 5, res.Next.Next.Next.Next.Next.Val)

	l2 := &ListNode{Val: 1}
	res2 := partition(l2, 0)
	assert.Equal(t, 1, res2.Val)
}
