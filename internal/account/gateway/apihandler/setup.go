package apihandler

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jonloureiro/tinybank/internal/account"
)

func New(container account.Container) http.Handler {
	router := chi.NewRouter()

	router.Post("/accounts",
		Create(container.CreateUsecase, container.CreatePresenter))

	router.Get("/accounts",
		List(container.ListUsecase, container.ListPresenter))

	router.Get("/accounts/{account-id}/balance",
		GetBalance(container.GetBalanceUsecase, container.GetBalancePresenter))

	fmt.Println("Setup accounts routes")

	return router
}
