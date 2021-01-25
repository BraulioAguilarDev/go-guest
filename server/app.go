package server

import (
	"net/http"
	"time"

	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"

	"github.com/brauliodev29/go-guest/config"
	"github.com/brauliodev29/go-guest/event"
	eventHTTP "github.com/brauliodev29/go-guest/event/delivery/http"
	eventPostgres "github.com/brauliodev29/go-guest/event/repository/postgres"
	eventUseCase "github.com/brauliodev29/go-guest/event/usecase"
	"github.com/brauliodev29/go-guest/pkg/database"
	firebase "github.com/brauliodev29/go-guest/pkg/firebase"
	"github.com/brauliodev29/go-guest/user"
	userHTTP "github.com/brauliodev29/go-guest/user/delivery/http"
	userPostgres "github.com/brauliodev29/go-guest/user/repository/postgres"
	userUseCase "github.com/brauliodev29/go-guest/user/usecase"
)

// App struct
type App struct {
	httpServer *http.Server
	Negroni    *negroni.Negroni
	Router     *mux.Router

	eventUC event.UseCase
	userUC  user.UseCase
}

// NewApp func
func NewApp() (*App, error) {
	// Conexion db
	db, err := database.Init()
	if err != nil {
		return nil, err
	}

	// Firebase config
	authClient, err := firebase.NewFirebase(config.Config.FirebaseJSON).NewClient()
	if err != nil {
		return nil, err
	}

	// Repository instance
	eventRepo := eventPostgres.NewEventRepository(db)
	userRepo := userPostgres.NewUserRepository(db)

	app := &App{
		Router:  mux.NewRouter().StrictSlash(true),
		eventUC: eventUseCase.NewEventUseCase(eventRepo),
		userUC:  userUseCase.NewAuthUseCase(userRepo, authClient),
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
	userHTTP.MakeUserHandlers(a.Router, *a.Negroni, a.userUC)

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
