package kernel

import (
	"context"
	"net/http"
	"time"

	"github.com/JustSteveKing/go-api/pkg/infrastructure/app"
	"github.com/JustSteveKing/go-api/pkg/infrastructure/config"
	"github.com/google/uuid"
	gohandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/micro/go-micro/v2/util/log"
	"go.uber.org/zap"
)

// Boot the Application
func Boot() *app.Application {
	// Configuration
	config := bootConfig()

	// Router
	router := bootRouter()

	// CORS
	corsHandler := gohandlers.CORS(gohandlers.AllowedOrigins([]string{"*"}))

	// Logger
	logger := bootLogger()
	defer logger.Sync() // flushes buffer, if any

	// Create and return and Application
	return &app.Application{
		Server: &http.Server{
			Addr:         ":8080",
			Handler:      corsHandler(requestIDMiddleware(router)),
			IdleTimeout:  120 * time.Second,
			ReadTimeout:  1 * time.Second,
			WriteTimeout: 1 * time.Second,
		},
		Router: router,
		Logger: logger,
		Config: config,
	}
}

func bootConfig() *config.Config {
	return config.New()
}

func bootRouter() *mux.Router {
	return mux.NewRouter()
}

func bootLogger() *zap.Logger {
	logger, logError := zap.NewProduction()
	if logError != nil {
		panic(logError)
	}

	return logger
}

// ContextKey is used for context.Context value. The value requires a key that is not primitive type.
type ContextKey string

// ContextKeyRequestID is the ContextKey for RequestID
const ContextKeyRequestID ContextKey = "requestID"

func requestIDMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		ctx := r.Context()

		id := uuid.New()

		ctx = context.WithValue(ctx, ContextKeyRequestID, id.String())

		r = r.WithContext(ctx)

		log.Debugf("Incomming request %s %s %s %s", r.Method, r.RequestURI, r.RemoteAddr, id.String())

		next.ServeHTTP(w, r)

		log.Debugf("Finished handling http req. %s", id.String())
	})
}
