package account

import (
	"context"
	"net/http"
)

type GetBalanceInput struct {
	AccountID string
}

type GetBalanceOutput struct {
	Balance int
}

type GetBalanceUsecase interface {
	Execute(context.Context, GetBalanceInput) (GetBalanceOutput, error)
}

type GetBalancePresenter interface {
	Render(http.ResponseWriter, GetBalanceOutput) error
}
