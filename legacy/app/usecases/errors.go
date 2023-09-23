package usecases

import "errors"

var (
	ErrAccountNotFound      = errors.New("account not found")
	ErrAccountAlreadyExists = errors.New("account already exists")
	ErrDatabaseUnknownError = errors.New("database unknown error")
)
