package social

import "errors"

// common errors
var ErrInvalidArgument = errors.New("invalid argument")

var ErrBadRequest = errors.New("bad request")

var ErrNotFound = errors.New("not found")

var ErrInvalidRequestPayload = errors.New("invalid request payload")

// auth error
var ErrAuthInvalidUserCredentials = errors.New("auth: invalid credentials")

var ErrAuthNotAuthorized = errors.New("auth: not authorized")

var ErrAuthNotHavePermission = errors.New("auth: not have permission")

var ErrIncorrectUsernameOrPassword = errors.New("auth: incorrect username or password")

// error
var ErrUserAlreadyExists = errors.New("user already exit")

type CustomErrorBadRequest struct {
	err error
}

// Error implements error.
func (c CustomErrorBadRequest) Error() string {
	return c.err.Error()
}

func NewCustomErrorBadRequestMessage(err string) error {
	return CustomErrorBadRequest{err: errors.New(err)}
}
