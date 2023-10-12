package accounts

import "context"

type FindAccountRepository interface {
	Find(context.Context, string) (Account, error)
}

type FindManyAccountsRepository interface {
	FindMany(context.Context) ([]Account, error)
}

type SaveAccountRepository interface {
	Save(context.Context, Account) error
}
