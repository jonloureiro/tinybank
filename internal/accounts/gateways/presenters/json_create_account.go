package presenters

import (
	"encoding/json"
	"net/http"

	"github.com/jonloureiro/tiny-bank/internal/accounts"
	"github.com/jonloureiro/tiny-bank/internal/accounts/gateways"
)

type createAccountJsonPresenter struct{}

func NewCreateAccountJsonPresenter() createAccountJsonPresenter {
	return createAccountJsonPresenter{}
}

var _ gateways.CreateAccountPresenter = (*createAccountJsonPresenter)(nil)

func (p createAccountJsonPresenter) Render(
	w http.ResponseWriter, output accounts.CreateAccountOutput,
) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	return json.NewEncoder(w).Encode(accountIDSchema{
		AccountID: output.AccountID,
	})
}
