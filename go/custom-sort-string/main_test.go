package customsortstring

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	t.Run("test0", func(t *testing.T) {
		assert.Equal(t, "cbad", customSortString("cba", "abcd"))
	})

	t.Run("test1", func(t *testing.T) {
		assert.Equal(t, "cbad", customSortString("cbafg", "abcd"))
	})

	t.Run("test2", func(t *testing.T) {
		assert.Equal(t, "eexvw", customSortString("exv", "xwvee"))
	})
}
