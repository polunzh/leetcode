package powxn

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	assert.Equal(t, 1024.00000, myPow(2.00000, 10))
	assert.Equal(t, true, myPow(2.10000, -3)-0.10798 < 0.0001)
}
