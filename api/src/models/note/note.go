package noteModel

import (
	"context"

	"github.com/ise0/kode-test/src/shared/db"
	"github.com/ise0/kode-test/src/shared/logger"
)

type Note struct {
	Id     int    `json:"note_id"`
	UserId string `json:"user_id"`
	Text   string `json:"body"`
}

func New(userId int, text string) (Note, error) {
	var n Note

	if err := db.DB.QueryRow(context.TODO(), `
		insert into notes(user_id, body) values ($1, $2)
		returning note_id, user_id, body
	`, userId, text).Scan(&n.Id, &n.UserId, &n.Text); err != nil {
		logger.L.Errorw("Query failed", "error", err)
		return Note{}, err
	}

	return n, nil
}

func GetList(userId int) ([]Note, error) {
	var res []Note

	rows, err := db.DB.Query(context.TODO(), `
		select note_id, user_id, body
		from notes
		where user_id = $1
	`, userId)

	if err != nil {
		logger.L.Errorw("Query failed", "error", err)
		return res, err
	}

	for rows.Next() {
		var n Note
		err = rows.Scan(&n.Id, &n.UserId, &n.Text)
		if err != nil {
			logger.L.Errorw("Query failed", "error", err)
			return res, err
		}
		res = append(res, n)
	}

	return res, nil
}
