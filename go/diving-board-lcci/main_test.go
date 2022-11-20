package divingboardlcci

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	t.Run("test0", func(t *testing.T) {
		assert.Equal(t, []int{}, divingBoard(1, 1, 0))
	})

	t.Run("test1", func(t *testing.T) {
		assert.Equal(t, []int{3, 4, 5, 6}, divingBoard(1, 2, 3))
	})
}
