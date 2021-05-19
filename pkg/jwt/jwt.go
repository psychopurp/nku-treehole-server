package jwt

import (
	jwt2 "github.com/dgrijalva/jwt-go"
)

var jwt2Secret = []byte("nku-treehole")

type Claims struct {
	UserID string `json:"user_id"`
	jwt2.StandardClaims
}

// 产生token的函数
func GenerateToken(userID string) (string, error) {
	claims := Claims{
		userID,
		jwt2.StandardClaims{Issuer: "nku-treehole"},
	}
	tokenClaims := jwt2.NewWithClaims(jwt2.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwt2Secret)
	return token, err
}

// 验证token的函数
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt2.ParseWithClaims(token, &Claims{}, func(token *jwt2.Token) (interface{}, error) {
		return jwt2Secret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
