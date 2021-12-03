package bitwiseComplement

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBitwiseComplement(t *testing.T) {
	assert.Equal(t, 2, bitwiseComplement(5))
	assert.Equal(t, 0, bitwiseComplement(7))
	assert.Equal(t, 5, bitwiseComplement(10))
}
