package presenters

import (
	"encoding/json"
	"net/http"

	"github.com/jonloureiro/tiny-bank/internal/accounts"
	"github.com/jonloureiro/tiny-bank/internal/accounts/gateways"
)

type listAccountsJsonPresenter struct{}

func NewListAccountsJsonPresenter() listAccountsJsonPresenter {
	return listAccountsJsonPresenter{}
}

var _ gateways.ListAccountsPresenter = (*listAccountsJsonPresenter)(nil)

func (p listAccountsJsonPresenter) Render(
	w http.ResponseWriter, output accounts.ListAccountsOutput,
) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	responseBody := make([]accountSchema, 0, len(output.Accounts))
	for _, account := range output.Accounts {
		responseBody = append(responseBody, accountSchema{
			ID:        account.ID(),
			Name:      account.Name(),
			CPF:       account.CPF(),
			Balance:   account.Balance(),
			CreatedAt: account.CreatedAt().String(),
		})
	}

	return json.NewEncoder(w).Encode(accountsSchema(responseBody))
}
