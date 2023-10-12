package frameworks

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jonloureiro/tiny-bank/internal/accounts"
	"github.com/jonloureiro/tiny-bank/internal/accounts/gateways"
	"github.com/jonloureiro/tiny-bank/internal/accounts/gateways/handlers"
)

type Routes struct {
	CreateAccountUsecase   accounts.CreateAccountUsecase
	CreateAccountPresenter gateways.CreateAccountPresenter
}

func (r Routes) Setup() http.Handler {
	routes := chi.NewRouter()

	routes.Post("/accounts", handlers.CreateAccountHandler(r.CreateAccountUsecase, r.CreateAccountPresenter))

	return routes
}
