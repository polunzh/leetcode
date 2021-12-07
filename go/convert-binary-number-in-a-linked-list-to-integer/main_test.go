package convertBinaryNumberInALinkedListToInteger

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetDecimalValue(t *testing.T) {
	t1 := &ListNode{Val: 1, Next: nil}
	assert.Equal(t, 1, getDecimalValue(t1))
	t2 := &ListNode{Val: 1, Next: &ListNode{Val: 0, Next: &ListNode{Val: 1, Next: nil}}}
	assert.Equal(t, 5, getDecimalValue(t2))
}
