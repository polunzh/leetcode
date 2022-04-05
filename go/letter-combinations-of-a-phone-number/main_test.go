package lettercombinationsofaphonenumber

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLetterCombinations(t *testing.T) {
	assert.ElementsMatch(t, []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}, letterCombinations("23"))
	assert.ElementsMatch(t, []string{}, letterCombinations(""))
}
