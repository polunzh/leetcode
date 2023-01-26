package strongpasswordcheckerii

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStrongPasswordCheckerII(t *testing.T) {
	assert.Equal(t, false, strongPasswordCheckerII(""))
	assert.Equal(t, true, strongPasswordCheckerII("IloveLe3tcode!"))
	assert.Equal(t, false, strongPasswordCheckerII("Me+You--IsMyDream"))
	assert.Equal(t, false, strongPasswordCheckerII("1aB!"))
	assert.Equal(t, false, strongPasswordCheckerII("A8(bblafkjiqow"))
	assert.Equal(t, false, strongPasswordCheckerII("ecuwcfoyajkolntovfniplayrxhzpmhrkhzonopcwxgupzhoupw"))
	assert.Equal(t, true, strongPasswordCheckerII("a1A!A!A!"))
}
