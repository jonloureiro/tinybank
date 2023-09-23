package entities_test

import (
	"testing"

	"github.com/jonloureiro/tiny-bank/legacy/app/entities"
	"github.com/jonloureiro/tiny-bank/legacy/app/vo"
	"github.com/jonloureiro/tiny-bank/legacy/extensions/jwt"
)

var validCPF, _ = vo.NewCPF("69029890100")

func TestNewAccount(t *testing.T) {
	t.Run("create account", func(t *testing.T) {
		account, err := entities.NewAccount("Jon", validCPF, "123456")
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
		if account == nil {
			t.Errorf("expected an account, got %v", account)
		}
	})

	t.Run("validate errors", func(t *testing.T) {
		testCases := map[string]struct {
			name        string
			cpf         *vo.CPF
			secret      string
			expectedErr error
		}{
			"empty name": {
				name:        "",
				cpf:         validCPF,
				secret:      "123456",
				expectedErr: entities.ErrEmptyName,
			},
			"empty secret": {
				name:        "Jon",
				cpf:         validCPF,
				secret:      "",
				expectedErr: entities.ErrEmptySecret,
			},
		}
		for desc, tC := range testCases {
			t.Run(desc, func(t *testing.T) {
				account, err := entities.NewAccount(tC.name, tC.cpf, tC.secret)
				if err != tC.expectedErr {
					t.Errorf("expected error %v, got %v", tC.expectedErr, err)
				}
				if account != nil {
					t.Errorf("expected no account, got %v", account)
				}
			})
		}
	})
}

func TestAuthenticateAccount(t *testing.T) {
	privateKey := "s3cr3t"
	secret := "123456"
	account, err := entities.NewAccount("Jon", validCPF, secret)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	got, _ := account.Authenticate(secret, privateKey)
	want, _ := jwt.Parse(got.Token, privateKey)
	if got.Token != want.Token {
		t.Errorf("expected token %v, got %v", want.Token, got.Token)
	}
}
