package inmemoryrepository

import (
	"context"

	"github.com/jonloureiro/tinybank/internal/account"
)

var _ account.FindRepository = (*inMemoryRepository)(nil)

func (repo *inMemoryRepository) Find(
	ctx context.Context, id string,
) (account.Account, error) {
	_account, ok := repo.storageByID[id]
	if !ok || _account == nil {
		return nil, account.ErrAccountNotFound
	}

	return *_account, nil
}
