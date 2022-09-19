package db

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/cloudybyte/shawty/server/util"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgtype"
	"github.com/jackc/pgx/v4"
)

type Session struct {
	Token      string
	Subject    string
	Created_at time.Time
	Expires_at time.Time
}

func CreateSession(state util.State, subject pgtype.UUID) (*Session, error) {
	var sess Session
	row := state.Db.QueryRow(context.Background(), "INSERT INTO sessions (subject) VALUES ($1) RETURNING *", subject)
	//is there a better way to do this?
	err := scanSession(&sess, row)
	if err != nil {
		return nil, err
	}
	return &sess, err
}

func CheckAndRenewSession(state util.State, token pgtype.UUID) (*Session, error) {
	var sess Session
	row := state.Db.QueryRow(context.Background(), "UPDATE sessions SET expires_at = current_timestamp + (2 * interval '1 week') WHERE token = $1 AND expires_at > current_timestamp RETURNING *", token)
	err := scanSession(&sess, row)
	if err != nil {
		fmt.Println(`error: `, err)
		fmt.Println("token", token)
		fmt.Println("token string", token)
		return nil, err
	}
	return &sess, nil
}

func RemoveSession(state util.State, token string) error {
	tag, err := state.Db.Exec(context.Background(), "DELETE FROM sessions WHERE token = $1", token)
	if err != nil {
		return err
	} else {
		if tag.RowsAffected() == 0 {
			return errors.New("not found")
			// }
		}
		return nil
	}
}

// ?
func scanSession(sess *Session, row pgx.Row) error {
	return row.Scan(&sess.Token, &sess.Subject, &sess.Created_at, &sess.Expires_at)
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		x, _ := c.Get("state")
		state := x.(util.State)
		y := c.Request.Header.Get("Authorization")
		var token pgtype.UUID
		err := token.Set(y)
		if err != nil {
			c.AbortWithStatus(400)
			return
		}

		session, err := CheckAndRenewSession(state, token)
		if err != nil {
			switch err.Error() {
			case "no rows in result set":
				c.AbortWithStatus(401)
				return
			default:
				fmt.Println("encountered error while validating session: ", err)
				c.AbortWithStatus(500)
				return
			}
		} else {
			c.Set("session", *session)
		}
		c.Next()
	}
}
