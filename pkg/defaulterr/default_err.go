package defaulterr

import (
	"errors"
	"fmt"
)

var (
	errConflict            = errors.New("conflict")
	errNotFound            = errors.New("not found")
	errMalformedParameters = errors.New("malformed parameters")
	errFailedDependency    = errors.New("failed dependency")
	errUnauthorized        = errors.New("unauthorized")
	errForbidden           = errors.New("forbidden")
	errLocked              = errors.New("locked")
	errUnknown             = errors.New("unknown")
)

func Wrap(err error, msg string) error {
	return fmt.Errorf("%w: %s", err, msg)
}

func Conflict() error {
	return errConflict
}

func NotFound() error {
	return errNotFound
}

func MalformedParameters() error {
	return errMalformedParameters
}

func FailedDependency() error {
	return errFailedDependency
}

func Unauthorized() error {
	return errUnauthorized
}

func Forbidden() error {
	return errForbidden
}

func Locked() error {
	return errLocked
}

func Unknown() error {
	return errUnknown
}
