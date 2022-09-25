package sortevenandoddindicesindependently

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	t.Run("[4, 1, 2, 3]", func(t *testing.T) {
		assert.Equal(t, []int{2, 3, 4, 1}, sortEvenOdd([]int{4, 1, 2, 3}))
	})

	t.Run("[2, 1]", func(t *testing.T) {
		assert.Equal(t, []int{2, 1}, sortEvenOdd([]int{2, 1}))
	})
}
