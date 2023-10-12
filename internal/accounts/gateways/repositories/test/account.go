package test

import (
	"fmt"
	"time"

	"github.com/jonloureiro/tiny-bank/internal/accounts"
)

type account struct {
	id      int
	uuid    string
	balance int
}

var _ accounts.Account = (*account)(nil)

func Account(id int) account {
	return account{id: id}
}

func (a *account) SetID(id string) {
	a.uuid = id
}

func (a *account) SetBalance(balance int) {
	a.balance = balance
}

func (a account) Balance() int {
	return a.balance
}

func (a account) CPF() string {
	return fmt.Sprintf("fake-account-cpf-%d", a.id)
}

func (account) CreatedAt() time.Time {
	return time.Time{}
}

func (a account) ID() string {
	if a.uuid != "" {
		return a.uuid
	}
	return fmt.Sprintf("fake-account-id-%d", a.id)
}

func (account) Name() string {
	return "fake-account-name"
}

func (account) Secret() string {
	return "fake-account-secret"
}
