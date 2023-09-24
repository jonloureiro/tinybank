package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/jonloureiro/tiny-bank/internal/accounts"
	"github.com/jonloureiro/tiny-bank/internal/common"
)

type createAccountBody struct {
	Name   string `json:"name"`
	CPF    string `json:"cpf"`
	Secret string `json:"secret"`
}

func createAccountHandler(
	uc accounts.CreateAccountUsecase,
	p accounts.CreateAccountPresenter,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		var body createAccountBody
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		output, err := uc.Execute(ctx, accounts.CreateAccountInput{
			Name:   body.Name,
			CPF:    body.CPF,
			Secret: body.Secret,
		})

		switch {
		case err == nil:
			if err := p.Render(w, output); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
			}

		case errors.Is(err, common.ErrFailedDependency):
			w.WriteHeader(http.StatusFailedDependency)

		case errors.Is(err, common.ErrConflict):
			w.WriteHeader(http.StatusConflict)

		default:
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}
