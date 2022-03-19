package powerOfThree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsPowerOfThree(t *testing.T) {
	assert.Equal(t, false, isPowerOfThree(0))
	assert.Equal(t, true, isPowerOfThree(1))
	assert.Equal(t, true, isPowerOfThree(3))
	assert.Equal(t, false, isPowerOfThree(10))
	assert.Equal(t, true, isPowerOfThree(27))
}
