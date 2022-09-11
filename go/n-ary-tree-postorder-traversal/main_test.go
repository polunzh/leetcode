package narytreepostordertraversal

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	root := Node{Val: 1, Children: []*Node{
		{Val: 3, Children: []*Node{{Val: 5}, {Val: 6}}}, {Val: 2}, {Val: 4}},
	}

	assert.Equal(t, []int{5, 6, 3, 2, 4, 1}, postorder2(&root))
}
