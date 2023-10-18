package account

import (
	"context"
	"net/http"
)

type CreateInput struct {
	Name   string
	CPF    string
	Secret string
}

type CreateOutput struct {
	AccountID string
}

type CreateUsecase interface {
	Execute(context.Context, CreateInput) (CreateOutput, error)
}

type CreatePresenter interface {
	Render(http.ResponseWriter, CreateOutput) error
}
