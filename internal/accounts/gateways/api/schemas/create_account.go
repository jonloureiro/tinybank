package schemas

type CreateAccountRequest struct {
	Name   string `json:"name"`
	CPF    string `json:"cpf"`
	Secret string `json:"secret"`
}

type CreateAccountResponse struct {
	AccountID string `json:"id"`
}
