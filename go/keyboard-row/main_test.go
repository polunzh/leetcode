package KeyboardRow

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindWords(t *testing.T) {
	assert.Equal(t, []string{"Alaska", "Dad"}, findWords([]string{"Hello", "Alaska", "Dad", "Peace"}))
	assert.Equal(t, []string{}, findWords([]string{"omk"}))
	assert.Equal(t, []string{"adsdf", "sfd"}, findWords([]string{"adsdf", "sfd"}))
}
