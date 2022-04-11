package capitalizethetitle

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	assert.Equal(t, "Capitalize The Title", capitalizeTitle("capiTalIze tHe titLe"))
	assert.Equal(t, "First Letter of Each Word", capitalizeTitle("First leTTeR of EACH Word"))
	assert.Equal(t, "i Love Leetcode", capitalizeTitle("i lOve leetcode"))
	assert.Equal(t, "l hv", capitalizeTitle("L hV"))
	assert.Equal(t, "l Cck n k", capitalizeTitle("l CCK n k"))
}
