package jwt

import (
	"nku-treehole-server/config"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateToken(t *testing.T) {
	config.Setup("../../.env.example")

	jwt := GetJWTCrypto()
	token, err := jwt.GenerateToken("1234")
	assert.Nil(t, err)

	userID, err := jwt.ValidateToken(token)
	assert.Nil(t, err)
	assert.Equal(t, userID, "1234")
}
