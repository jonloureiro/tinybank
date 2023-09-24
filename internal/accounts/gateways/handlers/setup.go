package handlers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jonloureiro/tiny-bank/internal/accounts"
)

type handlers struct {
	createAccountUsecase   accounts.CreateAccountUsecase
	createAccountPresenter accounts.CreateAccountPresenter
}

func New(
	createAccountUsecase accounts.CreateAccountUsecase,
	createAccountPresenter accounts.CreateAccountPresenter,
) handlers {
	return handlers{
		createAccountUsecase:   createAccountUsecase,
		createAccountPresenter: createAccountPresenter,
	}
}

func (h handlers) Setup() http.Handler {
	r := chi.NewRouter()

	r.Post("/accounts", createAccountHandler(h.createAccountUsecase, h.createAccountPresenter))

	return r
}
