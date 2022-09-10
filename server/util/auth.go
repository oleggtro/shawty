package util

import "github.com/gin-gonic/gin"

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func validateToken(c *gin.Context) {
	token := c.Request.Header.Get("Authentication")

	if token == "" {
		c.AbortWithStatus(401)
	} else {
		c.AbortWithStatus(401)
	}
}