package repositories

import (
	"context"
	"fmt"

	"github.com/jonloureiro/tiny-bank/internal/accounts"
	"github.com/jonloureiro/tiny-bank/internal/common"
)

var _ accounts.SaveAccountRepository = (*accountsRepositoryInMemory)(nil)

func (repo *accountsRepositoryInMemory) Save(
	ctx context.Context,
	acc accounts.Account,
) error {
	if _, ok := repo.storageByCPF[acc.CPF()]; ok {
		return fmt.Errorf("%w:account already exists", common.ErrConflict)
	}

	if _, ok := repo.storageByID[acc.ID()]; ok {
		return fmt.Errorf("%w:account already exists", common.ErrConflict)
	}

	repo.storageByCPF[acc.CPF()] = &acc
	repo.storageByID[acc.ID()] = &acc

	return nil
}
