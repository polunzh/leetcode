package checkifnumbersareascendinginasentence

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test1(t *testing.T) {
	assert.Equal(t, true, areNumbersAscending("1 box has 3 blue 4 red 6 green and 12 yellow marbles"))
}

func Test2(t *testing.T) {
	assert.Equal(t, false, areNumbersAscending("hello world 5 x 5"))
}

func Test3(t *testing.T) {
	assert.Equal(t, false, areNumbersAscending("sunset is at 7 51 pm overnight lows will be in the low 50 and 60 s"))
}
