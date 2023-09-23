package accounts

import "context"

type CreateAccountInput struct {
	Name   string
	CPF    string
	Secret string
}

type CreateAccountOutput struct {
	AccountID string
}

type CreateAccountUC interface {
	CreateAccount(
		context.Context,
		CreateAccountInput,
	) (CreateAccountOutput, error)
}
