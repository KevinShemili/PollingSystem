package unit

import (
	"gin/application/utility"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidatePassword_ValidPassword_ReturnTrue(t *testing.T) {
	password := "Password1"
	assert.True(t, utility.ValidatePassword(password))
}

func TestValidatePassword_ShortPassword_ReturnFalse(t *testing.T) {
	password := "Pass1"
	assert.False(t, utility.ValidatePassword(password))
}

func TestValidatePassword_NoUpperCase_ReturnFalse(t *testing.T) {
	password := "password1"
	assert.False(t, utility.ValidatePassword(password))
}

func TestValidatePassword_NoNrCase_ReturnFalse(t *testing.T) {
	password := "UnifiUnifi"
	assert.False(t, utility.ValidatePassword(password))
}
