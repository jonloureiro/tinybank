package account

import (
	"time"

	e "github.com/jonloureiro/tinybank/pkg/defaulterr"
)

// Errors used by account domain
var (
	ErrInvalidAccountName    = e.Wrap(e.FailedDependency(), "invalid account name")
	ErrInvalidAccountSecret  = e.Wrap(e.FailedDependency(), "invalid account secret")
	ErrIncompatibleCPFLength = e.Wrap(e.FailedDependency(), "incompatible cpf length")
	ErrInvalidAccountID      = e.Wrap(e.FailedDependency(), "invalid account id")
)

type Account interface {
	ID() string
	Name() string
	CPF() string
	Secret() string
	Balance() int
	CreatedAt() time.Time
}
