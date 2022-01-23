package indCommonCharacters

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCommonChars(t *testing.T) {
	assert.Equal(t, []string{"e","l","l"}, commonChars([]string{"bella","label","roller"}))
}