package account

import (
	"context"

	e "github.com/jonloureiro/tinybank/pkg/defaulterr"
)

// Errors used by account repository
var (
	ErrAccountAlreadyExists = e.Wrap(e.Conflict(), "account already exists")
	ErrAccountNotFound      = e.Wrap(e.NotFound(), "account not found")
)

type FindRepository interface {
	Find(context.Context, string) (Account, error)
}

type FindManyRepository interface {
	FindMany(context.Context) ([]Account, error)
}

type SaveRepository interface {
	Save(context.Context, Account) error
}
