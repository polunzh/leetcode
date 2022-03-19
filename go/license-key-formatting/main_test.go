package licenseKeyFormatting

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLicenseKeyFormatting(t *testing.T) {
	assert.Equal(t, "5F3Z-2E9W", licenseKeyFormatting("5F3Z-2e-9-w", 4))
	assert.Equal(t, "2-5G-3J", licenseKeyFormatting("2-5g-3-J", 2))
}
