package sortarraybyparity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	t.Run("[3, 1, 2, 4]", func(t *testing.T) {
		assert.Equal(t, []int{2, 4, 3, 1}, sortArrayByParity([]int{3, 1, 2, 4}))
	})

	t.Run("[0]", func(t *testing.T) {
		assert.Equal(t, []int{0}, sortArrayByParity([]int{0}))
	})

	t.Run("[0,1,2]", func(t *testing.T) {
		assert.Equal(t, []int{0, 2, 1}, sortArrayByParity([]int{0, 1, 2}))
	})
}
