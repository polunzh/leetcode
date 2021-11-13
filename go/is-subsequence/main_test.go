package IsSubsequence

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsSubsequence(t *testing.T){
	assert.True(t, IsSubsequence("abc", "ahbgdc"))
	assert.True(t, IsSubsequence("abd", "ahbgdc"))
	assert.False(t, IsSubsequence("abdd", "ahbgdc"))
}