package inmemoryrepository

import (
	"context"

	"github.com/jonloureiro/tinybank/internal/account"
)

var _ account.SaveRepository = (*inMemoryRepository)(nil)

func (repo *inMemoryRepository) Save(
	ctx context.Context,
	acc account.Account,
) error {
	if _, ok := repo.storageByCPF[acc.CPF()]; ok {
		return account.ErrAccountAlreadyExists
	}

	if _, ok := repo.storageByID[acc.ID()]; ok {
		return account.ErrAccountAlreadyExists
	}

	repo.storageByCPF[acc.CPF()] = &acc
	repo.storageByID[acc.ID()] = &acc

	return nil
}
