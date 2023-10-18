package jsonpresenter

import (
	"encoding/json"
	"net/http"

	"github.com/jonloureiro/tinybank/internal/account"
)

type getBalancePresenter struct{}

func NewGetBalance() getBalancePresenter {
	return getBalancePresenter{}
}

var _ account.GetBalancePresenter = (*getBalancePresenter)(nil)

func (p getBalancePresenter) Render(
	w http.ResponseWriter, output account.GetBalanceOutput,
) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	return json.NewEncoder(w).Encode(accountBalanceSchema{
		Balance: output.Balance,
	})
}
