package factory

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jonloureiro/tinybank/internal/account/gateway/apihandler"
)

// HTTPHandler builds an HTTP handler for the account domain
func HTTPHandler() http.Handler {
	r := chi.NewRouter()
	container := Container()
	apiHandler := apihandler.New(container)
	r.Mount("/api", apiHandler)
	return r
}
