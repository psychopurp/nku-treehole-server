package jwt

import "testing"

func TestGenerateToken(t *testing.T) {
	token, err := GenerateToken("1234")
	t.Log(token, err)
	claims, err := ParseToken(token)
	t.Logf("%#v %v", claims, err)
}
