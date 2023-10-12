package handlers

import (
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jonloureiro/tiny-bank/internal/accounts"
	"github.com/jonloureiro/tiny-bank/internal/accounts/gateways"
	"github.com/jonloureiro/tiny-bank/internal/common"
)

func GetAccountBalanceHandler(
	uc accounts.GetAccountBalanceUsecase,
	p gateways.GetAccountBalancePresenter,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		accountID := chi.URLParam(r, "account-id")

		output, err := uc.Execute(ctx, accounts.GetAccountBalanceInput{
			AccountID: accountID,
		})

		switch {
		case err == nil:
			if err := p.Render(w, output); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
			}

		case errors.Is(err, common.ErrNotFound):
			w.WriteHeader(http.StatusNotFound)

		default:
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}
