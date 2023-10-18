package apihandler

import (
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jonloureiro/tinybank/internal/account"
)

func GetBalance(
	uc account.GetBalanceUsecase,
	p account.GetBalancePresenter,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		accountID := chi.URLParam(r, "account-id")

		output, err := uc.Execute(ctx, account.GetBalanceInput{
			AccountID: accountID,
		})

		switch {
		case err == nil:
			if err := p.Render(w, output); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
			}

		case errors.Is(err, account.ErrAccountNotFound):
			w.WriteHeader(http.StatusNotFound)

		default:
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}
