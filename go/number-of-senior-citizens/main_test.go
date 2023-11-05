package numberofseniorcitizens

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test1(t *testing.T) {
	assert.Equal(t, 2, countSeniors([]string{"7868190130M7522", "5303914400F9211", "9273338290F4010"}))
}
