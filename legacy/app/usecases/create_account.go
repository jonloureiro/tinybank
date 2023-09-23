package usecases

import (
	"github.com/jonloureiro/tiny-bank/legacy/app/entities"
	"github.com/jonloureiro/tiny-bank/legacy/app/vo"
	"github.com/jonloureiro/tiny-bank/legacy/extensions/id"
)

type CreateAccountInput struct {
	Name   string
	CPF    string
	Secret string
}

type CreateAccountOutput struct {
	AccountID id.ID
}

func (uC *TinyBankUseCases) CreateAccount(input CreateAccountInput) (*CreateAccountOutput, error) {
	cpf, err := vo.NewCPF(input.CPF)
	if err != nil {
		return nil, err
	}
	_, err = uC.AccountsRepo.FindByCPF(cpf)
	if err == nil {
		return nil, ErrAccountAlreadyExists
	}
	account, err := entities.NewAccount(input.Name, cpf, input.Secret)
	if err != nil {
		return nil, err
	}
	err = uC.AccountsRepo.Create(account)
	if err == ErrAccountAlreadyExists {
		return nil, err
	}
	if err != nil {
		return nil, ErrDatabaseUnknownError
	}
	return &CreateAccountOutput{AccountID: account.ID}, nil
}
