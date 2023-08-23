package api

import (
	"github.com/go-chi/chi/v5"
	"github.com/ise0/kode-test/src/api/middlewares"
	noteRouter "github.com/ise0/kode-test/src/api/routes/notes"
	userRouter "github.com/ise0/kode-test/src/api/routes/user"
)

func New() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middlewares.Logger)
	r.Mount("/user", userRouter.New())
	r.Mount("/notes", noteRouter.New())
	return r
}
