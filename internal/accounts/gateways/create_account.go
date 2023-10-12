package gateways

import (
	"net/http"

	"github.com/jonloureiro/tiny-bank/internal/accounts"
)

type CreateAccountPresenter interface {
	Render(http.ResponseWriter, accounts.CreateAccountOutput) error
}
