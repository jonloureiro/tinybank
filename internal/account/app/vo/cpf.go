package vo

import (
	"github.com/jonloureiro/tinybank/internal/account"
)

type CPF struct {
	value string
}

var nilCPF = CPF{}

func ParseCPF(data string) (CPF, error) {
	c := CPF{data}

	if err := c.validate(); err != nil {
		return nilCPF, err
	}

	return c, nil
}

func (c CPF) validate() error {
	if len(c.value) != 11 {
		return account.ErrIncompatibleCPFLength
	}
	// TODO: to implement
	return nil
}

func (c CPF) Value() string {
	return c.value
}
