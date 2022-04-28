package countingbits

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	assert.Equal(t, []int{0}, countBits(0))
	assert.Equal(t, []int{0, 1, 1}, countBits(2))
	assert.Equal(t, []int{0,1,1,2,1,2}, countBits(5))
}