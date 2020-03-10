package auth

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"

	"git.codepku.com/examinate/exam/models"
)

var jwtSecret []byte

type Claims struct {
	jwt.StandardClaims
	*models.Examinees
}

// GenerateToken generate tokens used for auth
func GenerateToken(user *models.Examinees) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(24 * time.Hour)

	stdClaims := jwt.StandardClaims{
		ExpiresAt: expireTime.Unix(),
		IssuedAt: time.Now().Unix(),
		Issuer: "exam",
		Id: fmt.Sprintf("%d", user.ID),
	}

	claims := Claims{
		StandardClaims: stdClaims,
		Examinees: user,
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