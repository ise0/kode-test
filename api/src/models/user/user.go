package userModel

import (
	"context"
	"errors"

	"github.com/ise0/kode-test/src/shared/db"
	"github.com/ise0/kode-test/src/shared/logger"
	"github.com/ise0/kode-test/src/shared/passhash"
)

type user struct {
	Id       int    `json:"user_id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func New(username string, password string) (user, error) {
	if len(password) < 1 {
		return user{}, errors.New("weak password")
	}
	if len(username) < 1 {
		return user{}, errors.New("short username")
	}

	var u user
	passwordHash, err := passhash.Hash(password)
	if err != nil {
		return user{}, err
	}

	if err := db.DB.QueryRow(context.TODO(), `
		insert into users(username, user_password) values ($1, $2)
		returning user_id, username, user_password
	`, username, passwordHash).Scan(&u.Id, &u.Username, &u.Password); err != nil {
		logger.L.Errorw("Query failed", "error", err)
		return user{}, err
	}

	return u, nil
}

func GetByUsername(username string) (user, error) {
	var u user

	if err := db.DB.QueryRow(context.TODO(), `
		select user_id, username, user_password 
		from users 
		where username = $1
	`, username).Scan(&u.Id, &u.Username, &u.Password); err != nil {
		logger.L.Errorw("Query failed", "error", err)
		return user{}, err
	}

	return u, nil
}
