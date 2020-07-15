package app

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/JustSteveKing/go-api/pkg/infrastructure/config"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

// Application is our general purpose Application struct
type Application struct {
	Server *http.Server
	Router *mux.Router
	Logger *zap.Logger
	Config *config.Config
}

// Run will run the Application server
func (app *Application) Run() {
	err := app.Server.ListenAndServe()

	if err != nil {
		fmt.Println(err)
	}
}

// WaitForShutdown is a graceful way to handle server shutdown events
func WaitForShutdown(application *Application) {
	// Create a channel to listen for OS signals
	interruptChan := make(chan os.Signal, 1)
	signal.Notify(interruptChan, os.Interrupt, os.Kill, syscall.SIGINT, syscall.SIGTERM)

	// Block until we receive a signal to our channel
	<-interruptChan

	application.Logger.Info("Received shutdown signal, gracefully terminating")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	application.Server.Shutdown(ctx)
	os.Exit(0)
}
