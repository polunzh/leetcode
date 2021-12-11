package guessNumber

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGuessNumber(t *testing.T) {
	assert.Equal(t, 6, guessNumber(10))
	// assert.Equal(t, 3, guessNumber(4))
	// assert.Equal(t, 5, guessNumber(5))
	// assert.Equal(t, 5, guessNumber(100))
}