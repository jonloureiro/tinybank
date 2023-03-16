package usecases

import (
	"github.com/jonloureiro/tiny-bank/app/domain/entities"
	"github.com/jonloureiro/tiny-bank/app/domain/vo"
	"github.com/jonloureiro/tiny-bank/extensions/id"
)

type AccountsRepository interface {
	FindAccountByID(id id.ID) (*entities.Account, error)
	FindAccountByCPF(cpf *vo.CPF) (*entities.Account, error)
	CreateAccount(account *entities.Account) error
}

type TinyBankUseCases struct {
	AccountsRepo AccountsRepository
}
