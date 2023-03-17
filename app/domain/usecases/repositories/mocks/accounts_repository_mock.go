package mocks

import (
	"errors"

	"github.com/jonloureiro/tiny-bank/app/domain/entities"
	"github.com/jonloureiro/tiny-bank/app/domain/usecases"
	"github.com/jonloureiro/tiny-bank/app/domain/vo"
	"github.com/jonloureiro/tiny-bank/extensions/id"
)

const (
	CPFAlreadyExists    = "68347578133"
	CPFWithUnknownError = "70530694190"
)

type AccountsRepositoryMock struct {
	StorageByID  map[id.ID]*entities.Account
	StorageByCPF map[string]*entities.Account
	UnknownError bool
}

func NewAccountsRepositoryMock() *AccountsRepositoryMock {
	return &AccountsRepositoryMock{
		StorageByID:  make(map[id.ID]*entities.Account),
		StorageByCPF: make(map[string]*entities.Account),
	}
}

func (a *AccountsRepositoryMock) List() ([]*entities.Account, error) {
	if a.UnknownError {
		return nil, errors.New("unknown error")
	}
	accounts := make([]*entities.Account, 0)
	for _, account := range a.StorageByID {
		accounts = append(accounts, account)
	}
	return accounts, nil
}

func (a *AccountsRepositoryMock) FindByID(id id.ID) (*entities.Account, error) {
	account, ok := a.StorageByID[id]
	if !ok {
		return nil, usecases.ErrAccountNotFound
	}
	return account, nil
}

func (a *AccountsRepositoryMock) FindByCPF(cpf *vo.CPF) (*entities.Account, error) {
	if cpf.Value() == CPFAlreadyExists {
		account, _ := entities.NewAccount("Jon", cpf, "123456")
		return account, nil
	}
	account, ok := a.StorageByCPF[cpf.Value()]
	if !ok {
		return nil, usecases.ErrAccountNotFound
	}
	return account, nil
}

func (a *AccountsRepositoryMock) Create(account *entities.Account) error {
	if account.CPF.Value() == CPFWithUnknownError {
		return errors.New("unknown error")
	}
	if account.CPF.Value() == CPFAlreadyExists {
		return usecases.ErrAccountAlreadyExists
	}
	if _, ok := a.StorageByID[account.ID]; ok {
		return usecases.ErrAccountAlreadyExists
	}
	if _, ok := a.StorageByCPF[account.CPF.Value()]; ok {
		return usecases.ErrAccountAlreadyExists
	}
	a.StorageByID[account.ID] = account
	a.StorageByCPF[account.CPF.Value()] = account
	return nil
}
