package main

import (
	"net/http"

	accountsApp "github.com/jonloureiro/tiny-bank/internal/accounts/app"
	accountsFrameworks "github.com/jonloureiro/tiny-bank/internal/accounts/gateways/frameworks"
	accountsPresenters "github.com/jonloureiro/tiny-bank/internal/accounts/gateways/presenters"
	accountsRepositories "github.com/jonloureiro/tiny-bank/internal/accounts/gateways/repositories"
)

func main() {
	var (
		accountsRepositoryInMemory = accountsRepositories.NewRepositoryInMemory()
		createAccountUsecase       = accountsApp.NewCreateAccountUsecase(accountsRepositoryInMemory)
		createAccountJsonPresenter = accountsPresenters.NewJsonPresenter()
	)

	routes := accountsFrameworks.Routes{
		CreateAccountUsecase:   createAccountUsecase,
		CreateAccountPresenter: createAccountJsonPresenter,
	}
	if err := http.ListenAndServe(":3000", routes.Setup()); err != nil {
		panic(err)
	}
}
