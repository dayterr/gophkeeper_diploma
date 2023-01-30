package authjwt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateToken(t *testing.T) {
	testLogin := "user45"
	testKey := "dnfjdf5t84"

	_, err := CreateToken(testLogin, testKey)
	assert.NoError(t, err)
}

func TestEncryptPassword(t *testing.T) {
	testPass := "some_secret4567,djd"
	testKey := "dnfjdf5t84"

	enc := EncryptPassword(testPass, testKey)
	assert.NotEmpty(t, enc)
}