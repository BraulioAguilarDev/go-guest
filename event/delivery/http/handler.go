package http

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/brauliodev29/go-guest/event"
)

// Event struct
type Event struct {
	Name     string    `json:"name"`
	Location string    `json:"location"`
	Date     string    `json:"date"`
	Time     time.Time `json:"time"`
}

type Handler struct {
	useCase event.UseCase
}

func NewHandler(useCase event.UseCase) *Handler {
	return &Handler{
		useCase: useCase,
	}
}

func (h *Handler) Create() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var input Event

		body := json.NewDecoder(r.Body)
		if err := body.Decode(&input); err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}

		w.Write([]byte("hola"))
	})
}
