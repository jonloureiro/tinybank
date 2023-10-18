package apihandler

import (
	"errors"
	"net/http"

	"github.com/jonloureiro/tinybank/internal/account"
)

func List(
	uc account.ListUsecase,
	p account.ListPresenter,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		output, err := uc.Execute(ctx)

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
