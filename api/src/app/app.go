package app

import (
	"context"
	"net/http"
	"os"
	"time"

	"github.com/ise0/kode-test/src/api"
	"github.com/ise0/kode-test/src/shared/db"
	"github.com/ise0/kode-test/src/shared/logger"
	"github.com/ise0/kode-test/src/shared/retry"
)

func Start() {
	retry.Exec(context.TODO(), func() error {
		return db.Connect()
	}, retry.Options{Retries: -1, Delay: time.Second * 2})

	logger.L.Info("app started on port " + os.Getenv("PORT"))

	http.ListenAndServe(":"+os.Getenv("PORT"), api.New())
}
