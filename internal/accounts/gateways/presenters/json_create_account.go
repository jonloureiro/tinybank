package presenters

import (
	"encoding/json"
	"net/http"

	"github.com/jonloureiro/tiny-bank/internal/accounts"
)

type createAccountTarget struct {
	AccountID string `json:"id"`
}

var _ accounts.CreateAccountPresenter = createAccountJsonPresenter{}

func (p createAccountJsonPresenter) Render(
	w http.ResponseWriter, output accounts.CreateAccountOutput,
) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	return json.NewEncoder(w).Encode(createAccountTarget{
		AccountID: output.AccountID,
	})
}
