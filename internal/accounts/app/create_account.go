package app

import (
	"context"

	"github.com/jonloureiro/tiny-bank/internal/accounts"
)

type AccountsRepository interface {
	accounts.SaveAccountsRepository
}

type CreateAccountUC struct {
	accRepo AccountsRepository
}

var _ accounts.CreateAccountUC = CreateAccountUC{}

func NewCreateAccountUC(
	accRepo AccountsRepository,
) CreateAccountUC {
	return CreateAccountUC{
		accRepo: accRepo,
	}
}

func (uc CreateAccountUC) CreateAccount(
	ctx context.Context,
	input accounts.CreateAccountInput,
) (accounts.CreateAccountOutput, error) {
	acc, err := NewAccount(
		input.Name, input.CPF, input.Secret,
	)
	if err != nil {
		return accounts.CreateAccountOutput{}, err
	}

	if err := uc.accRepo.Save(ctx, acc); err != nil {
		return accounts.CreateAccountOutput{}, err
	}

	return uc.BuildOutput(acc), nil
}

func (CreateAccountUC) BuildOutput(
	account accounts.Account,
) accounts.CreateAccountOutput {
	return accounts.CreateAccountOutput{
		AccountID: account.ID(),
	}
}
