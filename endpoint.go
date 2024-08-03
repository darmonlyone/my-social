package social

import (
	"context"

	"github.com/go-kit/kit/endpoint"
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
		err := svc.Login(ctx, req.Username, req.Password)
		if err != nil {
			return nil, err
		}
		return EmptyResponse{}, nil
	}
}

// Account Requests and Responses
type LoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func makeRegisterEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(RegisterRequest)
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
		err := svc.StorePost(ctx, req.CreatedBy, req.Title, req.Content)
		if err != nil {
			return nil, err
		}
		return EmptyResponse{}, nil
	}
}

type StorePostRequest struct {
	CreatedBy AccountID `json:"createdBy" validate:"required"`
	Title     string    `json:"title" validate:"required"`
	Content   string    `json:"content" validate:"required"`
}

func makeUpdatePostEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdatePostRequest)
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
