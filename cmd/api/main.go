package main

import (
	"net/http"

	accountsApp "github.com/jonloureiro/tiny-bank/internal/accounts/app"
	accountsHTTPHandlers "github.com/jonloureiro/tiny-bank/internal/accounts/gateways/handlers"
	accountsPresenters "github.com/jonloureiro/tiny-bank/internal/accounts/gateways/presenters"
	accountsRepositories "github.com/jonloureiro/tiny-bank/internal/accounts/gateways/repositories"
)

func main() {
	var (
		accountsRepositoryInMemory = accountsRepositories.NewRepositoryInMemory()
		createAccountUsecase       = accountsApp.NewCreateAccountUsecase(accountsRepositoryInMemory)
		createAccountJsonPresenter = accountsPresenters.NewJsonPresenter()
	)

	accountsHandlers := accountsHTTPHandlers.New(
		createAccountUsecase,
		createAccountJsonPresenter,
	)
	if err := http.ListenAndServe(":3000", accountsHandlers.Setup()); err != nil {
		panic(err)
	}
}
