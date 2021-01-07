package http

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/brauliodev29/go-guest/event"
	"github.com/brauliodev29/go-guest/models"
	"github.com/brauliodev29/go-guest/pkg/entity"
	res "github.com/brauliodev29/go-guest/pkg/response"
)

// Event struct
type Event struct {
	Name     string    `json:"name"`
	Location string    `json:"location"`
	Date     string    `json:"date"`
	Time     time.Time `json:"time"`
}

// Handler func
type Handler struct {
	useCase event.UseCase
}

// NewHandler func
func NewHandler(useCase event.UseCase) *Handler {
	return &Handler{
		useCase: useCase,
	}
}

// Create func
func (h *Handler) Create() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var input Event

		body := json.NewDecoder(r.Body)
		if err := body.Decode(&input); err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}

		params := &models.Event{
			ID:       entity.NewID(),
			Name:     input.Name,
			Date:     input.Date,
			Location: input.Location,
		}

		data, err := h.useCase.CreateEvent(params)
		if err != nil {
			res.BuildError(w, http.StatusBadRequest, err)
			return
		}

		res.BuildJSON(w, http.StatusCreated, data)
	})
}
