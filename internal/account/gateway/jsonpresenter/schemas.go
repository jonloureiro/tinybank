package jsonpresenter

type accountIDSchema struct {
	AccountID string `json:"id"`
}

type accountSchema struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	CPF       string `json:"-"`
	Balance   int    `json:"balance"`
	CreatedAt string `json:"created_at"`
}

type accountsSchema []accountSchema

type accountBalanceSchema struct {
	Balance int `json:"balance"`
}
