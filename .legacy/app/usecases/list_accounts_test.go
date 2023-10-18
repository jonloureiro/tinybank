package usecases_test

import (
	"testing"

	"github.com/jonloureiro/tinybank/legacy/app/entities"
	"github.com/jonloureiro/tinybank/legacy/app/usecases"
	"github.com/jonloureiro/tinybank/legacy/app/usecases/repositories/mocks"
	"github.com/jonloureiro/tinybank/legacy/app/vo"
)

func TestListAccounts(t *testing.T) {
	var (
		validName   = "Test"
		validSecret = "123456"
		validCPF, _ = vo.NewCPF("69029890100")
	)
	t.Run("list accounts", func(t *testing.T) {
		accountsRepo := mocks.NewAccountsRepositoryMock()
		uC := usecases.TinyBankUseCases{AccountsRepo: accountsRepo}
		account, _ := entities.NewAccount(validName, validCPF, validSecret)
		_ = uC.AccountsRepo.Create(account)

		output, err := uC.ListAccount(usecases.ListAccountInput{})
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}

		if len(output.Accounts) != 1 {
			t.Errorf("expected 1 account, got %d", len(output.Accounts))
		}
	})

	t.Run("validate empty slice", func(t *testing.T) {
		want := make([]*entities.Account, 0)
		accountsRepo := mocks.NewAccountsRepositoryMock()
		uC := usecases.TinyBankUseCases{AccountsRepo: accountsRepo}
		output, err := uC.ListAccount(usecases.ListAccountInput{})
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
		if len(output.Accounts) != len(want) {
			t.Errorf("expected %d accounts, got %d", len(want), len(output.Accounts))
		}
	})

	t.Run("validate error", func(t *testing.T) {
		accountsRepo := mocks.NewAccountsRepositoryMock()
		accountsRepo.UnknownError = true
		uC := usecases.TinyBankUseCases{AccountsRepo: accountsRepo}
		_, err := uC.ListAccount(usecases.ListAccountInput{})
		if err != usecases.ErrDatabaseUnknownError {
			t.Errorf("expected error %v, got %v", usecases.ErrDatabaseUnknownError, err)
		}
	})
}
