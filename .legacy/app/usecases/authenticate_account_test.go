package usecases_test

import (
	"testing"

	"github.com/jonloureiro/tinybank/legacy/app/entities"
	"github.com/jonloureiro/tinybank/legacy/app/usecases"
	"github.com/jonloureiro/tinybank/legacy/app/usecases/repositories/mocks"
	"github.com/jonloureiro/tinybank/legacy/app/vo"
	"github.com/jonloureiro/tinybank/legacy/extensions/jwt"
)

func TestAuthenticate(t *testing.T) {
	const (
		validName   = "Test"
		validSecret = "123456"
		privateKey  = "s3cr3t"
	)

	t.Run("", func(t *testing.T) {
		accountsRepo := mocks.NewAccountsRepositoryMock()
		cpf, _ := vo.NewCPF(mocks.ValidCPF)
		account, _ := entities.NewAccount(validName, cpf, validSecret)
		accountsRepo.Create(account)
		uC := usecases.TinyBankUseCases{privateKey, accountsRepo}
		input := usecases.AuthenticateAccountInput{
			CPF:    account.CPF.Value(),
			Secret: account.Secret,
		}
		output, err := uC.AuthenticateAccount(&input)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
		if output.AccessToken.AccountId != account.ID {
			t.Errorf("expected account id %v, got %v", account.ID, output.AccessToken.AccountId)
		}
		_, err = jwt.Parse(output.AccessToken.Token, privateKey)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
	})
}
