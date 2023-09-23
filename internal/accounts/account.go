package accounts

import "time"

type Account interface {
	ID() string
	Name() string
	CPF() string
	Secret() string
	Balance() int
	CreatedAt() time.Time
}
