package http

import (
	"net/http"

	"github.com/brauliodev29/go-guest/user"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

// MakeUserHandlers func
func MakeUserHandlers(
	route *mux.Router,
	n negroni.Negroni,
	uc user.UseCase,
) {

	h := NewHandler(uc)

	route.Handle(
		"/sing-up",
		n.With(negroni.Wrap(h.Create())),
	).Methods(http.MethodPost, http.MethodOptions).Name("singUp")
}
