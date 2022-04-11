package detectcapital

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	assert.Equal(t, true, detectCapitalUse("USA"))
	assert.Equal(t, true, detectCapitalUse("leetcode"))
	assert.Equal(t, true, detectCapitalUse("Google"))
	assert.Equal(t, false, detectCapitalUse("FlaG"))
}