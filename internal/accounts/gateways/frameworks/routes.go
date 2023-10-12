package frameworks

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jonloureiro/tiny-bank/internal/accounts"
	"github.com/jonloureiro/tiny-bank/internal/accounts/gateways"
	"github.com/jonloureiro/tiny-bank/internal/accounts/gateways/handlers"
)

type Routes struct {
	CreateAccountUsecase   accounts.CreateAccountUsecase
	CreateAccountPresenter gateways.CreateAccountPresenter

	ListAccountsUsecase   accounts.ListAccountsUsecase
	ListAccountsPresenter gateways.ListAccountsPresenter
}

func (r Routes) Setup() http.Handler {
	routes := chi.NewRouter()

	routes.Post("/accounts", handlers.CreateAccountHandler(r.CreateAccountUsecase, r.CreateAccountPresenter))
	routes.Get("/accounts", handlers.ListAccountsHandler(r.ListAccountsUsecase, r.ListAccountsPresenter))

	fmt.Println("Setup accounts routes")

	return routes
}
