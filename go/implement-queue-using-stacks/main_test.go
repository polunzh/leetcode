package implementqueueusingstacks

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test1(t *testing.T) {
	obj := Constructor()
	obj.Push(1)
	obj.Push(2)

	v := obj.Peek()
	assert.Equal(t, 1, v)
	v = obj.Pop()
	assert.Equal(t, 1, v)

	empty := obj.Empty()
	assert.Equal(t, false, empty)
}

func Test2(t *testing.T) {
	obj := Constructor()
	obj.Push(1)
	obj.Push(2)

	v := obj.Pop()
	assert.Equal(t, 1, v)

	v = obj.Pop()
	assert.Equal(t, 2, v)

	empty := obj.Empty()
	assert.Equal(t, true, empty)
}