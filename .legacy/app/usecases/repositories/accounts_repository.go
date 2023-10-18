package repositories

import (
	"github.com/jonloureiro/tinybank/legacy/app/entities"
	"github.com/jonloureiro/tinybank/legacy/app/vo"
	"github.com/jonloureiro/tinybank/legacy/extensions/id"
)

type AccountsRepository interface {
	List() ([]*entities.Account, error)
	FindByID(id id.ID) (*entities.Account, error)
	FindByCPF(cpf *vo.CPF) (*entities.Account, error)
	Create(account *entities.Account) error
}
