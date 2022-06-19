package reversebits

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	assert.Equal(t, uint32(964176192), reverseBits(43261596))
}
