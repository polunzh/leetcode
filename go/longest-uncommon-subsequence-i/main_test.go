package longestuncommonsubsequencei

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	assert.Equal(t, 3, findLUSlength("aba", "cdc"))
	assert.Equal(t, 3, findLUSlength("aaa", "bbb"))
	assert.Equal(t, -1, findLUSlength("aaa", "aaa"))
}