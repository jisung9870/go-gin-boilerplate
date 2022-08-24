package middlewares

import (
	"net/http"

	"github.com/JisungPark0319/go-gin-boilerplate/auth"
	"github.com/gin-gonic/gin"
)

func Authorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		accessToken := c.GetHeader("Authorization")
		ok, err := auth.VerifyAccessToken(accessToken)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": err.Error()})
			return
		}
		if !ok {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Unauthorized token"})
			return
		}
		c.Next()
	}
}
