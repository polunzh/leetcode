package validparentheses

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test1(t *testing.T) {
	assert.Equal(t, true, isValid("()"))
}

func Test2(t *testing.T) {
	assert.Equal(t, true, isValid("()[]{}"))
}

func Test3(t *testing.T) {
	assert.Equal(t, false, isValid("(]"))
}

func Test4(t *testing.T) {
	assert.Equal(t, false, isValid("(])"))
}