package vos

import (
	"fmt"

	"github.com/jonloureiro/tiny-bank/internal/common"
)

type CPF struct {
	value string
}

var nilCPF = CPF{}

func NewCPF(data string) (CPF, error) {
	c := CPF{data}

	if err := c.validate(); err != nil {
		return nilCPF, fmt.Errorf("%w:invalid cpf", err)
	}

	return c, nil
}

func (c CPF) validate() error {
	if len(c.value) != 11 {
		return fmt.Errorf("%w:incompatible length", common.ErrFailedDependency)
	}
	// TODO: to implement
	return nil
}

func (c CPF) Value() string {
	return c.value
}
