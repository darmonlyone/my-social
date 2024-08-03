package social

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-playground/validator"
)

type SocialEndpoints struct {
	Login    endpoint.Endpoint
	Register endpoint.Endpoint

	FindPost     endpoint.Endpoint
	FindAllPosts endpoint.Endpoint
	StorePost    endpoint.Endpoint
	UpdatePost   endpoint.Endpoint
	DeletePost   endpoint.Endpoint
}

func MakeSocialEndpoints(svc Service) SocialEndpoints {
	return SocialEndpoints{
		Login:    makeLoginEndpoint(svc),
		Register: makeRegisterEndpoint(svc),

		FindPost:     makeFindPostEndpoint(svc),
		FindAllPosts: makeFindAllPostsEndpoint(svc),
		StorePost:    makeStorePostEndpoint(svc),
		UpdatePost:   makeUpdatePostEndpoint(svc),
		DeletePost:   makeDeletePostEndpoint(svc),
	}
}

func makeLoginEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(LoginRequest)
		account, err := svc.Login(ctx, req.Username, req.Password)
		if err != nil {
			return nil, err
		}
		return LoginResponse{UserID: account.ID}, nil
	}
}

var validate = validator.New()

// Account Requests and Responses
type LoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type LoginResponse struct {
	UserID string `json:"userId"`
}

func makeRegisterEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(RegisterRequest)
		if err := validate.Struct(req); err != nil {
			return nil, NewCustomErrorBadRequest(err)
		}
		err := svc.StoreAccount(ctx, req.Username, req.Password, req.Firstname, req.Lastname)
		if err != nil {
			return nil, err
		}
		return EmptyResponse{}, nil
	}
}

type RegisterRequest struct {
	Username  string `json:"username" validate:"required"`
	Password  string `json:"password" validate:"required"`
	Firstname string `json:"firstname" validate:"required"`
	Lastname  string `json:"lastname" validate:"required"`
}

func makeFindPostEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(FindPostRequest)
		if err := validate.Struct(req); err != nil {
			return nil, NewCustomErrorBadRequest(err)
		}
		post, err := svc.FindPost(ctx, req.ID)
		if err != nil {
			return nil, err
		}
		return FindPostResponse{Post: post}, nil
	}
}

// Post Requests and Responses
type FindPostRequest struct {
	ID string `json:"id" validate:"required"`
}

type FindPostResponse struct {
	Post *Post `json:"post" validate:"required"`
}

func makeFindAllPostsEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		posts, err := svc.FindAllPosts(ctx)
		if err != nil {
			return nil, err
		}
		return FindAllPostsResponse{Posts: posts}, nil
	}
}

type FindAllPostsResponse struct {
	Posts []*Post `json:"posts" validate:"required"`
}

func makeStorePostEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(StorePostRequest)
		if err := validate.Struct(req); err != nil {
			return nil, NewCustomErrorBadRequest(err)
		}
		err := svc.StorePost(ctx, req.Title, req.Content)
		if err != nil {
			return nil, err
		}
		return EmptyResponse{}, nil
	}
}

type StorePostRequest struct {
	Title     string    `json:"title" validate:"required"`
	Content   string    `json:"content" validate:"required"`
}

func makeUpdatePostEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdatePostRequest)
		if err := validate.Struct(req); err != nil {
			return nil, NewCustomErrorBadRequest(err)
		}
		err := svc.UpdatePost(ctx, req.ID, req.Title, req.Content)
		if err != nil {
			return nil, err
		}
		return EmptyResponse{}, nil
	}
}

type UpdatePostRequest struct {
	ID      PostID  `json:"id" validate:"required"`
	Title   *string `json:"title,omitempty"`
	Content *string `json:"content,omitempty"`
}

func makeDeletePostEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeletePostRequest)
		if err := validate.Struct(req); err != nil {
			return nil, NewCustomErrorBadRequest(err)
		}
		err := svc.DeletePost(ctx, req.ID)
		if err != nil {
			return nil, err
		}
		return EmptyResponse{}, nil
	}
}

type DeletePostRequest struct {
	ID string `json:"id" validate:"required"`
}

type EmptyResponse struct {
}
