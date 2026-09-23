package middleware

import (
	"strings"

	"seckill/internal/pkg/jwt"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.Request.Header.Get("Authorization")
		token := strings.Trim(strings.TrimSpace(tokenString), "Beearer")
		claims, err := jwt.ParseToken(token)
		if err != nil {
			c.Abort()
		}
		c.Set("username", claims.Username)
		c.Set("user_id", claims.UserID)
		c.Next()
	}
}
