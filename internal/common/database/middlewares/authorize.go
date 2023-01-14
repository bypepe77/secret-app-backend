package middlewares

import (
	"fmt"
	"net/http"

	"github.com/bypepe77/secret-app-backend/internal/common"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Authorize() gin.HandlerFunc {
	return func(c *gin.Context) {
		const BEARER_SCHEMA = "Bearer "
		authHeader := c.GetHeader("Authorization")
		tokenString := authHeader[len(BEARER_SCHEMA):]
		token, err := common.JWTAuthService().ValidateToken(tokenString)
		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			mcClaims, err := common.FromMapClaims(claims)
			if err != nil || mcClaims == nil {
				fmt.Println("err", err)
				c.AbortWithStatus(http.StatusUnauthorized)
			}
			c.Set("claims", mcClaims)
		} else {
			fmt.Println(err)
			c.AbortWithStatus(http.StatusUnauthorized)
		}

	}
}
