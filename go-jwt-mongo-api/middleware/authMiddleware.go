package middleware

import (
	"fmt"
	"go-jwt-mongo-api/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientToken := c.Request.Header.Get("token")
		if clientToken == "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("No token found")})
			c.Abort()
			return
		}
		claims, err := helper.ValidateToken(clientToken)
		if err != "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Invalid token")})
			c.Abort()
			return
		}
		c.Set("uid", claims.Id)
		c.Set("email", claims.Email)
		c.Set("username", claims.Username)
		c.Set("user_type", claims.User_Type)
		c.Next()
	}
}
