package ping

import (
	"net/http"

	"github.com/JustSteveKing/go-api/pkg/context/ping/routing"
	"github.com/JustSteveKing/go-api/pkg/infrastructure/app"
)

// BuildPingService is responsible for setting up the Ping Context for our Domain
func BuildPingService(app *app.Application) {
	// Create our Handler
	handler := routing.NewHandler(app)

	// Create a sub router for this service
	router := app.Router.Methods(http.MethodGet).Subrouter()

	// Register our service routes
	router.HandleFunc("/ping", handler.Handle).Name("ping:handle")
}
