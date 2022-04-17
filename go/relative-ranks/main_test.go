package relativeranks

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	// assert.Equal(t, []string{"Gold Medal", "Silver Medal", "Bronze Medal", "4", "5"}, findRelativeRanks([]int{5, 4, 3, 2, 1}))
	assert.Equal(t, []string{"Gold Medal", "5", "Bronze Medal", "Silver Medal", "4"}, findRelativeRanks([]int{10, 3, 8, 9, 4}))
	assert.Equal(t, []string{"Gold Medal"}, findRelativeRanks([]int{10}))
}
