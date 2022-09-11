package db

import (
	"context"
	"errors"
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

// Verifys a users password.
// Returns `NoMatch` on wrong password and `ErrNoRows` on user not found
func VerifyUser(state util.State, username string, password string) (*User, error) {
	var user User
	err := state.Db.QueryRow(context.Background(), "SELECT * FROM users where username = $1", username).Scan(&user.Id, &user.Username, &user.Password_Hash, &user.Created_at)
	if err != nil {
		return nil, err
	}
	match, err := argon2id.ComparePasswordAndHash(password, user.Password_Hash)
	if match == true {
		return &user, nil
	} else if (match == false) && (err == nil) {
		return nil, errors.New("NoMatch")
	} else {
		return nil, err
	}
}