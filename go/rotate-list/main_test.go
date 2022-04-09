package rotatelist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	l := &ListNode{Val: 1, Next: nil}
	l.Next = &ListNode{Val: 2, Next: nil}
	l.Next.Next = &ListNode{Val: 3, Next: nil}
	l.Next.Next.Next = &ListNode{Val: 4, Next: nil}
	l.Next.Next.Next.Next = &ListNode{Val: 5, Next: nil}

	rl := rotateRight(l, 3)
	assert.Equal(t, 3, rl.Val)
	assert.Equal(t, 4, rl.Next.Val)
	assert.Equal(t, 5, rl.Next.Next.Val)
	assert.Equal(t, 1, rl.Next.Next.Next.Val)
	assert.Equal(t, 2, rl.Next.Next.Next.Next.Val)
}

func Test1(t *testing.T) {
	l := &ListNode{Val: 0, Next: nil}
	l.Next = &ListNode{Val: 1, Next: nil}
	l.Next.Next = &ListNode{Val: 2, Next: nil}

	rl := rotateRight(l, 4)
	assert.Equal(t, 2, rl.Val)
	assert.Equal(t, 0, rl.Next.Val)
	assert.Equal(t, 1, rl.Next.Next.Val)
}

func Test2(t *testing.T) {
	l := &ListNode{Val: 0, Next: nil}
	l.Next = &ListNode{Val: 1, Next: nil}
	l.Next.Next = &ListNode{Val: 2, Next: nil}

	rl := rotateRight(l, 3)
	assert.Equal(t, 0, rl.Val)
	assert.Equal(t, 1, rl.Next.Val)
	assert.Equal(t, 2, rl.Next.Next.Val)
}

func Test3(t *testing.T) {
	l := &ListNode{Val: 1, Next: nil}

	rl := rotateRight(l, 0)
	assert.Equal(t, 1, rl.Val)
}