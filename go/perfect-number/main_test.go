package perfectnumber

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	assert.Equal(t, true, checkPerfectNumber(28))
	assert.Equal(t, false, checkPerfectNumber(7))
	assert.Equal(t, false, checkPerfectNumber(1))
}