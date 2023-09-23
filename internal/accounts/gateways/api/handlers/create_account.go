package handlers

import (
	"context"
	"errors"
	"net/http"

	"github.com/jonloureiro/tiny-bank/internal"
	"github.com/jonloureiro/tiny-bank/internal/accounts/gateways/api/schemas"
	"github.com/jonloureiro/tiny-bank/internal/accounts/usecases"
	"github.com/jonloureiro/tiny-bank/pkg/rest"
)

type CreateAccountUC interface {
	CreateAccount(
		context.Context,
		usecases.CreateAccountInput,
	) (usecases.CreateAccountOutput, error)
}

func CreateAccountHandler(uc CreateAccountUC) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		var body schemas.CreateAccountRequest
		if err := rest.Decode(r, &body); err != nil {
			rest.Status(w, http.StatusBadRequest)
			return
		}

		output, err := uc.CreateAccount(ctx, usecases.CreateAccountInput{
			Name:   body.Name,
			CPF:    body.CPF,
			Secret: body.Secret,
		})

		switch {
		case err == nil:
			rest.Encode(w, schemas.CreateAccountResponse{
				AccountID: output.AccountID,
			})

		case errors.Is(err, internal.ErrFailedDependency):
			rest.Status(w, http.StatusUnprocessableEntity)
			return

		case errors.Is(err, internal.ErrConflict):
			rest.Status(w, http.StatusConflict)
			return

		default:
			rest.Status(w, http.StatusInternalServerError)
			return
		}
	}
}
