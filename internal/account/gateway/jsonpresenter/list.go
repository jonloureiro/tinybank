package jsonpresenter

import (
	"encoding/json"
	"net/http"

	"github.com/jonloureiro/tinybank/internal/account"
)

type listPresenter struct{}

func NewList() listPresenter {
	return listPresenter{}
}

var _ account.ListPresenter = (*listPresenter)(nil)

func (p listPresenter) Render(
	w http.ResponseWriter, output account.ListOutput,
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
