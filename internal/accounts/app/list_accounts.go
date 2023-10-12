package app

import (
	"context"
	"sort"

	"github.com/jonloureiro/tiny-bank/internal/accounts"
)

type listAccounts_AccountsRepository interface {
	accounts.FindManyAccountsRepository
}

type listAccountsUsecase struct {
	accountsRepository listAccounts_AccountsRepository
}

var _ accounts.ListAccountsUsecase = (*listAccountsUsecase)(nil)

func NewListAccountsUsecase(
	accountsRepository listAccounts_AccountsRepository,
) listAccountsUsecase {
	return listAccountsUsecase{
		accountsRepository: accountsRepository,
	}
}

func (uc listAccountsUsecase) Execute(
	ctx context.Context,
) (accounts.ListAccountsOutput, error) {
	_accounts, err := uc.accountsRepository.FindMany(ctx)
	if err != nil {
		return accounts.ListAccountsOutput{}, err
	}

	sort.Slice(_accounts, func(i, j int) bool {
		return _accounts[i].ID() < _accounts[j].ID()
	})

	return uc.buildOutput(_accounts), nil
}

func (listAccountsUsecase) buildOutput(
	_accounts []accounts.Account,
) accounts.ListAccountsOutput {
	return accounts.ListAccountsOutput{
		Accounts: _accounts,
	}
}
