package accounts

import (
	"context"
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
