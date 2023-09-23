package usecases

import (
	"context"

	"github.com/jonloureiro/tiny-bank/internal/accounts"
)

type AccountsRepository interface {
	Save(context.Context, accounts.Account) error
}

type CreateAccountUC struct {
	accRepo AccountsRepository
}

func NewCreateAccountUC(
	accRepo AccountsRepository,
) CreateAccountUC {
	return CreateAccountUC{
		accRepo: accRepo,
	}
}

type CreateAccountInput struct {
	Name   string
	CPF    string
	Secret string
}

type CreateAccountOutput struct {
	AccountID string
}

func (uc CreateAccountUC) CreateAccount(
	ctx context.Context,
	input CreateAccountInput,
) (CreateAccountOutput, error) {
	acc, err := accounts.New(
		input.Name, input.CPF, input.Secret,
	)
	if err != nil {
		return CreateAccountOutput{}, err
	}

	if err := uc.accRepo.Save(ctx, acc); err != nil {
		return CreateAccountOutput{}, err
	}

	return uc.BuildOutput(acc), nil
}

func (CreateAccountUC) BuildOutput(
	account accounts.Account,
) CreateAccountOutput {
	return CreateAccountOutput{
		AccountID: account.ID(),
	}
}
