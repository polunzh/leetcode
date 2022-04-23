package reshapethematrix

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	assert.Equal(t, [][]int{{1, 2, 3, 4}}, matrixReshape([][]int{{1, 2}, {3, 4}}, 1, 4))
	assert.Equal(t, [][]int{{1, 2}, {3, 4}}, matrixReshape([][]int{{1, 2}, {3, 4}}, 2, 4))
	assert.Equal(t, [][]int{{1, 2}, {3, 4}}, matrixReshape([][]int{{1, 2, 3, 4}}, 2, 2))
	assert.Equal(t, [][]int{{1}, {2}, {3}, {4}}, matrixReshape([][]int{{1, 2, 3, 4}}, 4, 1))
}
