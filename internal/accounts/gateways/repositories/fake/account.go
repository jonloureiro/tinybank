package fake

import (
	"time"

	"github.com/jonloureiro/tiny-bank/internal/accounts"
)

type account struct{}

var _ accounts.Account = account{}

func Account() accounts.Account {
	return account{}
}

func (account) Balance() int {
	return 0
}

func (account) CPF() string {
	return "fake-account-cpf"
}

func (account) CreatedAt() time.Time {
	return time.Time{}
}

func (account) ID() string {
	return "fake-account-id"
}

func (account) Name() string {
	return "fake-account-name"
}

func (account) Secret() string {
	return "fake-account-secret"
}
