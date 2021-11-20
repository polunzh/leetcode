package numberComplement

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTwoSum(t *testing.T) {
	assert.Equal(t, 2, findComplement(5))
	assert.Equal(t, 0, findComplement(1))
}
