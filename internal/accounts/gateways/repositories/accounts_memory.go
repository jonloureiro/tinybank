package repositories

import (
	"context"
	"fmt"

	"github.com/jonloureiro/tiny-bank/internal"
	"github.com/jonloureiro/tiny-bank/internal/accounts"
)

type accountsRepositoryInMemory struct {
	storageByCPF map[string]*accounts.Account
	storageByID  map[string]*accounts.Account
}

func NewAccountsRepositoryInMemory() *accountsRepositoryInMemory {
	return &accountsRepositoryInMemory{
		storageByCPF: make(map[string]*accounts.Account),
		storageByID:  make(map[string]*accounts.Account),
	}
}

func (repo *accountsRepositoryInMemory) Save(
	ctx context.Context,
	acc accounts.Account,
) error {
	if _, ok := repo.storageByCPF[acc.CPF()]; ok {
		return fmt.Errorf("%w:account already exists", internal.ErrConflict)
	}

	if _, ok := repo.storageByID[acc.ID()]; ok {
		return fmt.Errorf("%w:account already exists", internal.ErrConflict)
	}

	repo.storageByCPF[acc.CPF()] = &acc
	repo.storageByID[acc.ID()] = &acc

	return nil
}
