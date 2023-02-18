package replacethesubstringforbalancedstring

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	assert.Equal(t, int64(0), findTheArrayConcVal([]int{}))
	assert.Equal(t, int64(0), findTheArrayConcVal([]int{0}))
	assert.Equal(t, int64(596), findTheArrayConcVal([]int{7, 52, 2, 4}))
	assert.Equal(t, int64(673), findTheArrayConcVal([]int{5, 14, 13, 8, 12}))
}
