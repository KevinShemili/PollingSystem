package unit

import (
	"gin/application/utility"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateEmail_ValidEmail_ReturnTrue(t *testing.T) {
	email := "test@example.com"
	assert.True(t, utility.ValidateEmail(email))
}

func TestValidateEmail_InvalidEmail_ReturnFalse(t *testing.T) {
	email := "invalid-email"
	assert.False(t, utility.ValidateEmail(email))
}
