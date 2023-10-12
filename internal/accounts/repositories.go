package accounts

import "context"

type FindManyAccountsRepository interface {
	FindMany(context.Context) ([]Account, error)
}

type SaveAccountRepository interface {
	Save(context.Context, Account) error
}
