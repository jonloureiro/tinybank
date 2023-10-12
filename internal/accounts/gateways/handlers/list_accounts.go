package handlers

import (
	"errors"
	"net/http"

	"github.com/jonloureiro/tiny-bank/internal/accounts"
	"github.com/jonloureiro/tiny-bank/internal/accounts/gateways"
	"github.com/jonloureiro/tiny-bank/internal/common"
)

func ListAccountsHandler(
	uc accounts.ListAccountsUsecase,
	p gateways.ListAccountsPresenter,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		output, err := uc.Execute(ctx)

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
