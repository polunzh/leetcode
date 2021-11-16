package addStrings

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddStrings(t *testing.T) {
	assert.Equal(t, "10", addStrings("1", "9"))
	assert.Equal(t, "10", addStrings("9", "1"))
	assert.Equal(t, "134", addStrings("11", "123"))
	assert.Equal(t, "533", addStrings("456", "77"))
	assert.Equal(t, "0", addStrings("0", "0"))
	assert.Equal(t, "1008", addStrings("999", "9"))
}
