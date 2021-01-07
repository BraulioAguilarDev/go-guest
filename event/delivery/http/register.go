package http

import (
	"net/http"

	"github.com/brauliodev29/go-guest/event"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

// MakeEventHandlers func
func MakeEventHandlers(
	route *mux.Router,
	n negroni.Negroni,
	uc event.UseCase,
) {
	h := NewHandler(uc)

	api := route.PathPrefix("/api").Subrouter()

	api.Handle(
		"/events",
		n.With(negroni.Wrap(h.Create())),
	).Methods(http.MethodPost, http.MethodOptions).Name("events")
}
