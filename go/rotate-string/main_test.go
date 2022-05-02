package rotatestring

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	assert.Equal(t, true, rotateString("abcde", "cdeab"))
	assert.Equal(t, false, rotateString("abcde", "abced"))
	assert.Equal(t, false, rotateString("abcde", "cdeabe"))
}
