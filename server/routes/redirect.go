package routes

import (
	"fmt"
	"net/http"

	"github.com/cloudybyte/shawty/server/db"
	"github.com/cloudybyte/shawty/server/util"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CreateShortlinkRequest struct {
	Url string `json:url`
}

type CreateShortlinkResponse struct {
	Short string `json:shortlink`
}

func CreateShortlink(c *gin.Context) {
	x, _ := c.Get("state")
	state := x.(util.State)

	y, e := c.Get("session")
	// session does not exist => unauthorized
	if e == false {
		c.AbortWithStatus(401)
		return
	}
	session := y.(db.Session)

	var req CreateShortlinkRequest
	if err := c.BindJSON(&req); err != nil {
		fmt.Printf(`error occurred while creating shortlink: $err`)
		c.AbortWithStatus(400)
		return
	}

	// `session.Subject` string can never be a non-uuid
	short, err := db.CreateRedirect(state, uuid.MustParse(session.Subject), req.Url)
	if err != nil {
		c.AbortWithStatus(500)
		return
	}

	response := CreateShortlinkResponse{
		Short: short.Id,
	}
	c.IndentedJSON(http.StatusOK, response)
}
func UseShortlink(c *gin.Context) {
	id := c.Param("id")
	x, _ := c.Get("state")
	state := x.(util.State)
	res, err := db.UseRedirect(state, id)
	if err != nil {
		fmt.Println(`encountered err while fetching redirect: $err`)
		c.AbortWithStatus(500)
		return
	} else {
		c.Redirect(301, res.RedirectTo)
		return
	}
}