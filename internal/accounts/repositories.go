package accounts

import (
	"context"
)

type SaveAccountsRepository interface {
	Save(context.Context, Account) error
}
