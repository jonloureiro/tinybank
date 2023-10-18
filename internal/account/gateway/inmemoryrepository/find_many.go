package inmemoryrepository

import (
	"context"

	"github.com/jonloureiro/tinybank/internal/account"
)

var _ account.FindManyRepository = (*inMemoryRepository)(nil)

func (repo *inMemoryRepository) FindMany(
	ctx context.Context,
) ([]account.Account, error) {
	_accounts := make([]account.Account, 0, len(repo.storageByID))
	for _, acc := range repo.storageByID {
		_accounts = append(_accounts, *acc)
	}

	return _accounts, nil
}
