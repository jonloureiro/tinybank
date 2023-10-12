package gateways

import (
	"net/http"

	"github.com/jonloureiro/tiny-bank/internal/accounts"
)

type CreateAccountPresenter interface {
	Render(http.ResponseWriter, accounts.CreateAccountOutput) error
}

type ListAccountsPresenter interface {
	Render(http.ResponseWriter, accounts.ListAccountsOutput) error
}

type GetAccountBalancePresenter interface {
	Render(http.ResponseWriter, accounts.GetAccountBalanceOutput) error
}
