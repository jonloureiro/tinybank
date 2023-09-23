package main

import (
	"net/http"

	accountsApp "github.com/jonloureiro/tiny-bank/internal/accounts/app"
	accountsApi "github.com/jonloureiro/tiny-bank/internal/accounts/gateways/api"
	accountsRepo "github.com/jonloureiro/tiny-bank/internal/accounts/gateways/repositories/inmemory"
)

func main() {
	// ------------------------ Repositories
	accRepo := accountsRepo.NewAccountsRepositoryInMemory()

	// ------------------------ Usecases
	createAccountUC := accountsApp.NewCreateAccountUC(accRepo)

	// ------------------------ HTTP
	httpRoutes := accountsApi.NewHttpRoutes(createAccountUC)
	httpHandler := httpRoutes.Setup()

	if err := http.ListenAndServe(":3000", httpHandler); err != nil {
		panic(err)
	}
}
