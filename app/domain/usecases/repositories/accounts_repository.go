package repositories

import (
	"github.com/jonloureiro/tiny-bank/app/domain/entities"
	"github.com/jonloureiro/tiny-bank/app/domain/vo"
	"github.com/jonloureiro/tiny-bank/extensions/id"
)

type AccountsRepository interface {
	List() ([]*entities.Account, error)
	FindByID(id id.ID) (*entities.Account, error)
	FindByCPF(cpf *vo.CPF) (*entities.Account, error)
	Create(account *entities.Account) error
}
