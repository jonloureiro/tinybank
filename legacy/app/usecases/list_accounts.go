package usecases

import "github.com/jonloureiro/tiny-bank/legacy/app/entities"

type ListAccountInput struct {
}

type ListAccountOutput struct {
	Accounts []*entities.Account
}

func (uC *TinyBankUseCases) ListAccount(input ListAccountInput) (*ListAccountOutput, error) {
	accounts, err := uC.AccountsRepo.List()
	if err != nil {
		return nil, ErrDatabaseUnknownError
	}
	return &ListAccountOutput{Accounts: accounts}, nil
}
