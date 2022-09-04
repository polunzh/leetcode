package uniquenumberofoccurrences

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test1(t *testing.T) {
	res := uniqueOccurrences([]int{1,2})
	assert.Equal(t, false, res)
}

func Test2(t *testing.T) {
	res := uniqueOccurrences([]int{1,2,2,1,1,3})
	assert.Equal(t, true, res)
}