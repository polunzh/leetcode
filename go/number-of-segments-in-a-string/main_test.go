package numberofsegmentsinastring

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test1(t *testing.T) {
	assert.Equal(t, 5, countSegments("Hello, my name is John"))
}

func Test2(t *testing.T) {
	assert.Equal(t, 1, countSegments("Hello"))
}

func Test3(t *testing.T) {
	assert.Equal(t, 1, countSegments(" Hello"))
}

func Test4(t *testing.T) {
	assert.Equal(t, 1, countSegments(" Hello "))
}