package util

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
)

type State struct {
	Db *pgx.Conn
}

// func AuthMiddlewareOld() gin.HandlerFunc {
// return func(c *gin.Context) {
// validateToken(c)
// c.Next()
// }
// }

func Site(state State) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("state", state)
		c.Next()
	}
}