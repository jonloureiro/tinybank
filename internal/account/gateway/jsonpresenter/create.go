package jsonpresenter

import (
	"encoding/json"
	"net/http"

	"github.com/jonloureiro/tinybank/internal/account"
)

type createPresenter struct{}

func NewCreate() createPresenter {
	return createPresenter{}
}

var _ account.CreatePresenter = (*createPresenter)(nil)

func (p createPresenter) Render(
	w http.ResponseWriter, output account.CreateOutput,
) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	return json.NewEncoder(w).Encode(accountIDSchema{
		AccountID: output.AccountID,
	})
}
