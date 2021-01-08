package http

import (
	"encoding/json"
	"net/http"

	"github.com/brauliodev29/go-guest/event"
	"github.com/brauliodev29/go-guest/pkg/entity"
	res "github.com/brauliodev29/go-guest/pkg/response"
	"github.com/gorilla/mux"
)

// Event struct
type Event struct {
	Name     string `json:"name"`
	Location string `json:"location"`
	Date     string `json:"date"`
	Time     string `json:"time"`
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

		p := &Event{
			Name:     input.Name,
			Date:     input.Date,
			Location: input.Location,
			Time:     input.Time,
		}

		data, err := h.useCase.CreateEvent(p.Name, p.Location, p.Date, p.Time)
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

		var input Event

		body := json.NewDecoder(r.Body)
		if err := body.Decode(&input); err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}

		p := &Event{
			Name:     input.Name,
			Date:     input.Date,
			Location: input.Location,
			Time:     input.Time,
		}

		id, _ := entity.StringToID(query["uid"])

		if err := h.useCase.UpdateEvent(p.Name, p.Location, p.Date, p.Time, id); err != nil {
			res.BuildError(w, http.StatusBadRequest, err)
			return
		}

		res.BuildJSON(w, http.StatusOK, "Updated")
	})
}
