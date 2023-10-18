package usecase

import (
	"context"

	"github.com/jonloureiro/tinybank/internal/account"
	"github.com/jonloureiro/tinybank/internal/account/app/vo"
)

type getBalance_AccountsRepository interface {
	account.FindRepository
}

type getBalance struct {
	accountsRepository getBalance_AccountsRepository
}

var _ account.GetBalanceUsecase = (*getBalance)(nil)

func NewGetBalance(
	accountsRepository getBalance_AccountsRepository,
) getBalance {
	return getBalance{
		accountsRepository: accountsRepository,
	}
}

func (uc getBalance) Execute(
	ctx context.Context, input account.GetBalanceInput,
) (account.GetBalanceOutput, error) {
	parsedAccountID, err := vo.ParseAccountID(input.AccountID)
	if err != nil {
		return account.GetBalanceOutput{}, err
	}

	acc, err := uc.accountsRepository.Find(ctx, parsedAccountID.Value())
	if err != nil {
		return account.GetBalanceOutput{}, err
	}

	return uc.buildOutput(acc), nil
}

func (getBalance) buildOutput(
	acc account.Account,
) account.GetBalanceOutput {
	return account.GetBalanceOutput{
		Balance: acc.Balance(),
	}
}
