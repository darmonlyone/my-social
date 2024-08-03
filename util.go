package social

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	kithttp "github.com/go-kit/kit/transport/http"
)

const userIDKey string = "userID"

func UserIDFromContext(ctx context.Context) (string, bool) {
	userID, ok := ctx.Value(userIDKey).(string)
	return userID, ok
}

type ErrorEncoder func(ctx context.Context, err error, w http.ResponseWriter)

type errorer interface {
	error() error
}

func MakeEncodeResponse(encodeError ErrorEncoder, status int) kithttp.EncodeResponseFunc {
	return func(ctx context.Context, w http.ResponseWriter, response interface{}) error {
		if e, ok := response.(errorer); ok && e.error() != nil {
			encodeError(ctx, e.error(), w)
			return nil
		}
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(status)
		return json.NewEncoder(w).Encode(response)
	}
}

func DecodeMappingBodyRequest[T any](_ context.Context, r *http.Request) (interface{}, error) {
	if r.Body == nil || r.ContentLength == 0 {
		// If the body is empty, return the default value.
		return nil, nil
	}
	var req T
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, ErrInvalidRequestPayload
	}
	return req, nil
}

func HandleCommonErrors(err error, w http.ResponseWriter) {
	if errors.As(err, &CustomErrorBadRequest{}) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	switch err {
	case ErrIncorrectUsernameOrPassword:
		w.WriteHeader(http.StatusUnauthorized)
	case ErrAuthNotAuthorized:
		w.WriteHeader(http.StatusUnauthorized)
	case ErrUserAlreadyExists:
		w.WriteHeader(http.StatusBadRequest)
	case ErrNotFound:
		w.WriteHeader(http.StatusNotFound)
	case ErrBadRequest:
		w.WriteHeader(http.StatusBadRequest)
	case ErrInvalidRequestPayload:
		w.WriteHeader(http.StatusBadRequest)
	case ErrAuthNotHavePermission:
		w.WriteHeader(http.StatusForbidden)
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}
}
