package util

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
)

type State struct {
	Db *pgx.Conn
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		validateToken(c)
		c.Next()
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

func Site(state State) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("state", state)
		c.Next()
	}
}