package server

import (
	"net/http"
	"time"

	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"

	"github.com/brauliodev29/go-guest/event"
	eventHTTP "github.com/brauliodev29/go-guest/event/delivery/http"
	eventPostgres "github.com/brauliodev29/go-guest/event/repository/postgres"
	eventUseCase "github.com/brauliodev29/go-guest/event/usecase"
	"github.com/brauliodev29/go-guest/pkg/database"
)

// App struct
type App struct {
	httpServer *http.Server
	Negroni    *negroni.Negroni
	Router     *mux.Router

	eventUC event.UseCase
}

// NewApp func
func NewApp() (*App, error) {
	db, err := database.Init()
	if err != nil {
		return nil, err
	}

	// Repository instance
	eventRepo := eventPostgres.NewEventRepository(db)

	app := &App{
		Router:  mux.NewRouter().StrictSlash(true),
		eventUC: eventUseCase.NewEventUseCase(eventRepo),
		Negroni: negroni.New(
			negroni.NewLogger(),
			negroni.NewRecovery(),
		),
	}

	return app, nil
}

// Run func
func (a *App) Run(port string) error {

	// Make handlers
	eventHTTP.MakeEventHandlers(a.Router, *a.Negroni, a.eventUC)

	http.Handle("/", a.Router)

	a.httpServer = &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		Addr:         ":" + port,
		Handler:      context.ClearHandler(http.DefaultServeMux),
	}

	if err := a.httpServer.ListenAndServe(); err != nil {
		return err
	}

	return nil
}
