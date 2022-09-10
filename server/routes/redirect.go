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
	x := c.MustGet(gin.AuthUserKey).(string)
	fmt.Println(x)
	c.IndentedJSON(http.StatusOK, "hello")
}
