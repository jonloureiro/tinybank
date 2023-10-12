package app

import (
	"context"

	"github.com/jonloureiro/tiny-bank/internal/accounts"
	"github.com/jonloureiro/tiny-bank/internal/accounts/app/domain"
)

type AccountsRepository interface {
	accounts.SaveAccountsRepository
}

type CreateAccountUsecase struct {
	accRepo AccountsRepository
}

var _ accounts.CreateAccountUsecase = (*CreateAccountUsecase)(nil)

func NewCreateAccountUsecase(
	accRepo AccountsRepository,
) CreateAccountUsecase {
	return CreateAccountUsecase{
		accRepo: accRepo,
	}
}

func (uc CreateAccountUsecase) Execute(
	ctx context.Context,
	input accounts.CreateAccountInput,
) (accounts.CreateAccountOutput, error) {
	acc, err := domain.NewAccount(
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

func (CreateAccountUsecase) BuildOutput(
	account accounts.Account,
) accounts.CreateAccountOutput {
	return accounts.CreateAccountOutput{
		AccountID: account.ID(),
	}
}
