package main

import (
	"net/http"

	accountsFactory "github.com/jonloureiro/tiny-bank/internal/accounts/gateways/factories"
)

func main() {
	accountsRoutes := accountsFactory.AccountsFactory()
	if err := http.ListenAndServe(":3000", accountsRoutes.Setup()); err != nil {
		panic(err)
	}
}
