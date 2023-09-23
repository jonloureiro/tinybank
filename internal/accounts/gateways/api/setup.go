package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jonloureiro/tiny-bank/internal/accounts"
	"github.com/jonloureiro/tiny-bank/internal/accounts/gateways/api/handlers"
)

type httpRoutes struct {
	createAccountUC accounts.CreateAccountUC
}

func NewHttpRoutes(
	createAccountUC accounts.CreateAccountUC,
) httpRoutes {
	return httpRoutes{
		createAccountUC: createAccountUC,
	}
}

func (h httpRoutes) Setup() http.Handler {
	r := chi.NewRouter()

	r.Post("/accounts", handlers.CreateAccountHandler(h.createAccountUC))

	return r
}
