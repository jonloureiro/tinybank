package presenters

import (
	"encoding/json"
	"net/http"

	"github.com/jonloureiro/tiny-bank/internal/accounts"
	"github.com/jonloureiro/tiny-bank/internal/accounts/gateways"
)

type getAccountBalanceJsonPresenter struct{}

func NewGetAccountBalanceJsonPresenter() getAccountBalanceJsonPresenter {
	return getAccountBalanceJsonPresenter{}
}

var _ gateways.GetAccountBalancePresenter = (*getAccountBalanceJsonPresenter)(nil)

func (p getAccountBalanceJsonPresenter) Render(
	w http.ResponseWriter, output accounts.GetAccountBalanceOutput,
) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	return json.NewEncoder(w).Encode(accountBalanceSchema{
		Balance: output.Balance,
	})
}
