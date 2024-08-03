package social

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"

	"github.com/go-chi/chi/middleware"
	"go.uber.org/zap"
)

func NewHTTPServer(
	logger *zap.Logger,
	s Service,
) http.Handler {

	r := chi.NewRouter()
	r.Use(
		middleware.CleanPath,
		middleware.Timeout(60*time.Second),
		middleware.Recoverer,
	)

	return r
}

func ListenAndServe(ctx context.Context, addr string, handler http.Handler, logger *zap.Logger) {
	l := logger.With(zap.String("logContext", "httpServer"))
	server := &http.Server{
		Addr:    addr,
		Handler: handler,
	}
	l.Info(fmt.Sprintf("listening to %s", addr))

	go func() {
		<-ctx.Done()
		shutdownCtx, _ := context.WithTimeout(context.Background(), 30*time.Second)
		go func() {
			<-shutdownCtx.Done()
			if shutdownCtx.Err() == context.DeadlineExceeded {
				l.Fatal("graceful shutdown timed out, forcing exit")
			}
		}()

		err := server.Shutdown(shutdownCtx)
		if err != nil {
			l.Fatal("failed shutting down server", zap.Error(err))
		}
	}()

	err := server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		l.Fatal("graceful shutdown failed", zap.Error(err))
	}
	l.Info("gracefully shutdown server")
}
