package accounts

import (
	"context"
)

type ListAccountsOutput struct {
	Accounts []Account
}

type ListAccountsUsecase interface {
	Execute(
		context.Context,
	) (ListAccountsOutput, error)
}
