package noteRouter

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/ise0/kode-test/src/api/middlewares"
	noteController "github.com/ise0/kode-test/src/controllers/notes"
	noteModel "github.com/ise0/kode-test/src/models/note"
	"github.com/ise0/kode-test/src/shared/logger"
)

func New() *chi.Mux {
	m := chi.NewRouter()
	r := m.With(middleware.AllowContentType("application/json"), middlewares.Auth)
	r.Post("/", addNote)
	r.Get("/", getNotes)
	return m
}

type note struct {
	NoteId int    `json:"noteId"`
	Text   string `json:"text"`
}

func addNote(w http.ResponseWriter, r *http.Request) {
	userId, ok := r.Context().Value("userId").(int)
	if !ok {
		logger.L.Errorw("this route must be used within a auth middleware")
		http.Error(w, "something went wrong", 500)
		return
	}
	var body struct {
		Text string `json:"text"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil || strings.TrimSpace(body.Text) == "" {
		http.Error(w, "something went wrong", 500)
		return
	}

	n, err := noteController.AddNew(userId, body.Text)
	if err != nil {
		http.Error(w, "something went wrong", 500)
		return
	}

	res := note{n.Id, n.Text}
	json.NewEncoder(w).Encode(res)
}

func getNotes(w http.ResponseWriter, r *http.Request) {
	userId, ok := r.Context().Value("userId").(int)
	if !ok {
		logger.L.Errorw("this route must be used within a auth middleware")
		http.Error(w, "something went wrong", 500)
		return
	}

	n, err := noteModel.GetList(userId)
	if err != nil {
		http.Error(w, "something went wrong", 500)
		return
	}

	var res []note
	for _, v := range n {
		res = append(res, note{v.Id, v.Text})
	}
	json.NewEncoder(w).Encode(res)
}
