package db

import (
	"context"
	"github.com/alexedwards/argon2id"
	"github.com/cloudybyte/shawty/server/util"
	"time"
)

type User struct {
	Id            string
	Username      string
	Password_Hash string
	Created_at    time.Time
}

func CreateUser(state util.State, username string, password string) (result *User, err error) {
	hash, err := argon2id.CreateHash(password, argon2id.DefaultParams)
	if err != nil {
		return nil, err
	}
	user := User{}
	res := state.Db.QueryRow(context.Background(), "INSERT INTO users (username, password) VALUES ($1, $2)", username, hash)
	res.Scan(user.Id, user.Username, user.Password_Hash, user.Created_at)

	return &user, nil
}