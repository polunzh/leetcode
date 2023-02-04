package numberofrecentcalls

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test0(t *testing.T) {
	obj := Constructor()

	assert.Equal(t, 1, obj.Ping(1))
	assert.Equal(t, 2, obj.Ping(100))
	assert.Equal(t, 3, obj.Ping(3001))
	assert.Equal(t, 3, obj.Ping(3002))
}

func Test1(t *testing.T) {
	obj := Constructor()

	assert.Equal(t, 1, obj.Ping(642))
	assert.Equal(t, 2, obj.Ping(1849))
	assert.Equal(t, 1, obj.Ping(4921))
	assert.Equal(t, 2, obj.Ping(5936))
	assert.Equal(t, 3, obj.Ping(5957))
}