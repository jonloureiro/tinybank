package main

import (
	"net/http"

	accountfactory "github.com/jonloureiro/tinybank/internal/account/factory"
)

func main() {
	accountHandler := accountfactory.HTTPHandler()

	if err := http.ListenAndServe(":3000", accountHandler); err != nil {
		panic(err)
	}
}
