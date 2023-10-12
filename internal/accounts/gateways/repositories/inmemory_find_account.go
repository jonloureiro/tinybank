package repositories

import (
	"context"
	"fmt"

	"github.com/jonloureiro/tiny-bank/internal/accounts"
	"github.com/jonloureiro/tiny-bank/internal/common"
)

var _ accounts.FindAccountRepository = (*accountsRepositoryInMemory)(nil)

func (repo *accountsRepositoryInMemory) Find(
	ctx context.Context, id string,
) (accounts.Account, error) {
	_account, ok := repo.storageByID[id]
	if !ok || _account == nil {
		return nil, fmt.Errorf("%w:FindAccountRepository", common.ErrNotFound)
	}

	return *_account, nil
}
