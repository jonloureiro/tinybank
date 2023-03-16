package vo

import "errors"

var (
	ErrInvalidLength = errors.New("invalid cpf, incompatible length")
)

type CPF struct {
	data string
}

func NewCPF(data string) (*CPF, error) {
	c := CPF{data}
	err := c.validate()
	if err != nil {
		return nil, err
	}
	return &c, nil
}

func (c *CPF) validate() error {
	if len(c.data) != 11 {
		return ErrInvalidLength
	}
	// TODO: to implement
	return nil
}

func (c *CPF) Value() string {
	return c.data
}
