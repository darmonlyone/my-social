package social

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	kithttp "github.com/go-kit/kit/transport/http"
	"go.uber.org/zap"
)

func NewHTTPServer(
	logger *zap.Logger,
	svc Service,
) http.Handler {

	router := chi.NewRouter()
	router.Use(
		middleware.CleanPath,
		middleware.Timeout(60*time.Second),
		middleware.Recoverer,
	)

	opts := []kithttp.ServerOption{
		kithttp.ServerErrorEncoder(EncodeError),
	}

	endpoints := MakeSocialEndpoints(svc)

	loginHandler := MakeLoginHTTPHandler(opts, endpoints, logger, svc)
	registerHandler := MakeRegisterHTTPHandler(opts, endpoints, logger, svc)

	createPostHandler := MakeCreatePostHTTPHandler(opts, endpoints, logger, svc)
	listPostHandler := MakeListPostHTTPHandler(opts, endpoints, logger, svc)
	getPostHandler := MakeGetPostHTTPHandler(opts, endpoints, logger, svc)
	editPostHandler := MakeEditPostHTTPHandler(opts, endpoints, logger, svc)
	deletePostHandler := MakeDeletePostHTTPHandler(opts, endpoints, logger, svc)

	router.Group(func(r chi.Router) {
		r.Post("/login", loginHandler.ServeHTTP)
		r.Post("/register", registerHandler.ServeHTTP)

		// Protected routes
		r.With(BasicAuthMiddleware(svc)).Group(func(r chi.Router) {
			r.Post("/post.create", createPostHandler.ServeHTTP)
			r.Post("/post.list", listPostHandler.ServeHTTP)
			r.Post("/post.get", getPostHandler.ServeHTTP)
			r.Post("/post.edit", editPostHandler.ServeHTTP)
			r.Post("/post.delete", deletePostHandler.ServeHTTP)
		})

	})
	return router
}

func MakeLoginHTTPHandler(opts []kithttp.ServerOption, endpoints SocialEndpoints, _ *zap.Logger, _ Service) http.Handler {
	return kithttp.NewServer(
		endpoints.Login,
		DecodeMappingBodyRequest[LoginRequest],
		MakeEncodeResponse(EncodeError, 200),
		opts...,
	)
}

func MakeRegisterHTTPHandler(opts []kithttp.ServerOption, endpoints SocialEndpoints, _ *zap.Logger, _ Service) http.Handler {
	return kithttp.NewServer(
		endpoints.Register,
		DecodeMappingBodyRequest[RegisterRequest],
		MakeEncodeResponse(EncodeError, 201),
		opts...,
	)
}

func MakeCreatePostHTTPHandler(opts []kithttp.ServerOption, endpoints SocialEndpoints, _ *zap.Logger, _ Service) http.Handler {
	return kithttp.NewServer(
		endpoints.StorePost,
		DecodeMappingBodyRequest[StorePostRequest],
		MakeEncodeResponse(EncodeError, 201),
		opts...,
	)
}

func MakeListPostHTTPHandler(opts []kithttp.ServerOption, endpoints SocialEndpoints, _ *zap.Logger, _ Service) http.Handler {
	return kithttp.NewServer(
		endpoints.FindAllPosts,
		decodeEmptyRequest,
		MakeEncodeResponse(EncodeError, 200),
		opts...,
	)
}

func MakeGetPostHTTPHandler(opts []kithttp.ServerOption, endpoints SocialEndpoints, _ *zap.Logger, _ Service) http.Handler {
	return kithttp.NewServer(
		endpoints.FindPost,
		DecodeMappingBodyRequest[FindPostRequest],
		MakeEncodeResponse(EncodeError, 200),
		opts...,
	)
}

func MakeEditPostHTTPHandler(opts []kithttp.ServerOption, endpoints SocialEndpoints, _ *zap.Logger, _ Service) http.Handler {
	return kithttp.NewServer(
		endpoints.UpdatePost,
		DecodeMappingBodyRequest[UpdatePostRequest],
		MakeEncodeResponse(EncodeError, 200),
		opts...,
	)
}

func MakeDeletePostHTTPHandler(opts []kithttp.ServerOption, endpoints SocialEndpoints, _ *zap.Logger, _ Service) http.Handler {
	return kithttp.NewServer(
		endpoints.DeletePost,
		DecodeMappingBodyRequest[DeletePostRequest],
		MakeEncodeResponse(EncodeError, 200),
		opts...,
	)
}

func decodeEmptyRequest(_ context.Context, _ *http.Request) (interface{}, error) {
	return nil, nil
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

// encode errors from business-logic
func EncodeError(_ context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	HandleCommonErrors(err, w)

	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}

func BasicAuthMiddleware(svc Service) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			auth := r.Header.Get("Authorization")
			if auth == "" {
				http.Error(w, "Authorization header required", http.StatusUnauthorized)
				return
			}

			parts := strings.SplitN(auth, " ", 2)
			if len(parts) != 2 || parts[0] != "Basic" {
				http.Error(w, "Invalid authorization format", http.StatusUnauthorized)
				return
			}

			payload, err := base64.StdEncoding.DecodeString(parts[1])
			if err != nil {
				http.Error(w, "Invalid base64 encoding", http.StatusUnauthorized)
				return
			}

			cred := strings.SplitN(string(payload), ":", 2)
			if len(cred) != 2 {
				http.Error(w, "Invalid authorization value", http.StatusUnauthorized)
				return
			}

			username, password := cred[0], cred[1]
			account, err := svc.Login(r.Context(), username, password)
			if err != nil {
				http.Error(w, "Invalid credentials", http.StatusUnauthorized)
				return
			}

			ctx := context.WithValue(r.Context(), UserIDKey, account.ID)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
