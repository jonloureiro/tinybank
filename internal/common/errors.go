package common

import "errors"

var (
	ErrConflict            = errors.New("conflict")
	ErrNotFound            = errors.New("not found")
	ErrMalformedParameters = errors.New("malformed parameters")
	ErrFailedDependency    = errors.New("failed dependency")
	ErrUnauthorized        = errors.New("unauthorized")
	ErrForbidden           = errors.New("forbidden")
	ErrLocked              = errors.New("locked")
	ErrUnknown             = errors.New("unknown")
)
