package main

import (
	"github.com/JustSteveKing/go-api/pkg/application/registry/ping"
	application "github.com/JustSteveKing/go-api/pkg/infrastructure/app"
	"github.com/JustSteveKing/go-api/pkg/infrastructure/kernel"
	"github.com/joho/godotenv"
)

// init is invoked before main()
func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		panic("No .env file found")
	}
}

func main() {
	// Create our application
	app := kernel.Boot()

	// Build our services
	ping.BuildPingService(app)

	// Run our Application in a coroutine
	go func() {
		app.Run()
	}()

	// Wait for termination signals and shut down gracefully
	application.WaitForShutdown(app)
}
