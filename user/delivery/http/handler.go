package http

import (
	"encoding/json"
	"net/http"

	"github.com/brauliodev29/go-guest/pkg/entity"
	res "github.com/brauliodev29/go-guest/pkg/response"
	"github.com/brauliodev29/go-guest/user"
	"github.com/gorilla/mux"
)

// User struct
type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

// Handler struct
type Handler struct {
	useCase user.UseCase
}

// NewHandler func
func NewHandler(useCase user.UseCase) *Handler {
	return &Handler{
		useCase: useCase,
	}
}

// Create func
func (h *Handler) Create() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var input User

		body := json.NewDecoder(r.Body)
		if err := body.Decode(&input); err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}

		p := &User{
			FirstName: input.FirstName,
			LastName:  input.LastName,
			Email:     input.Email,
		}

		data, err := h.useCase.CreateUser(
			p.FirstName,
			p.LastName,
			p.Email,
		)

		if err != nil {
			res.BuildError(w, http.StatusBadRequest, err)
			return
		}

		res.BuildJSON(w, http.StatusCreated, data)
	})
}

// Update func
func (h *Handler) Update() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		query := mux.Vars(r)

		var input User

		body := json.NewDecoder(r.Body)
		if err := body.Decode(&input); err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}

		p := &User{
			FirstName: input.FirstName,
			LastName:  input.LastName,
		}

		id, _ := entity.StringToID(query["uid"])

		if err := h.useCase.UpdateUser(p.FirstName, p.LastName, id); err != nil {
			res.BuildError(w, http.StatusBadRequest, err)
			return
		}

		res.BuildJSON(w, http.StatusOK, "Updated")
	})
}
