package entities

import (
	"errors"
	"time"

	"github.com/jonloureiro/tinybank/legacy/app/vo"
	"github.com/jonloureiro/tinybank/legacy/extensions/id"
	"github.com/jonloureiro/tinybank/legacy/extensions/jwt"
)

var (
	ErrEmptyName          = errors.New("empty name")
	ErrEmptySecret        = errors.New("empty secret")
	ErrInvalidCredentials = errors.New("invalid credentials")
)

type Account struct {
	ID        id.ID
	Name      string
	CPF       *vo.CPF
	Secret    string
	CreatedAt time.Time
}

func NewAccount(name string, cpf *vo.CPF, secret string) (*Account, error) {
	account := Account{
		ID:        id.New(),
		Name:      name,
		CPF:       cpf,
		Secret:    secret,
		CreatedAt: time.Now(),
	}
	err := account.validate()
	if err != nil {
		return nil, err
	}
	return &account, nil
}

func (a *Account) validate() error {
	if a.Name == "" {
		return ErrEmptyName
	}
	if a.Secret == "" {
		return ErrEmptySecret
	}
	return nil
}

func (a *Account) Authenticate(secret, privateKey string) (*jwt.Token, error) {
	if a.Secret != secret {
		return nil, ErrInvalidCredentials
	}
	token, _ := jwt.New(a.ID, privateKey)
	return token, nil
}
