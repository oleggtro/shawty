package routes

import (
	"fmt"

	"github.com/cloudybyte/shawty/server/db"
	"github.com/cloudybyte/shawty/server/util"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type CreateSessionRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type CreateSessionResponse struct {
	Token string `json:"token"`
}

func CreateSession(c *gin.Context) {
	x, _ := c.Get("state")
	state := x.(util.State)

	req := CreateSessionRequest{}
	c.BindJSON(&req)
	user, err := db.VerifyUser(state, req.Username, req.Password)
	if err != nil {
		switch err.Error() {
		case "NoMatch":
			c.AbortWithStatus(401)
		case "no rows in result set":
			// == user not found
			c.AbortWithStatus(401)
		default:
			fmt.Println("probably encountered db error: ", err)
			c.AbortWithStatus(500)
		}
		return
	}

	sess, err := db.CreateSession(state, user.Id)
	if err != nil {
		c.AbortWithStatus(500)
		log.Error("While creating session: ", err)
		return
	}

	res := CreateSessionResponse{sess.Token}
	c.JSON(200, res)
}
