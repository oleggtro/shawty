package routes

import (
	"fmt"
	"time"

	"github.com/cloudybyte/shawty/server/db"
	"github.com/cloudybyte/shawty/server/util"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type CreateUserRequest struct {
	Username string `json:"username" binding:required`
	Password string `json:"password" binding:required`
}

type CreateUserResponse struct {
	Token string `json:"token"`
}

func CreateUser(c *gin.Context) {
	x, _ := c.Get("state")
	state := x.(util.State)
	var req CreateUserRequest
	if err := c.BindJSON(&req); err != nil {
		fmt.Printf(`error occurred while creating user: $err`)
		c.AbortWithStatus(500)
		return
	}
	user, err := db.CreateUser(state, req.Username, req.Password)
	if err != nil {
		log.Error("fucked up while creating user: ", err)
		c.AbortWithStatus(500)
		return
	}

	time.Sleep(750 * time.Millisecond)

	session, err := db.CreateSession(state, user.Id)
	if err != nil {
		c.AbortWithStatus(500)
		log.Error("While creating session: ", err)
		return
	}

	// session, err := db.CreateSession(state, user.Id)

	// if err != nil {
	// log.Error("fucked up while creating session: ", err)
	// c.AbortWithStatus(500)
	// return
	// }
	res := CreateUserResponse{session.Token}
	c.JSON(200, res)

}