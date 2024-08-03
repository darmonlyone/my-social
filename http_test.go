package social_test

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	social "github.com/darmonlyone/my-social"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

type mockService struct{}

func (m *mockService) Login(ctx context.Context, username string, password string) (*social.Account, error) {
	return &social.Account{ID: "123", Username: username, HashedPassword: password}, nil
}

func (m *mockService) FindAccount(ctx context.Context, id string) (*social.Account, error) {
	return &social.Account{ID: id, Username: "testuser"}, nil
}

func (m *mockService) StoreAccount(ctx context.Context, username string, password string, firstname string, lastname string) error {
	return nil
}

func (m *mockService) FindPost(ctx context.Context, id string) (*social.Post, error) {
	return &social.Post{ID: id, Title: "Test Post", Content: "This is a test post"}, nil
}

func (m *mockService) FindAllPosts(ctx context.Context) ([]*social.Post, error) {
	return []*social.Post{
		{ID: "1", Title: "Post 1", Content: "Content 1"},
		{ID: "2", Title: "Post 2", Content: "Content 2"},
	}, nil
}

func (m *mockService) StorePost(ctx context.Context, title string, content string) error {
	return nil
}

func (m *mockService) UpdatePost(ctx context.Context, id social.PostID, title *string, content *string) error {
	return nil
}

func (m *mockService) DeletePost(ctx context.Context, id string) error {
	return nil
}

func setupRouter() http.Handler {
	logger, _ := zap.NewProduction()
	svc := &mockService{}

	router := chi.NewRouter()
	router.Use(
		middleware.CleanPath,
		middleware.Timeout(60*time.Second),
		middleware.Recoverer,
	)

	opts := []kithttp.ServerOption{
		kithttp.ServerErrorEncoder(social.EncodeError),
	}

	endpoints := social.MakeSocialEndpoints(svc)

	loginHandler := social.MakeLoginHTTPHandler(opts, endpoints, logger, svc)
	registerHandler := social.MakeRegisterHTTPHandler(opts, endpoints, logger, svc)
	createPostHandler := social.MakeCreatePostHTTPHandler(opts, endpoints, logger, svc)
	listPostHandler := social.MakeListPostHTTPHandler(opts, endpoints, logger, svc)
	getPostHandler := social.MakeGetPostHTTPHandler(opts, endpoints, logger, svc)
	editPostHandler := social.MakeEditPostHTTPHandler(opts, endpoints, logger, svc)
	deletePostHandler := social.MakeDeletePostHTTPHandler(opts, endpoints, logger, svc)

	router.Group(func(r chi.Router) {
		r.Post("/login", loginHandler.ServeHTTP)
		r.Post("/register", registerHandler.ServeHTTP)

		r.With(social.BasicAuthMiddleware(svc)).Group(func(r chi.Router) {
			r.Post("/post.create", createPostHandler.ServeHTTP)
			r.Post("/post.list", listPostHandler.ServeHTTP)
			r.Post("/post.get", getPostHandler.ServeHTTP)
			r.Post("/post.edit", editPostHandler.ServeHTTP)
			r.Post("/post.delete", deletePostHandler.ServeHTTP)
		})
	})

	return router
}

func TestLoginHttp(t *testing.T) {
	router := setupRouter()

	reqBody := `{"username":"testuser","password":"testpass"}`
	req := httptest.NewRequest("POST", "/login", strings.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var response social.LoginResponse
	json.NewDecoder(w.Body).Decode(&response)
	assert.Equal(t, "123", response.UserID)
}

func TestRegisterHttp(t *testing.T) {
	router := setupRouter()

	reqBody := `{"username":"testuser","password":"testpass","firstname":"John","lastname":"Doe"}`
	req := httptest.NewRequest("POST", "/register", strings.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestFindPostHttp(t *testing.T) {
	router := setupRouter()

	reqBody := `{"id":"1"}`
	req := httptest.NewRequest("POST", "/post.get", strings.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Basic "+basicAuth("testuser", "testpass"))
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var response social.FindPostResponse
	json.NewDecoder(w.Body).Decode(&response)
	assert.Equal(t, "1", response.Post.ID)
	assert.Equal(t, "Test Post", response.Post.Title)
}

func TestFindAllPostsHttp(t *testing.T) {
	router := setupRouter()

	req := httptest.NewRequest("POST", "/post.list", nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Basic "+basicAuth("testuser", "testpass"))
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var response social.FindAllPostsResponse
	json.NewDecoder(w.Body).Decode(&response)
	assert.Equal(t, 2, len(response.Posts))
}

func TestStorePostHttp(t *testing.T) {
	router := setupRouter()

	reqBody := `{"title":"New Post","content":"New content"}`
	req := httptest.NewRequest("POST", "/post.create", strings.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Basic "+basicAuth("testuser", "testpass"))
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestUpdatePostHttp(t *testing.T) {
	router := setupRouter()

	reqBody := `{"id":"1","title":"Updated Post","content":"Updated content"}`
	req := httptest.NewRequest("POST", "/post.edit", strings.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Basic "+basicAuth("testuser", "testpass"))
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestDeletePostHttp(t *testing.T) {
	router := setupRouter()

	reqBody := `{"id":"1"}`
	req := httptest.NewRequest("POST", "/post.delete", strings.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Basic "+basicAuth("testuser", "testpass"))
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}
