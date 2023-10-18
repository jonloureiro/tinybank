package usecase

import (
	"context"

	"github.com/jonloureiro/tinybank/internal/account"
	"github.com/jonloureiro/tinybank/internal/account/app/domain"
)

type create_AccountsRepository interface {
	account.SaveRepository
}

type create struct {
	accountsRepository create_AccountsRepository
}

var _ account.CreateUsecase = (*create)(nil)

func NewCreate(
	accountsRepository create_AccountsRepository,
) create {
	return create{
		accountsRepository: accountsRepository,
	}
}

func (uc create) Execute(
	ctx context.Context,
	input account.CreateInput,
) (account.CreateOutput, error) {
	acc, err := domain.NewAccount(
		input.Name, input.CPF, input.Secret,
	)
	if err != nil {
		return account.CreateOutput{}, err
	}

	if err := uc.accountsRepository.Save(ctx, acc); err != nil {
		return account.CreateOutput{}, err
	}

	return uc.buildOutput(acc), nil
}

func (create) buildOutput(
	acc account.Account,
) account.CreateOutput {
	return account.CreateOutput{
		AccountID: acc.ID(),
	}
}
