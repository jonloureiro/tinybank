package entities

import (
	"errors"
	"time"

	"github.com/jonloureiro/go-challenge/app/domain/vo"
	"github.com/jonloureiro/go-challenge/extensions/id"
)

var (
	ErrEmptyName   = errors.New("empty name")
	ErrEmptySecret = errors.New("empty secret")
)

type Account struct {
	ID        string
	Name      string
	CPF       vo.CPF
	Secret    string
	CreatedAt time.Time
}

func NewAccount(name string, cpf string, secret string) (*Account, error) {
	c, err := vo.NewCPF(cpf)
	if err != nil {
		return nil, err
	}
	a := Account{
		ID:        id.New(),
		Name:      name,
		CPF:       *c,
		Secret:    secret,
		CreatedAt: time.Now(),
	}
	err = a.validate()
	if err != nil {
		return nil, err
	}
	return &a, nil
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
