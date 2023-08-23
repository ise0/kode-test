package db

import (
	"context"
	"os"

	"github.com/ise0/kode-test/src/shared/logger"
	"github.com/jackc/pgx/v5"
)

var DB *pgx.Conn

func Connect() error {
	d, err := pgx.Connect(context.TODO(), os.Getenv("DB_CONNECTION"))

	if err != nil {
		logger.L.Errorw("Failed to connect to db", "error", err)
		return err
	}

	DB = d
	return nil
}
