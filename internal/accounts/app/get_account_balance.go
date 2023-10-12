package app

import (
	"context"

	"github.com/jonloureiro/tiny-bank/internal/accounts"
	"github.com/jonloureiro/tiny-bank/internal/accounts/app/domain/vos"
)

type getAccountBalance_AccountsRepository interface {
	accounts.FindAccountRepository
}

type getAccountBalanceUsecase struct {
	accountsRepository getAccountBalance_AccountsRepository
}

var _ accounts.GetAccountBalanceUsecase = (*getAccountBalanceUsecase)(nil)

func NewGetAccountBalanceUsecase(
	accountsRepository getAccountBalance_AccountsRepository,
) getAccountBalanceUsecase {
	return getAccountBalanceUsecase{
		accountsRepository: accountsRepository,
	}
}

func (uc getAccountBalanceUsecase) Execute(
	ctx context.Context, input accounts.GetAccountBalanceInput,
) (accounts.GetAccountBalanceOutput, error) {
	parsedAccountID, err := vos.ParseAccountID(input.AccountID)
	if err != nil {
		return accounts.GetAccountBalanceOutput{}, err
	}

	account, err := uc.accountsRepository.Find(ctx, parsedAccountID.Value())
	if err != nil {
		return accounts.GetAccountBalanceOutput{}, err
	}

	return uc.buildOutput(account), nil
}

func (getAccountBalanceUsecase) buildOutput(
	account accounts.Account,
) accounts.GetAccountBalanceOutput {
	return accounts.GetAccountBalanceOutput{
		Balance: account.Balance(),
	}
}
