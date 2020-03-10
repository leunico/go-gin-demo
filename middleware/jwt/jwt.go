package jwt

import (
	"strings"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"

	"git.codepku.com/examinate/exam/pkg/e"
	"git.codepku.com/examinate/exam/pkg/auth"
)

// JWT is jwt middleware
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int

		code = e.SUCCESS
		authString := c.Request.Header.Get("Authorization")
        kv := strings.Split(authString, " ")
        if len(kv) != 2 || kv[0] != "Bearer" {
            code = e.INVALID_AUTH_PARAMS
        } else {
			token := kv[1]
			if token == "" {
				code = e.INVALID_AUTH_PARAMS
			} else {
				claims, err := auth.ParseToken(token)
				if err != nil {
					switch err.(*jwt.ValidationError).Errors {
					case jwt.ValidationErrorExpired:
						code = e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
					default:
						code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
					}
				} else {
					c.Set("authUser", *claims.Examinees)
				}
			}
		}
		
		if code != e.SUCCESS {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message":  e.GetMsg(code),
			})
			c.Abort()
			return
		}

		c.Next()
	}
}