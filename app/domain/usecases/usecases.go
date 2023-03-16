package usecases

import (
	"github.com/jonloureiro/go-challenge/app/domain/entities"
	"github.com/jonloureiro/go-challenge/app/domain/vo"
	"github.com/jonloureiro/go-challenge/extensions/id"
)

type AccountsRepository interface {
	FindAccountByID(id id.ID) (*entities.Account, error)
	FindAccountByCPF(cpf *vo.CPF) (*entities.Account, error)
	CreateAccount(account *entities.Account) error
}

type TinyBankUseCases struct {
	AccountsRepo AccountsRepository
}
