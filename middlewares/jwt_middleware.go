package middlewares

import (
	"log"
	"net/http"
	"strings"

	"web-project/utils"

	"github.com/gin-gonic/gin"
)

func JwtAuthorizationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		tokenString := strings.TrimSpace(strings.SplitN(authHeader, "Bearer", 2)[1])
		claims, err := utils.ValidateAndGetClaims(tokenString)
		if err == nil {
			log.Println(claims["exp"])
			c.Set("user_name", claims["user_name"])
			c.Next()
		} else {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}
