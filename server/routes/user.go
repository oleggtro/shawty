package routes

import (
	"fmt"

	"github.com/cloudybyte/shawty/server/db"
	"github.com/cloudybyte/shawty/server/util"
	"github.com/gin-gonic/gin"
)

type CreateUserRequest struct {
	Username string `json:username binding:required`
	Password string `json:password binding:required`
}

type CreateUserResponse struct {
	Token string `json:token`
}

func CreateUser(c *gin.Context) {
	state, _ := c.Get("state")
	var req CreateUserRequest
	if err := c.BindJSON(&req); err != nil {
		fmt.Printf(`error occurred while creating user: $err`)
		c.AbortWithStatus(500)
		return
	}
	user, _ := db.CreateUser(state.(util.State), req.Username, req.Password)
	fmt.Println(user)

}