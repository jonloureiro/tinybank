package main

import (
	"net/http"

	"github.com/jonloureiro/tiny-bank/internal/accounts/gateways/api"
	"github.com/jonloureiro/tiny-bank/internal/accounts/gateways/repositories"
	"github.com/jonloureiro/tiny-bank/internal/accounts/usecases"
)

func main() {
	// ------------------------ Repositories
	accRepo := repositories.NewAccountsRepositoryInMemory()

	// ------------------------ Usecases
	createAccountUC := usecases.NewCreateAccountUC(accRepo)

	// ------------------------ HTTP
	httpRoutes := api.NewHttpRoutes(createAccountUC)
	httpHandler := httpRoutes.Setup()

	if err := http.ListenAndServe(":3000", httpHandler); err != nil {
		panic(err)
	}
}
