package findthehighestaltitude

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	t.Run("test0", func(t *testing.T) {
		assert.Equal(t, 0, largestAltitude([]int{}))
	})

	t.Run("test1", func(t *testing.T) {
		assert.Equal(t, 1, largestAltitude([]int{-5, 1, 5, 0, -7}))
	})

	t.Run("test2", func(t *testing.T) {
		assert.Equal(t, 0, largestAltitude([]int{-4, -3, -2, -1, 4, 3, 2}))
	})
}
