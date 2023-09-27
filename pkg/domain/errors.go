package domain

import "errors"

var (
	ErrBadRequest                 = errors.New("bad-request")
	ErrNotFound                   = errors.New("not-found")
	ErrInternal                   = errors.New("internal")
	ErrInvalidAuthorizationHeader = errors.New("invalid-authorization-header")
	ErrInvalidToken               = errors.New("invalid-token")
	ErrUnauthorized               = errors.New("unauthorized")
	ErrUserNotFound               = errors.New("user-not-found")
	ErrEnvironmentWithNameExists  = errors.New("environment-with-name-exists")
	ErrProjectWithNameExists      = errors.New("project-with-name-exists")
)
