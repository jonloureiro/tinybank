package usecases

import (
	"github.com/jonloureiro/tiny-bank/legacy/app/vo"
	"github.com/jonloureiro/tiny-bank/legacy/extensions/jwt"
)

type AuthenticateAccountInput struct {
	CPF    string
	Secret string
}

type AuthenticateAccountOutput struct {
	AccessToken *jwt.Token
}

func (uC *TinyBankUseCases) AuthenticateAccount(input *AuthenticateAccountInput) (*AuthenticateAccountOutput, error) {
	cpf, err := vo.NewCPF(input.CPF)
	if err != nil {
		return nil, err
	}
	account, err := uC.AccountsRepo.FindByCPF(cpf)
	if err != nil {
		return nil, err
	}
	token, err := account.Authenticate(input.Secret, uC.PrivateKey)
	if err != nil {
		return nil, err
	}
	return &AuthenticateAccountOutput{token}, nil
}
