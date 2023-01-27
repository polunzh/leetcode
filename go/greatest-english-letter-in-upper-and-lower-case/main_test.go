package greatestenglishletterinupperandlowercase

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	assert.Equal(t, "E", greatestLetter("lEeTcOdE"))
	assert.Equal(t, "R", greatestLetter("arRAzFif"))
	assert.Equal(t, "", greatestLetter("AbCdEfGhIjK"))
}
