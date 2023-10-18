package apihandler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/jonloureiro/tinybank/internal/account"
	"github.com/jonloureiro/tinybank/pkg/defaulterr"
)

type createBody struct {
	Name   string `json:"name"`
	CPF    string `json:"cpf"`
	Secret string `json:"secret"`
}

func Create(
	uc account.CreateUsecase,
	p account.CreatePresenter,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		var body createBody
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		output, err := uc.Execute(ctx, account.CreateInput{
			Name:   body.Name,
			CPF:    body.CPF,
			Secret: body.Secret,
		})

		switch {
		case err == nil:
			if err := p.Render(w, output); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
			}

		case errors.Is(err, defaulterr.FailedDependency()):
			w.WriteHeader(http.StatusUnprocessableEntity)

		case errors.Is(err, account.ErrAccountAlreadyExists):
			w.WriteHeader(http.StatusConflict)

		default:
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}
