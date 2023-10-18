package account

import (
	"context"
	"net/http"
)

type ListOutput struct {
	Accounts []Account
}

type ListUsecase interface {
	Execute(context.Context) (ListOutput, error)
}

type ListPresenter interface {
	Render(http.ResponseWriter, ListOutput) error
}
