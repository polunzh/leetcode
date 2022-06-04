package stringtointegeratoi

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test1(t *testing.T) {
	assert.Equal(t, 42, myAtoi("42"))
}

func Test2(t *testing.T) {
	assert.Equal(t, -42, myAtoi("-42"))
}

func Test3(t *testing.T) {
	assert.Equal(t, 4193, myAtoi("4193 with words"))
}

func Test4(t *testing.T) {
	assert.Equal(t, -42, myAtoi("   -42"))
}

func Test5(t *testing.T) {
	assert.Equal(t, -2147483648, myAtoi("-91283472332"))
}

func Test6(t *testing.T) {
	assert.Equal(t, 2147483647, myAtoi("91283472332"))
}

func Test7(t *testing.T) {
	assert.Equal(t, 0, myAtoi("00000-42a1234"))
}

func Test8(t *testing.T) {
	assert.Equal(t, 12345678, myAtoi("  0000000000012345678"))
}

func Test9(t *testing.T) {
	assert.Equal(t, 0, myAtoi("  0000000000-012345678"))
}

func Test10(t *testing.T) {
	assert.Equal(t, -5, myAtoi("-5-"))
}