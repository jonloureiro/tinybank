package usecase

import (
	"context"
	"sort"

	"github.com/jonloureiro/tinybank/internal/account"
)

type list_AccountsRepository interface {
	account.FindManyRepository
}

type list struct {
	accountsRepository list_AccountsRepository
}

var _ account.ListUsecase = (*list)(nil)

func NewList(
	accountsRepository list_AccountsRepository,
) list {
	return list{
		accountsRepository: accountsRepository,
	}
}

func (uc list) Execute(
	ctx context.Context,
) (account.ListOutput, error) {
	_accounts, err := uc.accountsRepository.FindMany(ctx)
	if err != nil {
		return account.ListOutput{}, err
	}

	sort.Slice(_accounts, func(i, j int) bool {
		return _accounts[i].ID() < _accounts[j].ID()
	})

	return uc.buildOutput(_accounts), nil
}

func (list) buildOutput(
	_accounts []account.Account,
) account.ListOutput {
	return account.ListOutput{
		Accounts: _accounts,
	}
}
