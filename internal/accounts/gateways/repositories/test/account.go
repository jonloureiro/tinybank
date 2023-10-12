package test

import (
	"fmt"
	"time"

	"github.com/jonloureiro/tiny-bank/internal/accounts"
)

type account struct {
	id int
}

var _ accounts.Account = (*account)(nil)

func Account(id int) accounts.Account {
	return account{id}
}

func (account) Balance() int {
	return 0
}

func (a account) CPF() string {
	return fmt.Sprintf("fake-account-cpf-%d", a.id)
}

func (account) CreatedAt() time.Time {
	return time.Time{}
}

func (a account) ID() string {
	return fmt.Sprintf("fake-account-id-%d", a.id)
}

func (account) Name() string {
	return "fake-account-name"
}

func (account) Secret() string {
	return "fake-account-secret"
}
