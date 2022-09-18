package db

import (
	"context"
	"math/rand"
	"time"

	"github.com/cloudybyte/shawty/server/util"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
)

type Redirect struct {
	Id         string
	RedirectTo string
	Owner      uuid.UUID
	CreatedAt  time.Time
	Uses       int
}

func CreateRedirect(state util.State, owner uuid.UUID, target string) (*Redirect, error) {
	var res Redirect
	id := genId()
	row := state.Db.QueryRow(context.Background(), "INSERT INTO redirects (id, owner, redirect_to) VALUES ($1, $2, $3) RETURNING *", id, owner, target)
	err := res.Scan(row)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func UseRedirect(state util.State, id string) (*Redirect, error) {
	var res Redirect
	row := state.Db.QueryRow(context.Background(), "UPDATE redirects SET uses = uses + 1 WHERE id = $1 RETURNING *", id)
	err := res.Scan(row)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (red *Redirect) Scan(row pgx.Row) error {
	return row.Scan(&red.Id, &red.RedirectTo, &red.Owner, &red.CreatedAt, &red.Uses)
}

func genId() string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	// the link length
	s := make([]rune, 7)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}