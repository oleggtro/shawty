package db

import (
	"context"
	"time"

	"github.com/cloudybyte/shawty/server/util"
	"github.com/jackc/pgx/v4"
)

type Session struct {
	Token      string
	Subject    string
	Created_at time.Time
	Expires_at time.Time
}

func CreateSession(state util.State, subject string) (*Session, error) {
	var sess Session
	row := state.Db.QueryRow(context.Background(), "INSERT INTO sessions (subject) VALUES ($1) RETURNING *", subject)
	//is there a better way to do this?
	err := scanSession(&sess, row)
	if err != nil {
		return nil, err
	}
	return &sess, err
}

func CheckSession(state util.State, token string) (*Session, error) {
	var sess Session
	row := state.Db.QueryRow(context.Background(), "SELECT * FROM sessions WHERE token = $1", token)
	err := scanSession(&sess, row)
	if err != nil {
		return nil, err
	}
	return &sess, nil
}

func scanSession(sess *Session, row pgx.Row) error {
	return row.Scan(&sess.Token, &sess.Subject, &sess.Created_at, &sess.Expires_at)
}