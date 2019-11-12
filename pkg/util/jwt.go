package util

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtSecret []byte

type Claims struct {
	Certificates string `json:"certificates"`
	Password string `json:"password"`
	jwt.StandardClaims
}

// GenerateToken generate tokens used for auth
func GenerateToken(certificates, password string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(24 * time.Hour)

	claims := Claims{
		certificates,
		password,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer: "exam",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

// ParseToken parsing token
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}