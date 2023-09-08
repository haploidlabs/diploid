package domain

import "errors"

var (
	ErrBadRequest                 = errors.New("bad-request")
	ErrInternal                   = errors.New("internal")
	ErrInvalidAuthorizationHeader = errors.New("invalid-authorization-header")
	ErrInvalidToken               = errors.New("invalid-token")
	ErrUnauthorized               = errors.New("unauthorized")
	ErrUserNotFound               = errors.New("user-not-found")
)
