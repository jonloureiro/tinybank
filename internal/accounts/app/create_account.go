package app

import (
	"context"

	"github.com/jonloureiro/tiny-bank/internal/accounts"
	"github.com/jonloureiro/tiny-bank/internal/accounts/app/domain"
)

type createAccount_AccountsRepository interface {
	accounts.SaveAccountRepository
}

type createAccountUsecase struct {
	accountsRepository createAccount_AccountsRepository
}

var _ accounts.CreateAccountUsecase = (*createAccountUsecase)(nil)

func NewCreateAccountUsecase(
	accountsRepository createAccount_AccountsRepository,
) createAccountUsecase {
	return createAccountUsecase{
		accountsRepository: accountsRepository,
	}
}

func (uc createAccountUsecase) Execute(
	ctx context.Context,
	input accounts.CreateAccountInput,
) (accounts.CreateAccountOutput, error) {
	acc, err := domain.NewAccount(
		input.Name, input.CPF, input.Secret,
	)
	if err != nil {
		return accounts.CreateAccountOutput{}, err
	}

	if err := uc.accountsRepository.Save(ctx, acc); err != nil {
		return accounts.CreateAccountOutput{}, err
	}

	return uc.buildOutput(acc), nil
}

func (createAccountUsecase) buildOutput(
	account accounts.Account,
) accounts.CreateAccountOutput {
	return accounts.CreateAccountOutput{
		AccountID: account.ID(),
	}
}
