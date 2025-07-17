package dtos

type CriarClienteRequest struct {
	Nome     string `json:"nome" validate:"required,min=2,max=100"`
	Telefone string `json:"telefone" validate:"required,min=10,max=15"`
	Email    string `json:"email" validate:"required,email"`
	Endereco string `json:"endereco" validate:"required,min=10,max=200"`
}

type AtualizarClienteRequest struct {
	Nome     *string `json:"nome,omitempty" validate:"omitempty,min=2,max=100"`
	Telefone *string `json:"telefone,omitempty" validate:"omitempty,min=10,max=15"`
	Email    *string `json:"email,omitempty" validate:"omitempty,email"`
	Endereco *string `json:"endereco,omitempty" validate:"omitempty,min=10,max=200"`
}

type AlterarStatusClienteRequest struct {
	Status string `json:"status" validate:"required,oneof=ativo inativo"`
}

type ClienteResponse struct {
	ID        int    `json:"id"`
	Nome      string `json:"nome"`
	Telefone  string `json:"telefone"`
	Email     string `json:"email"`
	Endereco  string `json:"endereco"`
	Status    string `json:"status"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
