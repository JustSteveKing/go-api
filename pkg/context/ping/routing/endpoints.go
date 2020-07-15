package routing

import (
	"net/http"

	"github.com/JustSteveKing/go-api/pkg/context/ping/responses"
	"github.com/JustSteveKing/go-api/pkg/infrastructure/app"
	responseFactory "github.com/JustSteveKing/go-http-response"
)

// Handler is the http.Handler for this request
type Handler struct {
	app *app.Application
}

// NewHandler will create a new Handler to handle this request
func NewHandler(app *app.Application) *Handler {
	return &Handler{app}
}

// Handle will handle the incoming request
func (handler *Handler) Handle(response http.ResponseWriter, request *http.Request) {
	handler.app.Logger.Info("Ping Handler Dispatched.")

	responseFactory.Send(
		response,
		http.StatusOK,
		&responses.Response{
			Message: "Service Online",
		},
		handler.app.Config.HTTP.Content,
	)

	return
}
