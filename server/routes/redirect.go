package routes

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateShortlinkRequest struct {
	Url string `json:url`
}

type CreateShortlinkResponse struct {
}

func CreateShortlink(c *gin.Context) {
	x, exists := c.Get(gin.AuthUserKey)
	fmt.Println(x, exists)
	c.IndentedJSON(http.StatusOK, "hello")
}
