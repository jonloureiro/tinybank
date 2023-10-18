package test

import (
	"fmt"
	"time"

	"github.com/jonloureiro/tinybank/internal/account"
)

type fakeAccount struct {
	id      int
	uuid    string
	balance int
}

var _ account.Account = (*fakeAccount)(nil)

func Account(id int) fakeAccount {
	return fakeAccount{id: id}
}

func (a *fakeAccount) SetID(id string) {
	a.uuid = id
}

func (a *fakeAccount) SetBalance(balance int) {
	a.balance = balance
}

func (a fakeAccount) Balance() int {
	return a.balance
}

func (a fakeAccount) CPF() string {
	return fmt.Sprintf("fake-account-cpf-%d", a.id)
}

func (fakeAccount) CreatedAt() time.Time {
	return time.Time{}
}

func (a fakeAccount) ID() string {
	if a.uuid != "" {
		return a.uuid
	}
	return fmt.Sprintf("fake-account-id-%d", a.id)
}

func (fakeAccount) Name() string {
	return "fake-account-name"
}

func (fakeAccount) Secret() string {
	return "fake-account-secret"
}
