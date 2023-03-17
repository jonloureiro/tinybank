package usecases_test

import (
	"testing"

	"github.com/jonloureiro/tiny-bank/app/domain/entities"
	"github.com/jonloureiro/tiny-bank/app/domain/usecases"
	"github.com/jonloureiro/tiny-bank/app/domain/usecases/repositories/mocks"
	"github.com/jonloureiro/tiny-bank/app/domain/vo"
)

func TestCreateAccount(t *testing.T) {
	const (
		validName   = "Test"
		validSecret = "123456"
		validCPF    = "69029890100"
		invalidCPF  = "1"
	)

	t.Run("create account", func(t *testing.T) {
		accountsRepo := mocks.NewAccountsRepositoryMock()
		uC := usecases.TinyBankUseCases{AccountsRepo: accountsRepo}
		input := usecases.CreateAccountInput{
			Name:   validName,
			CPF:    validCPF,
			Secret: validSecret,
		}
		output, err := uC.CreateAccount(input)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
		account, err := uC.AccountsRepo.FindByID(output.AccountID)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
		if account.Name != input.Name {
			t.Errorf("expected name %s, got %s", input.Name, account.Name)
		}
	})

	t.Run("validate", func(t *testing.T) {
		testCases := map[string]struct {
			input usecases.CreateAccountInput
			err   error
		}{
			"invalid cpf": {
				input: usecases.CreateAccountInput{
					Name:   validName,
					CPF:    invalidCPF,
					Secret: validSecret,
				},
				err: vo.ErrInvalidLength,
			},
			"account already exists": {
				input: usecases.CreateAccountInput{
					Name:   validName,
					CPF:    mocks.CPFAlreadyExists,
					Secret: validSecret,
				},
				err: usecases.ErrAccountAlreadyExists,
			},
			"invalid input": {
				input: usecases.CreateAccountInput{
					Name:   "",
					CPF:    validCPF,
					Secret: validSecret,
				},
				err: entities.ErrEmptyName,
			},
			"id already exists in base": {
				input: usecases.CreateAccountInput{
					Name:   validName,
					CPF:    mocks.CPFAlreadyExists,
					Secret: validSecret,
				},
				err: usecases.ErrAccountAlreadyExists,
			},
			"error saving account": {
				input: usecases.CreateAccountInput{
					Name:   validName,
					CPF:    mocks.CPFWithUnknownError,
					Secret: validSecret,
				},
				err: usecases.ErrDatabaseUnknownError,
			},
		}
		for desc, tC := range testCases {
			t.Run(desc, func(t *testing.T) {
				accountsRepo := mocks.NewAccountsRepositoryMock()
				uC := usecases.TinyBankUseCases{AccountsRepo: accountsRepo}
				_, err := uC.CreateAccount(tC.input)
				if err != tC.err {
					t.Errorf("expected error %v, got %v", tC.err, err)
				}
			})
		}
	})
}
