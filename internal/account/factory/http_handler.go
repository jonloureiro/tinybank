package factory

import (
	"net/http"

	"github.com/jonloureiro/tinybank/internal/account/gateway/httphandler"
)

// HTTPHandler builds an HTTP handler for the account domain
func HTTPHandler() http.Handler {
	container := Container()
	return httphandler.New(container)
}
