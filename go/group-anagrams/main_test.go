package groupanagrams

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	assert.Equal(t, [][]string{{""}}, groupAnagrams([]string{""}))
}
