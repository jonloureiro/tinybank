package repositories

import (
	"context"

	"github.com/jonloureiro/tiny-bank/internal/accounts"
)

var _ accounts.FindManyAccountsRepository = (*accountsRepositoryInMemory)(nil)

func (repo *accountsRepositoryInMemory) FindMany(
	ctx context.Context,
) ([]accounts.Account, error) {
	_accounts := make([]accounts.Account, 0, len(repo.storageByID))
	for _, acc := range repo.storageByID {
		_accounts = append(_accounts, *acc)
	}

	return _accounts, nil
}
