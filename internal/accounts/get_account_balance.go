package accounts

import "context"

type GetAccountBalanceInput struct {
	AccountID string
}

type GetAccountBalanceOutput struct {
	Balance int
}

type GetAccountBalanceUsecase interface {
	Execute(
		context.Context,
		GetAccountBalanceInput,
	) (GetAccountBalanceOutput, error)
}
