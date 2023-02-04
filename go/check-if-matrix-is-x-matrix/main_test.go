package checkifmatrixisxmatrix

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	grid := [][]int{{2, 0, 0, 1}, {0, 3, 1, 0}, {0, 5, 2, 0}, {4, 0, 0, 2}}

	assert.Equal(t, true, checkXMatrix(grid))
}
