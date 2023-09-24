package accounts

import (
	"context"
	"net/http"
)

type SaveAccountsRepository interface {
	Save(context.Context, Account) error
}

type CreateAccountInput struct {
	Name   string
	CPF    string
	Secret string
}

type CreateAccountOutput struct {
	AccountID string
}

type CreateAccountUsecase interface {
	Execute(
		context.Context,
		CreateAccountInput,
	) (CreateAccountOutput, error)
}

type CreateAccountPresenter interface {
	Render(http.ResponseWriter, CreateAccountOutput) error
}
