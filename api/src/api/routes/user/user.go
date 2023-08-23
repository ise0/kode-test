package userRouter

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	userModel "github.com/ise0/kode-test/src/models/user"
	"github.com/ise0/kode-test/src/shared/authjwt"
	"github.com/ise0/kode-test/src/shared/passhash"
)

func New() *chi.Mux {
	m := chi.NewRouter()
	r := m.With(middleware.AllowContentType("application/json"))
	r.Post("/sign-in", SignIn)
	r.Post("/sign-up", signUp)
	return m
}

type user struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type response struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Jwt      string `json:"jwt"`
}

func SignIn(w http.ResponseWriter, r *http.Request) {
	var u user

	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		http.Error(w, "something went wrong", 500)
		return
	}

	user, err := userModel.GetByUsername(u.Username)
	if err != nil {
		http.Error(w, "something went wrong", 500)
		return
	}

	if !passhash.Verify(u.Password, user.Password) {
		http.Error(w, "something went wrong", 500)
		return
	}

	token, err := authjwt.SignAuthJwt(user.Id)
	if err != nil {
		http.Error(w, "something went wrong", 500)
		return
	}

	res := response{user.Id, user.Username, token}
	json.NewEncoder(w).Encode(res)
}

func signUp(w http.ResponseWriter, r *http.Request) {
	var u user

	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		http.Error(w, "something went wrong", 500)
		return
	}

	user, err := userModel.New(u.Username, u.Password)
	if err != nil {
		http.Error(w, "something went wrong", 500)
		return
	}

	token, err := authjwt.SignAuthJwt(user.Id)
	if err != nil {
		http.Error(w, "something went wrong", 500)
		return
	}

	res := response{user.Id, user.Username, token}

	json.NewEncoder(w).Encode(res)
}
